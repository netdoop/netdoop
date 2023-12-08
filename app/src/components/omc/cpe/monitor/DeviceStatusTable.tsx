import React, { useEffect, useRef, useState } from 'react';
import { history, useIntl } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Dropdown, Button, MenuProps, Modal } from 'antd';
import { MoreOutlined, UnorderedListOutlined, EditOutlined, CloudSyncOutlined, ProfileOutlined } from '@ant-design/icons';


import { fetchDevices, rebootDevice, getDeviceParameterNames } from '@/models/device';
import { useGroups } from '@/models/groups';
import { FetchParams, SearchItem, updateSearchItemWithValue, updateSearchItemWithValues } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { DeviceOnlineStatusIcon, UpgradeDeviceModal, UploadDeviceFileModal } from '../../device';
import { formatTimestamp2, getDuratinon } from '@/utils/format';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

interface DeviceInfoTableProps {
  groupIds: number[],
}

const DeviceInfoTable: React.FC<DeviceInfoTableProps> = ({
  groupIds
}) => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();
  const { groupNameById } = useGroups();

  const [selectedDevice, setSelectedDevice] = useState<API.Device>();
  const [uploadFileModalVisible, setUploadFileModalVisible] = useState(false);
  const [upgradeModalVisible, setUpgradeModalVisible] = useState(false);


  useEffect(() => {
    ref.current?.reload();
  }, [groupIds]);

  const handleRequest = async (params: {
    groupIds: number[];
    SerialNumber?: string;
    IMSI?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", "cpe")
    searchItems = updateSearchItemWithValues(searchItems, "group", params.groupIds)
    searchItems = updateSearchItemWithValue(searchItems, "serial_number", params.SerialNumber)
    searchItems = updateSearchItemWithValue(searchItems, "meta_data.IMSI", params.IMSI)

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

  const handleInformation = async (record: API.Device) => {
    history.push(history.location.pathname + '/' + record.Id + '/information')
  };
  // const handleStatistics = async (record: API.Device) => {
  //   history.push(history.location.pathname + '/' + record.Id + '/statistics')
  // };
  const handleSetting = (record: API.Device) => {
    history.push(history.location.pathname + '/' + record.Id + '/setting')
  };
  const handleReboot = (record: API.Device) => {
    Modal.confirm({
      title: 'Do you want to reboot this device?',
      okText: 'Confirm',
      cancelText: 'Cancel',
      onOk: () => {
        rebootDevice(record)
      },
    });
  };
  const handleSync = (record: API.Device) => {
    getDeviceParameterNames(record, "Device.", false)
  };
  const handleUploadFile = (record: API.Device) => {
    setSelectedDevice(record);
    setUploadFileModalVisible(true);
  };
  const handleUprade = (record: API.Device) => {
    setSelectedDevice(record);
    setUpgradeModalVisible(true);
  };

  const moreItems = (record: API.Device): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'information',
        icon: (<ProfileOutlined />),
        disabled: !(access.canGetOMCDevice && access.canGetOMCDeviceParameterValues),
        label: (
          <a onClick={() => handleInformation(record)}>
            {intl.formatMessage({ id: 'common.information' })}
          </a>
        ),
      },
      {
        key: 'setting',
        icon: (<EditOutlined />),
        disabled: !(access.canGetOMCDevice && access.canGetOMCDeviceParameterValues && access.canSetOMCDeviceParameterValues),
        label: (
          <a onClick={() => handleSetting(record)}>
            {intl.formatMessage({ id: 'common.setting' })}
          </a>
        ),
      },
      {
        key: 'actions',
        icon: (<UnorderedListOutlined />),
        label: intl.formatMessage({ id: 'common.actions' }),
        disabled: (
          !access.canRebootOMCDevice &&
          !access.canGetOMCDeviceParameterNames &&
          !access.canUploadFileToOMCDevice &&
          !(access.canUpgradeOMCDevice && access.canListOMCFirmwares)
        ),
        children: [
          {
            key: 'sync',
            icon: (<CloudSyncOutlined />),
            disabled: !access.canGetOMCDeviceParameterNames,
            label: (
              <a onClick={() => handleSync(record)}>
                {intl.formatMessage({ id: 'common.actions.sync' })}
              </a>
            )
          },
          {
            key: 'reboot',
            icon: (<CloudSyncOutlined />),
            disabled: !access.canRebootOMCDevice,
            label: (
              <a onClick={() => handleReboot(record)}>
                {intl.formatMessage({ id: 'common.actions.reboot' })}
              </a>
            ),
          },
          {
            key: 'upload-file',
            icon: (<CloudSyncOutlined />),
            disabled: !access.canUploadFileToOMCDevice,
            label: (
              <a onClick={() => handleUploadFile(record)}>
                {intl.formatMessage({ id: 'common.actions.upload-file' })}
              </a>
            ),
          },
          {
            key: 'upgrade',
            icon: (<CloudSyncOutlined />),
            disabled: !(access.canUpgradeOMCDevice && access.canListOMCFirmwares),
            label: (
              <a onClick={() => handleUprade(record)}>
                {intl.formatMessage({ id: 'common.actions.upgrade' })}
              </a>
            ),
          },
        ],
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
      title: '  ',
      dataIndex: 'Online',
      key: 'Online',
      search: false,
      render: (text: any, record: API.Device) => <DeviceOnlineStatusIcon device={record} />,
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
      title: 'MAC',
      dataIndex: ['Properties', 'MACAddress'],
      key: 'MACAddress',
      search: false,
    },
    {
      title: 'IP',
      dataIndex: ['Properties', 'IPAddress'],
      key: 'IPAddress',
      search: false,
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
      title: 'Software Version',
      dataIndex: ['Properties', 'SoftwareVersion'],
      key: 'SoftwareVersion',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Software Version",
      },
    },
    {
      title: 'Group',
      dataIndex: 'Group',
      key: 'Group',
      search: false,
      render: (text: any, record: API.Device) => <span>{groupNameById(record.GroupId)}</span>,
    },

    {
      title: 'Cell ID',
      dataIndex: ['Properties', 'CellID'],
      key: 'CellID',
      search: false,
    },
    {
      title: 'PCI',
      dataIndex: ['Properties', 'PCI'],
      key: 'PCI',
      search: false,
    },
    {
      title: 'ECGI/NCGI',
      dataIndex: ['Properties', 'CGI'],
      key: 'CGI',
      search: false,
      ellipsis: true,
    },
    {
      title: 'Run Time',
      dataIndex: ['Properties', 'RunTime'],
      key: 'RunTime',
      search: false,
      render: (text: any) => getDuratinon(text),
    },
    {
      title: 'Inform Time',
      dataIndex: 'LastInformTime',
      key: 'LastInformTime',
      search: false,
      render: (text: any) => formatTimestamp2(text, 1000000),
    },
  ];

  return (
    <>
      <ProTable
        {...proTableLayout}
        rowKey="Id"
        columns={columns}
        actionRef={ref}
        params={{ groupIds: groupIds }}
        request={handleRequest}
        // dataSource={devices}
        // loading={loading}
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
      <UploadDeviceFileModal
        visible={uploadFileModalVisible}
        device={selectedDevice}
        onCancel={() => setUploadFileModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <UpgradeDeviceModal
        visible={upgradeModalVisible}
        device={selectedDevice}
        onCancel={() => setUpgradeModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default DeviceInfoTable;
