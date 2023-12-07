import React, { useRef, useState, } from 'react';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Button, Dropdown, MenuProps, Space, Switch, } from 'antd';
import { DeleteOutlined, EditOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons';
import { formatTimestamp } from '@/utils/format';
import { CreateUserForm, UpdateUserForm, DeleteUserForm, SetUserRoleModal } from '@/components/iam';
import { fetchUsers, setUserEnable, setUserDisable } from '@/models/iam_users';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { useIntl } from '@umijs/max';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const UserTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  const [selectedUser, setSelectedUser] = useState<API.User | undefined>(undefined);
  const [createModalVisible, setCreateModalVisible] = useState(false);
  const [updateModalVisible, setUpdateModalVisible] = useState(false);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);
  const [setRoleModalVisibale, setSetRoleModalVisible] = useState(false);

  const handleRequest = async (params: {
    Name?: string;
    Enable?: boolean;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "name", params.Name)
    searchItems = updateSearchItemWithValue(searchItems, "enable", params.Enable)

    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchUsers(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  // Handle delete button click
  const handleDeleteUser = async (record: API.User) => {
    setSelectedUser(record);
    setDeleteModalVisible(true);
  };

  const handleEditUser = (record: API.User) => {
    setSelectedUser(record);
    setUpdateModalVisible(true);
  };

  const handleSetUserRole = (record: API.User) => {
    setSelectedUser(record);
    setSetRoleModalVisible(true);
  };

  const moreItems = (record: API.User): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'setRole',
        icon: (<EditOutlined />),
        disabled: record.Default || !access.canCreateIAMUser,
        label: (
          <a onClick={() => handleSetUserRole(record)} >
            {intl.formatMessage({ id: 'iam.user.edit-role' })}
          </a>
        ),
      },
      {
        key: 'edit',
        icon: (<EditOutlined />),
        disabled: record.Default || !access.canUpdateIAMUser,
        label: (
          <a onClick={() => handleEditUser(record)}>
            {intl.formatMessage({ id: 'common.edit' })}
          </a>
        ),
      },
      {
        key: 'delete',
        disabled: record.Default || !access.canDeleteIAMUser,
        icon: (<DeleteOutlined />),
        label: (
          <a onClick={() => handleDeleteUser(record)}>
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
      render: (text: any, record: API.User) => (
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
      title: 'Alias',
      dataIndex: 'Alias',
      key: 'Alias',
      search: false,
    },
    {
      title: 'Roles',
      dataIndex: 'roles',
      key: 'roles',
      search: false,
      render: (text: any, record: API.User) => record.Roles?.join(', '),
    },

    {
      title: 'Default',
      dataIndex: 'Default',
      key: 'Default',
      search: false,
      render: (text: any, record: API.User) => (<Switch checked={record.Default} />),

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
        defaultValue: "0",
      },
      render: (text: any, record: API.User) => (
        <Switch checked={record.Enable} onChange={
          (checked) => {
            if (checked) {
              setUserEnable(record)
            } else {
              setUserDisable(record)
            }
            ref.current?.reload();
          }}
        />),
    },
    {
      title: 'Created',
      dataIndex: 'Created',
      key: 'Created',
      render: (text: any, record: API.User) => <span>{formatTimestamp(record.Created)}</span>,
    },
    {
      title: 'Updated',
      dataIndex: 'Updated',
      key: 'Updated',
      render: (text: any, record: API.User) => <span>{formatTimestamp(record.Updated)}</span>,
    },

  ];
  return (
    <>
      <ProTable
        {...proTableLayout}
        rowKey="Id"
        columns={columns}
        params={{}}
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
      <CreateUserForm
        visible={createModalVisible}
        onCancel={() => setCreateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <UpdateUserForm
        visible={updateModalVisible}
        user={selectedUser}
        onCancel={() => setUpdateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <DeleteUserForm
        visible={deleteModalVisible}
        user={selectedUser}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <SetUserRoleModal
        visible={setRoleModalVisibale}
        user={selectedUser}
        onCancel={() => setSetRoleModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default UserTable;