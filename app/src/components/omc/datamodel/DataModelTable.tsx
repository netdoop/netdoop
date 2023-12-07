
import React, { useRef, useState } from 'react';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Switch, Space, Button, Dropdown, MenuProps } from 'antd';
import { DeleteOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';
import { DeleteDataModelForm } from '@/components/omc/datamodel';
import { fetchDataModels } from '@/models/datamodel';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { useIntl, history } from '@umijs/max';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const DataModelTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl();
  const ref = useRef<ActionType>();

  const [selectedDataModel, setSelectedDataModel] = useState<API.DataModel | undefined>(undefined);
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
    const result = await fetchDataModels(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleDeleteDataModel = async (record: API.DataModel) => {
    setSelectedDataModel(record);
    setDeleteModalVisible(true);
  };

  const moreItems = (record: API.DataModel): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'delete',
        disabled: record.Default || !access.canDeleteOMCDataModel,
        icon: (<DeleteOutlined />),
        label: (
          <a onClick={() => handleDeleteDataModel(record)}>
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
      render: (text: any, record: API.DataModel) => (
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
      title: 'Product Type',
      dataIndex: 'ProductType',
      key: 'ProductType',
      search: false,
    },
    {
      title: 'Parameter Path',
      dataIndex: 'ParameterPath',
      key: 'ParameterPath',
      search: false,
    },
    {
      title: 'Default',
      dataIndex: 'Default',
      key: 'Default',
      search: false,
      render: (text: any, record: API.DataModel) => (<Switch checked={record.Default} />),
    },
    {
      title: 'Created',
      dataIndex: 'Created',
      key: 'Created',
      search: false,
      render: (text: any, record: API.DataModel) => <span>{formatTimestamp2(record.Created)}</span>,
    },
    {
      title: 'Updated',
      dataIndex: 'Updated',
      key: 'Updated',
      search: false,
      render: (text: any, record: API.DataModel) => <span>{formatTimestamp2(record.Updated)}</span>,
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
      <DeleteDataModelForm
        visible={deleteModalVisible}
        record={selectedDataModel}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default DataModelTable;