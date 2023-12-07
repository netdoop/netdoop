import React, { useRef } from 'react';
import { history, useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Dropdown, Button, MenuProps, Tag } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';

import { SEVERITY_COLOR_MAP } from '@/models/alarm';

import { fetchDeviceAlarms } from '@/models/device_alarm';
import { FetchParams, SearchItem, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const AlarmInfoTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  const handleRequest = async (params: {
    PerceivedSeverity?: string;
    AlarmIdentifier?: string;
    EventType?: string;
    AlarmCleared?: number;
    SerialNumber?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "perceived_severity", params?.PerceivedSeverity)
    searchItems = updateSearchItemWithValue(searchItems, "alarm_identifier", params?.AlarmIdentifier)
    searchItems = updateSearchItemWithValue(searchItems, "event_type", params?.EventType)
    searchItems = updateSearchItemWithValue(searchItems, "alarm_cleared", params?.AlarmCleared)
    searchItems = updateSearchItemWithValue(searchItems, "serial_number", params?.SerialNumber)

    sort["time"] = "descend"

    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchDeviceAlarms(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleDetail = async (record: API.DeviceAlarm) => {
    history.push(history.location.pathname + '/' + record.Time + '/detail')
  };

  // Define the actions for the table
  const moreItems = (record: API.DeviceAlarm): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'detail',
        disabled: !access.canViewOMCAlarmDetailPage,
        label: (
            <a onClick={() => handleDetail(record)}>
              {intl.formatMessage({ id: 'common.detail' })}
            </a>
        ),
      },
    ]
    return { items }
  }


  const columns: ProColumns[] = [
    {
      title: 'Time',
      dataIndex: 'Time',
      key: 'Time',
      fixed: 'left' as 'left',
      width: 180,
      search: false,
      render: (text: any, record: API.DeviceAlarm) => (
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
      title: 'Perceived Severity',
      dataIndex: 'PerceivedSeverity',
      key: 'PerceivedSeverity',
      valueType: 'select',
      valueEnum: {
        "": { text: "All" },
        "Critical": { text: "Critical" },
        "Major": { text: "Major" },
        "Minor": { text: "Minor" },
        "Warning": { text: "Warning" },
      },
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Perceived Severity",
        defaultValue: "",
      },
      render: (text: any, record: API.DeviceAlarm) => (
        <Tag bordered style={{ width: 56, textAlign: 'center' }} color={SEVERITY_COLOR_MAP[record.PerceivedSeverity || 'Unknown']}>
          {record.PerceivedSeverity}
        </Tag>
      ),
    },
    {
      title: 'Alarm Identifier',
      dataIndex: 'AlarmIdentifier',
      key: 'AlarmIdentifier',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Alarm Identifier",
      },
    },
    {
      title: 'Network Element',
      dataIndex: 'Device',
      key: 'Device',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Network Element",
      },
      render: (text: any, record: API.DeviceAlarm) => (
        <span>Model={record.ProductClass} SN={record.SerialNumber}</span>
      ),
    },
    {
      title: 'Event Type',
      dataIndex: 'EventType',
      key: 'EventType',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Event Type",
      },
    },
    {
      title: 'Alarm Status',
      dataIndex: 'AlarmCleared',
      key: 'AlarmCleared',
      render: (text: any, record: API.DeviceAlarm) => {
        if (record.AlarmCleared) {
          return "Cleared"
        } else {
          return "Active"
        }
      },
      valueType: 'select',
      valueEnum: {
        "1": { text: "Cleared" },
        "0": { text: "Active" },
      },
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Active",
        defaultValue: "0",
      },
    },
    {
      title: 'Event Time',
      dataIndex: 'AlarmRaisedTime',
      key: 'AlarmRaisedTime',
      search: false,
      render: (text: any, record: API.DeviceAlarm) => formatTimestamp2(record.AlarmRaisedTime),
    },
    {
      title: 'Update Time',
      dataIndex: 'AlarmChangedTime',
      key: 'AlarmChangedTime',
      search: false,
      render: (text: any, record: API.DeviceAlarm) => formatTimestamp2(record.AlarmChangedTime),
    },
    {
      title: 'Probable Cause',
      dataIndex: 'ProbableCause',
      key: 'ProbableCause',
      search: false,
    },
    {
      title: 'Specific Problem',
      dataIndex: 'SpecificProblem',
      key: 'SpecificProblem',
      search: false,
    },
  ];


  return (
    <>
      <ProTable
        {...proTableLayout}
        rowKey="Time"
        params={{}}
        request={handleRequest}
        columns={columns}
        actionRef={ref}
        search={{
          span: 6,
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
            {/* TODO: add create alarm button */}
          </Space>]}
      />
    </>
  );
};

export default AlarmInfoTable;