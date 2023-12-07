import React, { useEffect, useRef } from 'react';
import { history, useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Dropdown, Button, MenuProps } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';
import { fetchDeviceMethodCalls } from '@/models/device_method_call';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

interface Props {
  productType: string;
}

const DeviceMethodCallTable: React.FC<Props> = ({ productType }) => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  useEffect(() => {
    ref.current?.reload();
  }, [productType]);

  const handleRequest = async (params: {
    EventType?: string;
    State?: number;
    FaultCode?: number;
    FaultString?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
    searchItems = updateSearchItemWithValue(searchItems, "event_type", params.EventType)
    searchItems = updateSearchItemWithValue(searchItems, "state", params.State)

    searchItems = updateSearchItemWithValue(searchItems, "fault_code", params.FaultCode)
    searchItems = updateSearchItemWithValue(searchItems, "fault_string", params.FaultString)

    sort["time"] = "descend"
    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchDeviceMethodCalls(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleDetail = async (record: API.DeviceMethodCall) => {
    history.push(history.location.pathname + '/' + record.Time + '/detail')
  };

  // Define the actions for the table
  const moreItems = (record: API.DeviceMethodCall): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'detail',
        disabled: !access.canGetOMCDeviceMethodCall,
        label: (
            <a onClick={() => handleDetail(record)}>
              {intl.formatMessage({ id: 'common.detail' })}
            </a>
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
      render: (text: any, record: API.DeviceMethodCall) => (
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
      title: 'Method Name',
      dataIndex: 'MethodName',
      key: 'MethodName',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Method Name",
      },
    },
    {
      title: 'State',
      dataIndex: 'State',
      key: 'State',
      valueType: 'select',
      valueEnum: {
        0: { text: "Queued" },
        1: { text: "Request Send" },
        2: { text: "Response Received" },
        3: { text: "Response Timeout" },
        100: { text: "Unknow" },
      },
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "State",
      },
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
      render: (text: any, record: API.DeviceMethodCall) => (
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
    {
      title: 'Network Element',
      dataIndex: 'Device',
      key: 'Device',
      search: false,
      render: (text: any, record: API.DeviceMethodCall) => (
        <span>Model={record.ProductClass} SN={record.SerialNumber}</span>
      ),
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

export default DeviceMethodCallTable;