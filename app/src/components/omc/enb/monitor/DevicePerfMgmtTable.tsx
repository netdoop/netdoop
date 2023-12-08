import React, { useEffect, useRef } from 'react';
import { useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Modal, Tooltip } from 'antd';


import { fetchDevices, setDevicePerfEnable, setDevicePerfDisable } from '@/models/device';
import { useGroups } from '@/models/groups';
import { FetchParams, SearchItem, updateSearchItemWithValue, updateSearchItemWithValues } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { formatTimestamp2 } from '@/utils/format';
import { PauseCircleOutlined, PlayCircleOutlined, UploadOutlined } from '@ant-design/icons';
import { proTableLayout } from '@/constants/style';

interface Props {
  groupIds: number[],
}

const DevicePerfMgmtTable: React.FC<Props> = ({
  groupIds
}) => {
  const intl = useIntl()
  const ref = useRef<ActionType>();
  const { groupNameById } = useGroups();

  useEffect(() => {
    ref.current?.reload();
  }, [groupIds, groupNameById]);

  const handleRequest = async (params: {
    groupIds: number[];
    SerialNumber?: string;
    Name?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", "enb")
    searchItems = updateSearchItemWithValues(searchItems, "group", params.groupIds)
    searchItems = updateSearchItemWithValue(searchItems, "serial_number", params.SerialNumber)
    searchItems = updateSearchItemWithValue(searchItems, "name", params.Name)

    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchDevices(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleEnable = (record: API.Device) => {
    Modal.confirm({
      title: 'Do you want to enable performance upload?',
      okText: 'Confirm',
      cancelText: 'Cancel',
      onOk: () => {
        setDevicePerfEnable(record)
      },
    });
  };
  const handleDisable = (record: API.Device) => {
    Modal.confirm({
      title: 'Do you want to disable performance upload?',
      okText: 'Confirm',
      cancelText: 'Cancel',
      onOk: () => {
        setDevicePerfDisable(record)
      },
    });
  };

  const columns: ProColumns[] = [
    {
      title: 'ID',
      dataIndex: 'Id',
      key: 'Id',
      search: false,
    },
    {
      title: 'Status',
      dataIndex: ['Properties', 'PerfMgmtEnable'],
      key: 'PerfMgmtEnable',
      search: false,
      render: (text: any) => {
        const v: string = text as string;
        // const {} = record.Properties;
        return (
          <>
            {v === "1" && (<UploadOutlined color='green' />)}
            {v !== "1" && (<UploadOutlined color='grey' />)}
          </>
        )
      },
    },
    {
      title: 'Serial Number',
      dataIndex: 'SerialNumber',
      key: 'SerialNumber',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Serial Number",
      },
    },
    {
      title: 'Cell Name',
      dataIndex: 'Name',
      key: 'Name',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Name",
      },
    },
    {
      title: 'Periodic Upload Interval',
      dataIndex: ['Properties', 'PerfMgmtPeriodicUploadInterval'],
      key: 'PerfMgmtPeriodicUploadInterval',
      search: false,
    },
    {
      title: 'Last Upload Time',
      dataIndex: ['Properties', 'PerfMgmtLastUploadTime'],
      key: 'PerfMgmtLastUploadTime',
      search: false,
      render: (text: any) => formatTimestamp2(text, 1000000),
    },
    {
      title: 'OUI',
      dataIndex: 'Oui',
      key: 'Oui',
      search: false,
    },
    {
      title: 'Product Class',
      dataIndex: 'ProductClass',
      key: 'ProductClass',
      search: false,
    },
    {
      title: 'Group',
      dataIndex: 'Group',
      key: 'Group',
      search: false,
      render: (text: any, record: API.Device) => <span>{groupNameById(record.GroupId)}</span>,
    },
    {
      title: 'Operations',
      dataIndex: 'Operations',
      key: 'Operations',
      search: false,
      render: (text: any, record: API.Device) => {
        const data = record.Properties as unknown as { PerfMgmtEnable?: string, PerfMgmtPeriodicUploadInterval?: string };
        return (
          <>
            {data.PerfMgmtEnable === "1" && (<Tooltip title={intl.formatMessage({ id: 'common.stop' })}><PauseCircleOutlined onClick={() => handleDisable(record)} /> </Tooltip>)}
            {data.PerfMgmtEnable !== "1" && (<Tooltip title={intl.formatMessage({ id: 'common.start' })}><PlayCircleOutlined onClick={() => handleEnable(record)} /> </Tooltip>)}
          </>
        )
      },
    },
  ];

  return (
    <>
      <ProTable
        {...proTableLayout}
        rowKey="Id"
        columns={columns}
        actionRef={ref}
        params={{ groupIds: groupIds }}
        request={handleRequest}
        // dataSource={devices}
        // loading={loading}
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
      />
    </>
  );
};

export default DevicePerfMgmtTable;
