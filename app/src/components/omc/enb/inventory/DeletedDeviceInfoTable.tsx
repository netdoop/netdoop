import React, { useEffect, useRef, useState } from 'react';
import { useIntl } from '@umijs/max';
import { ActionType, ProCard, ProColumns, ProTable } from '@ant-design/pro-components';
import { Switch, Dropdown, Button, MenuProps } from 'antd';
import { MoreOutlined, DeleteOutlined, RollbackOutlined } from '@ant-design/icons';

import { useGroups } from '@/models/groups';
import { fetchDeletedDevices } from '@/models/device';
import { SortOrder } from 'antd/es/table/interface';
import { FetchParams, SearchItem, updateSearchItemWithValue, updateSearchItemWithValues } from '@/models/common';
import { DeleteDeletedDeviceModal, RecoverDeletedDeviceModal } from '../../device';
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
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);
  const [recoverModalVisible, setRecoverModalVisible] = useState(false);

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
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", "enb")
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
    const result = await fetchDeletedDevices(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleSelectGroups = (selectedKeys: React.Key[]) => {
    const ids: number[] = selectedKeys.map((key: React.Key) => parseInt(key as string));
    setGroupIds(ids)
  }

  const handlePermanentlyDelete = async (record: API.Device) => {
    setSelectedDevice(record);
    setDeleteModalVisible(true);
  };
  const handleRecover = async (record: API.Device) => {
    setSelectedDevice(record);
    setRecoverModalVisible(true);
  };
  const moreItems = (record: API.Device): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'recover',
        icon: (<RollbackOutlined />),
        disabled: !access.canRecoverOMCDeletedDevice,
        label: (
          <a onClick={() => handleRecover(record)}>
            {intl.formatMessage({ id: 'common.recover' })}
          </a>
        ),
      },
      {
        key: 'permanently-delete',
        icon: (<DeleteOutlined />),
        disabled: !access.canDeleteOMCDeletedDevice,
        label: (
          <a onClick={() => handlePermanentlyDelete(record)}>
            {intl.formatMessage({ id: 'common.permanently-delete' })}
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
        <Switch checked={record.Enable} />),
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
          />
        </ProCard>
      </ProCard>
      <DeleteDeletedDeviceModal
        visible={deleteModalVisible}
        device={selectedDevice}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <RecoverDeletedDeviceModal
        visible={recoverModalVisible}
        device={selectedDevice}
        onCancel={() => setRecoverModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default DeviceInfoTable;

