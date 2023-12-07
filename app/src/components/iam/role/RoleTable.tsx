
import React, { useRef, useState } from 'react';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Switch, Space, Button, Dropdown, MenuProps } from 'antd';
import { DeleteOutlined, EditOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons';
import { formatTimestamp } from '@/utils/format';
import { CreateRoleForm, UpdateRoleForm, DeleteRoleForm, EditRoleRulesModal } from '@/components/iam';
import { fetchRoles, setRoleDisable, setRoleEnable } from '@/models/iam_roles';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { useIntl } from '@umijs/max';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const RoleTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl();
  const ref = useRef<ActionType>();

  const [selectedRole, setSelectedRole] = useState<API.Role | undefined>(undefined);
  const [createModalVisible, setCreateModalVisible] = useState(false);
  const [updateModalVisible, setUpdateModalVisible] = useState(false);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);
  const [editRulesModalVisible, setEditRulesModalVisible] = useState(false);

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
    const result = await fetchRoles(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }


  // Handle delete button click
  const handleDeleteRole = async (record: API.Role) => {
    setSelectedRole(record);
    setDeleteModalVisible(true);
  };

  const handleEditRole = (record: API.Role) => {
    setSelectedRole(record);
    setUpdateModalVisible(true);
  };

  const handleEditRules = (record: API.Role) => {
    setSelectedRole(record);
    setEditRulesModalVisible(true);
  };

  const moreItems = (record: API.Role): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'edit-rules',
        disabled: record.Default || !access.canCreateIAMRole,
        icon: (<EditOutlined />),
        label: (
          <a onClick={() => handleEditRules(record)}>
            {intl.formatMessage({ id: 'iam.role.edit-rules' })}
          </a>
        ),
      },
      {
        key: 'edit',
        disabled: record.Default || !access.canUpdateIAMRoles,
        icon: (<EditOutlined />),
        label: (
          <a onClick={() => handleEditRole(record)}>
            {intl.formatMessage({ id: 'common.edit' })}
          </a>
        ),
      },
      {
        key: 'delete',
        disabled: record.Default || !access.canDeleteIAMRole,
        icon: (<DeleteOutlined />),
        label: (
          <a onClick={() => handleDeleteRole(record)}>
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
      render: (text: any, record: API.Role) => (
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
      title: 'Default',
      dataIndex: 'Default',
      key: 'Default',
      search: false,
      render: (text: any, record: API.Role) => (<Switch checked={record.Default} />),
    },
    {
      title: 'Enable',
      dataIndex: 'Enable',
      key: 'Enable',
      valueType: 'select',
      valueEnum: {
        "1": { text: "Enable" },
        "0": { text: "Disabled" },
      },
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Enable",
        defaultValue: "1",
      },
      render: (text: any, record: API.Role) => (
        <Switch checked={record.Enable} onChange={
          (checked) => {
            if (checked) {
              setRoleEnable(record);
            } else {
              setRoleDisable(record)
            }
            ref.current?.reload()
          }}
        />),
    },
    {
      title: 'Created',
      dataIndex: 'Created',
      key: 'Created',
      search: false,
      render: (text: any, record: API.Role) => <span>{formatTimestamp(record.Created)}</span>,
    },
    {
      title: 'Updated',
      dataIndex: 'Updated',
      key: 'Updated',
      search: false,
      render: (text: any, record: API.Role) => <span>{formatTimestamp(record.Updated)}</span>,
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
      <CreateRoleForm
        visible={createModalVisible}
        onCancel={() => setCreateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <UpdateRoleForm
        visible={updateModalVisible}
        role={selectedRole}
        onCancel={() => setUpdateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <DeleteRoleForm
        visible={deleteModalVisible}
        role={selectedRole}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <EditRoleRulesModal
        visible={editRulesModalVisible}
        role={selectedRole}
        onCancel={() => setEditRulesModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default RoleTable;