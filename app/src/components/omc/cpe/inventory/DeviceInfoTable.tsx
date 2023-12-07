import React, { useEffect, useRef, useState } from 'react';
import { history, useIntl } from '@umijs/max';
import { ActionType, ProCard, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Switch, Dropdown, Button, MenuProps } from 'antd';
import { PlusOutlined, MoreOutlined, UnorderedListOutlined, EditOutlined, ApartmentOutlined, DeleteOutlined } from '@ant-design/icons';
import { CreateDeviceModal, UpdateDeviceInfoModal, DeleteDeviceModal, SetDeviceGroupModal } from '../../device';

import { useGroups } from '@/models/groups';
import { fetchDevices, setDeviceEnable, setDeviceDisable } from '@/models/device';
import { SortOrder } from 'antd/es/table/interface';
import { FetchParams, SearchItem, updateSearchItemWithValue, updateSearchItemWithValues } from '@/models/common';
import { GroupTree } from '../../group';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const DeviceInfoTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  const { groupNameById } = useGroups();
  const [groupIds, setGroupIds] = useState<number[]>([]);

  const [selectedDevice, setSelectedDevice] = useState<API.Device>();
  const [createModalVisible, setCreateModalVisible] = useState(false);
  const [updateModalVisible, setUpdateModalVisible] = useState(false);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);
  const [setGroupModalVisible, setSetGroupModalVisible] = useState(false);

  useEffect(() => {
    ref.current?.reload();
  }, [groupIds]);

  const handleRequest = async (params: {
    groupIds?: number[];
    SerialNumber?: string;
    IMSI?: string;
    Enable?: boolean;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    // }, sort: Record<string, SortOrder>, filter: Record<string, (string | number)[] | null>) => {
    // console.log(params, sort, filter)
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", "cpe")
    searchItems = updateSearchItemWithValues(searchItems, "group", params.groupIds)
    searchItems = updateSearchItemWithValue(searchItems, "serial_number", params.SerialNumber)
    searchItems = updateSearchItemWithValue(searchItems, "meta_data.IMSI", params.IMSI)
    searchItems = updateSearchItemWithValue(searchItems, "enable", params.Enable)

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


  const handleSelectGroups = (selectedKeys: React.Key[]) => {
    const ids: number[] = selectedKeys.map((key: React.Key) => parseInt(key as string));
    setGroupIds(ids)
  }

  const handleInformation = async (record: API.Device) => {
    history.push(history.location.pathname + '/' + record.Id + '/information')
  };
  const handleSetting = (record: API.Device) => {
    setSelectedDevice(record);
    setUpdateModalVisible(true);
  };
  const handleDelete = async (record: API.Device) => {
    setSelectedDevice(record);
    setDeleteModalVisible(true);
  };
  const handleSetGroup = (record: API.Device) => {
    setSelectedDevice(record);
    setSetGroupModalVisible(true);
  };
  const moreItems = (record: API.Device): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'information',
        icon: (<UnorderedListOutlined />),
        disabled: !access.canGetOMCDevice,
        label: (
          <a onClick={() => handleInformation(record)}>
            {intl.formatMessage({ id: 'common.information' })}
          </a>
        ),
      },
      {
        key: 'setting',
        icon: (<EditOutlined />),
        disabled: !access.canUpdateOMCDevice,
        label: (
          <a onClick={() => handleSetting(record)}>
            {intl.formatMessage({ id: 'common.setting' })}
          </a>
        ),
      },
      {
        key: 'move-to-group',
        icon: (<ApartmentOutlined />),
        disabled: !(access.canSetOMCDeviceGroup && access.canListOMCGroups),
        label: (
          <a onClick={() => handleSetGroup(record)}>
            {intl.formatMessage({ id: 'omc.device.move-to-group' })}
          </a>
        ),
      },
      {
        key: 'delete',
        icon: (<DeleteOutlined />),
        disabled: !access.canDeleteOMCDevice,
        label: (
          <a onClick={() => handleDelete(record)}>
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
      fixed: 'left' as 'left',
      width: 100,
      render: (text: any, record: API.Device) => (
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
      title: 'MAC',
      dataIndex: ['Properties', 'MACAddress'],
      key: 'MACAddress',
      search: false,
    },
    {
      title: 'IMSI',
      dataIndex: ['MetaData', 'IMSI'],
      key: 'IMSI',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "IMSI",
      },
    },
    {
      title: 'Latitude',
      dataIndex: ['Properties', 'Latitude'],
      key: 'Latitude',
      search: false,
    },
    {
      title: 'Longitude',
      dataIndex: ['Properties', 'Longitude'],
      key: 'Longitude',
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
      render: (text: any, record: API.Device) => (
        <Switch checked={record.Enable} onChange={
          (checked) => {
            if (!checked) {
              setDeviceDisable(record);
            } else {
              setDeviceEnable(record);
            }
            ref.current?.reload();
          }}
        />),
    },
  ];


  return (
    <>
      <ProCard gutter={4} style={{}}>
        <ProCard colSpan={4} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
          <GroupTree onSelect={handleSelectGroups} />
        </ProCard>
        <ProCard colSpan={20} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
          <ProTable
            {...proTableLayout}
            rowKey="Id"
            columns={columns}
            params={{ groupIds: groupIds }}
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
                <PlusOutlined
                  className="ant-pro-table-toolbar-item-iconButton"
                  onClick={() => setCreateModalVisible(true)}
                  style={{ fontSize: '12px' }}
                />
              </Space>
            ]}
          />
        </ProCard>
      </ProCard>

      <CreateDeviceModal
        visible={createModalVisible}
        onCancel={() => setCreateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <UpdateDeviceInfoModal
        visible={updateModalVisible}
        device={selectedDevice}
        onCancel={() => setUpdateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <SetDeviceGroupModal
        visible={setGroupModalVisible}
        onCancel={() => setSetGroupModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
        device={selectedDevice}
      />
      <DeleteDeviceModal
        visible={deleteModalVisible}
        device={selectedDevice}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default DeviceInfoTable;

