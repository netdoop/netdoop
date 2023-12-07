import React, { useEffect, useRef } from 'react';
import { history, useAccess, useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Dropdown, Button, MenuProps } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { fetchAuditLogs } from '@/models/audit_logs';
import { proTableLayout } from '@/constants/style';
import { Access } from '@umijs/max';



const AuditLogsTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  useEffect(() => {
    ref.current?.reload();
  }, []);

  const handleRequest = async (params: {
    ApiName?: string;
    Method?: string;

  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "api_name", params.ApiName)
    searchItems = updateSearchItemWithValue(searchItems, "method", params.Method)

    sort["time"] = "descend"
    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchAuditLogs(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleDetail = async (record: API.AuditLog) => {
    history.push(history.location.pathname + '/audit-logs/' + record.Time + '/detail')
  };

  // Define the actions for the table
  const moreItems = (record: API.AuditLog): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'detail',
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
      render: (text: any, record: API.AuditLog) => (
        <div>
          {formatTimestamp2(record.Time)}
          <Dropdown menu={moreItems(record)}>
            <Access accessible={access.canGetAuditLog}>
              <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
                <MoreOutlined />
              </Button>
            </Access>
          </Dropdown>
        </div>
      ),
    },
    {
      title: 'User',
      dataIndex: 'UserName',
      key: 'UserName',
    },
    {
      title: 'API',
      dataIndex: 'ApiName',
      key: 'ApiName',
    },
    {
      title: 'Method',
      dataIndex: 'Method',
      key: 'Method',
    },
    {
      title: 'Path',
      dataIndex: 'Path',
      key: 'Path',
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

export default AuditLogsTable;