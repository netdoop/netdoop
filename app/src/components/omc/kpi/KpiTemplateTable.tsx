
import React, { useRef, useState } from 'react';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Switch, Space, Button, Dropdown, MenuProps } from 'antd';
import { DeleteOutlined, EditOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';
import { DeleteKpiTemplateForm } from '@/components/omc/kpi';
import { fetchKPITemplates } from '@/models/kpi_temp';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { useIntl, history } from '@umijs/max';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const KpiTemplateTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl();
  const ref = useRef<ActionType>();

  const [selectedKpiTemplate, setSelectedKpiTemplate] = useState<API.KPITemplate | undefined>(undefined);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);

  const handleRequest = async (params: {
    Name?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "name", params.Name)

    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchKPITemplates(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  // Handle delete button click
  const handleDeleteKpiTemplate = async (record: API.KPITemplate) => {
    setSelectedKpiTemplate(record);
    setDeleteModalVisible(true);
  };

  const handleEditKpiTemplate = (record: API.KPITemplate) => {
    history.push('/omc/perf/KpiTemp/edit/' + record.Id)
  };

  const moreItems = (record: API.KPITemplate): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'edit',
        disabled: record.Default || !access.canUpdateOMCKPITemplate,
        icon: (<EditOutlined />),
        label: (
          <a onClick={() => handleEditKpiTemplate(record)}>
            {intl.formatMessage({ id: 'common.edit' })}
          </a>
        ),
      },
      {
        key: 'delete',
        disabled: record.Default || !access.canDeleteOMCKPITemplate,
        icon: (<DeleteOutlined />),
        label: (
          <a onClick={() => handleDeleteKpiTemplate(record)}>
            {intl.formatMessage({ id: 'common.delete' })}
          </a>
        ),
      },
    ];
    return { items };
  }

  const columns: ProColumns[] = [
    {
      title: 'ID',
      dataIndex: 'Id',
      key: 'Id',
      search: false,
      render: (text: any, record: API.KPITemplate) => (
        <>
          {text}
          <Dropdown menu={moreItems(record)}>
            <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
              <MoreOutlined />
            </Button>
          </Dropdown>
        </>
      ),
    },
    {
      title: 'Name',
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
      title: 'Periodic Interval',
      dataIndex: 'PeriodicInterval',
      key: 'PeriodicInterval',
      search: false,
    },
    {
      title: 'Select Type',
      dataIndex: 'SelectType',
      key: 'SelectType',
      search: false,
    },
    {
      title: 'Default',
      dataIndex: 'Default',
      key: 'Default',
      search: false,
      render: (text: any, record: API.KPITemplate) => (<Switch checked={record.Default} />),
    },
    {
      title: 'Created',
      dataIndex: 'Created',
      key: 'Created',
      search: false,
      render: (text: any, record: API.KPITemplate) => <span>{formatTimestamp2(record.Created)}</span>,
    },
    {
      title: 'Updated',
      dataIndex: 'Updated',
      key: 'Updated',
      search: false,
      render: (text: any, record: API.KPITemplate) => <span>{formatTimestamp2(record.Updated)}</span>,
    },
  ];
  return (
    <>
      <ProTable
        {...proTableLayout}
        rowKey="Id"
        columns={columns}
        request={handleRequest}
        actionRef={ref}
        search={{
          span: 4,
          labelWidth: 0,
        }}
        options={{
          density: false,
          fullScreen: true,
          setting: true,
          reload: true,
        }}
        toolBarRender={() => [
          <Space key="custom-options">
            <PlusOutlined
              className="ant-pro-table-toolbar-item-iconButton"
              onClick={() => history.push('/omc/perf/KpiTemp/create')}
              style={{ fontSize: '20px' }}
            />
          </Space>

        ]}
      />
      <DeleteKpiTemplateForm
        visible={deleteModalVisible}
        record={selectedKpiTemplate}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default KpiTemplateTable;