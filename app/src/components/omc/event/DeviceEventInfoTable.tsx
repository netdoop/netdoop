import React, { useEffect, useRef } from 'react';
import { history, useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Dropdown, Button, MenuProps } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';
import { fetchDeviceEvents } from '@/models/device_event';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

interface Props {
  productType: string;
}

const DeviceEventInfoTable: React.FC<Props> = ({ productType }) => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  useEffect(() => {
    ref.current?.reload();
  }, [productType]);

  const handleRequest = async (params: {
    EventType?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
    searchItems = updateSearchItemWithValue(searchItems, "event_type", params.EventType)

    sort["time"] = "descend"
    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchDeviceEvents(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleDetail = async (record: API.DeviceEvent) => {
    history.push(history.location.pathname + '/' + record.Time + '/detail')
  };

  // Define the actions for the table
  const moreItems = (record: API.DeviceEvent): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'detail',
        disabled: !access.canGetOMCDeviceEvent,
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
      render: (text: any, record: API.DeviceEvent) => (
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
      title: 'Event Type',
      dataIndex: 'EventType',
      key: 'EventType',
      valueType: 'select',
      valueEnum: {
        "Inform": { text: "Inform" },
      },
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Event Type",
        defaultValue: "Inform",
      },
    },
    {
      title: 'Details',
      dataIndex: 'Details',
      key: 'Details',
      search: false,
      render: (text: any, record: API.DeviceEvent) => {
        const events = record.MetaData as unknown as Record<string, string> || {};
        return (
          <>
            {Object.entries(events).map(([key]) => (
              <span key={key}>{key}</span>
            ))}
          </>
        );
      },
    },
    {
      title: 'Network Element',
      dataIndex: 'Device',
      key: 'Device',
      search: false,
      render: (text: any, record: API.DeviceEvent) => (
        <span>Model={record.ProductClass} SN={record.SerialNumber}</span>
      ),
    },
    {
      title: 'Current Time',
      dataIndex: 'CurrentTime',
      key: 'CurrentTime',
      search: false,
      render: (text: any, record: API.DeviceEvent) => formatTimestamp2(record.CurrentTime || undefined),
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

export default DeviceEventInfoTable;