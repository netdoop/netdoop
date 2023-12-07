import React, { useState, useRef } from 'react';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Button, Dropdown, MenuProps, Space, Switch } from 'antd';
import { DeleteOutlined, EditOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';
import { CreateProductForm, UpdateProductForm, DeleteProductForm } from '@/components/omc/product';
import { fetchProducts, setProductDisable, setProductEnable } from '@/models/product';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { useIntl } from '@umijs/max';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const ProductsTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  const [selectedProduct, setSelectedProduct] = useState<API.Product | null>(null);
  const [createModalVisible, setCreateModalVisible] = useState(false);
  const [updateModalVisible, setUpdateModalVisible] = useState(false);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);

  const handleRequest = async (params: {
    Oui?: string;
    ProductClass?: string;
    Menufacture?: string;
    Name?: string;
    Enable?: boolean;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "oui", params.Oui)
    searchItems = updateSearchItemWithValue(searchItems, "product_class", params.ProductClass)
    searchItems = updateSearchItemWithValue(searchItems, "menufacture", params.Menufacture)
    searchItems = updateSearchItemWithValue(searchItems, "name", params.Name)
    searchItems = updateSearchItemWithValue(searchItems, "enable", params.Enable)

    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchProducts(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }
  const handleDeleteProduct = async (record: API.Product) => {
    setSelectedProduct(record);
    setDeleteModalVisible(true);
  };

  const handleEditProduct = (record: API.Product) => {
    setSelectedProduct(record);
    setUpdateModalVisible(true);
  };

  const moreItems = (record: API.Product): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'edit',
        icon: (<EditOutlined />),
        disabled: !access.canUpdateOMCProduct,
        label: (
          <a onClick={() => handleEditProduct(record)}>
            {intl.formatMessage({ id: 'common.edit' })}
          </a>
        ),
      },
      {
        key: 'delete',
        icon: (<DeleteOutlined />),
        disabled: !access.canDeleteOMCProduct,
        label: (
          <a onClick={() => handleDeleteProduct(record)}>
            {intl.formatMessage({ id: 'common.delete' })}
          </a>
        ),
      },
    ]
    return { items }
  }

  const columns: ProColumns[] = [
    {
      title: 'ID',
      dataIndex: 'Id',
      key: 'Id',
      search: false,
      render: (text: any, record: API.Product) => (
        <div>
          {text}
          <Dropdown menu={moreItems(record)}>
            <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
              <MoreOutlined />
            </Button>
          </Dropdown>
        </div>
      ),
    },
    {
      title: 'OUI',
      dataIndex: 'Oui',
      key: 'Oui',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "OUI",
      },
    },
    {
      title: 'Product Class',
      dataIndex: 'ProductClass',
      key: 'ProductClass',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Product Class",
      },
    },
    {
      title: 'Manufacturer',
      dataIndex: 'Manufacturer',
      key: 'Manufacturer',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Manufacturer",
      },
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
      title: 'Enable',
      dataIndex: 'Enable',
      key: 'Enable',
      valueType: 'select',
      valueEnum: {
        "1": { text: "Enable" },
        "0": { text: "Disable" },
      },
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Enable",
        defaultValue: "1",
      },
      render: (text: any, record: API.Product) => (
        <Switch checked={record.Enable} onChange={
          (checked) => {
            if (!checked) {
              setProductDisable(record)
            } else {
              setProductEnable(record)
            }
            ref.current?.reload();
          }}
        />),
    },
    {
      title: 'Created',
      dataIndex: 'Created',
      key: 'Created',
      render: (text: any, record: API.Product) => <span>{formatTimestamp2(record.Created)}</span>,
    },
    {
      title: 'Updated',
      dataIndex: 'Updated',
      key: 'Updated',
      render: (text: any, record: API.Product) => <span>{formatTimestamp2(record.Updated)}</span>,
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
              onClick={() => setCreateModalVisible(true)}
              style={{ fontSize: '20px' }}
            />
          </Space>

        ]}
      />
      <CreateProductForm
        visible={createModalVisible}
        onCancel={() => setCreateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <UpdateProductForm
        visible={updateModalVisible}
        product={selectedProduct}
        onCancel={() => setUpdateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <DeleteProductForm
        visible={deleteModalVisible}
        product={selectedProduct}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default ProductsTable;
