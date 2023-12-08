import React, { useRef } from 'react';
import { useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Dropdown, Button, MenuProps } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { formatBytes, formatTimestamp2 } from '@/utils/format';
import { fetchDeviceTransferLogs } from '@/models/device_transfer';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

interface DeviceUploadFileTableProps {
  productType: string,
  type: string;
}

const DeviceUploadFileTable: React.FC<DeviceUploadFileTableProps> = ({ productType, type }) => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  const handleRequest = async (params: {
    FaultCode?: number;
    FaultString?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
    searchItems = updateSearchItemWithValue(searchItems, "transfer_type", "upload")
    searchItems = updateSearchItemWithValue(searchItems, "file_type", type)
    searchItems = updateSearchItemWithValue(searchItems, "fault_code", params.FaultCode)
    searchItems = updateSearchItemWithValue(searchItems, "fault_string", params.FaultString)

    sort["time"] = "descend"
    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchDeviceTransferLogs(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  // Define the actions for the table
  const moreItems = (record: API.DeviceTransferLog): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'download',
        disabled: record.S3Object === undefined || !access.canGetOMCTransferLog,
        label: (
          <>
            {record.S3Object ? <a href={record.S3Object?.DownloadUrl}>
              {intl.formatMessage({ id: 'common.actions.download' })}
            </a> : <span>{intl.formatMessage({ id: 'common.actions.download' })}</span>}
          </>
        ),
      },
    ]
    return { items }
  }

  // Define the columns for the table
  const columns: ProColumns[] = [
    {
      title: 'Time',
      dataIndex: 'Time',
      key: 'Time',
      search: false,
      fixed: 'left' as 'left',
      width: 180,
      render: (text: any, record: API.DeviceTransferLog) => (
        <div>
          {formatTimestamp2(record.Time)}
          <Dropdown menu={moreItems(record)}>
            <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
              <MoreOutlined />
            </Button>
          </Dropdown>
        </div>
      ),
    },
    {
      title: 'File Name',
      dataIndex: 'FileName',
      key: 'FileName',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "File Name",
      },
    },
    {
      title: 'Network Element',
      dataIndex: 'Device',
      key: 'Device',
      search: false,
      render: (text: any, record: API.DeviceTransferLog) => (
        <span>Model={record.ProductClass} SN={record.SerialNumber}</span>
      ),
    },
    {
      title: 'File Size',
      dataIndex: ['S3Object', "FileSize"],
      key: 'File Size',
      search: false,
      render: (text: any, record: API.DeviceTransferLog) => (
        <span>{formatBytes(record.S3Object?.FileSize || -1)}</span>
      ),
    },
    {
      title: 'Fault Code',
      dataIndex: 'FaultCode',
      key: 'FaultCode',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Fault Code",
      },
      render: (text: any, record: API.DeviceTransferLog) => (
        <span>{record.FaultCode === 0 ? '-' : record.FaultCode}</span>
      ),
    },
    {
      title: 'Fault String',
      dataIndex: 'FaultString',
      key: 'FaultString',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Fault String",
      },
    },

  ];


  return (
    <>
      <ProTable
        {...proTableLayout}
        rowKey="Time"
        columns={columns}
        request={handleRequest}
        actionRef={ref}
        search={{
          span: 4,
          labelWidth: 0,
        }}
        scroll={{ x: 'max-content' }}
        options={{
          density: false,
          fullScreen: true,
          setting: true,
          reload: true,
        }}
        toolBarRender={() => [
          <Space key="custom-options">
            {/* TODO: add create event button */}
          </Space>]}
      />
    </>
  );
};

export default DeviceUploadFileTable;