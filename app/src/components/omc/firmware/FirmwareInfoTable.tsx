import React, { useRef, useState } from 'react';
import { useIntl, request } from '@umijs/max';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Button, Dropdown, MenuProps, Space } from 'antd';
import { DeleteOutlined, DownloadOutlined, MoreOutlined, PlusOutlined } from '@ant-design/icons';
import DeleteFirmwareModal from './DeleteFirmwareModal';
import CreateFirmwareModal from './CreateFirmwareModal';
import { fetchFirmwares } from '@/models/firmware';
import { SearchItem, FetchParams, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { formatBytes } from '@/utils/format';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

interface Props {
  productType: string;
};

const FirmwareInfoTable: React.FC<Props> = ({ productType }) => {
  const access = useAccess()
  const intl = useIntl()
  const ref = useRef<ActionType>();

  const [selectedFirmware, setSelectedFirmware] = useState<API.Firmware>();
  const [createModalVisible, setCreateModalVisible] = useState(false);
  const [deleteModalVisible, setDeleteModalVisible] = useState(false);

  const handleRequest = async (params: {
    Version?: string,
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
    searchItems = updateSearchItemWithValue(searchItems, "version", params.Version)

    sort['ID'] = "descend"
    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchFirmwares(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const handleDelete = async (record: API.Firmware) => {
    setSelectedFirmware(record);
    setDeleteModalVisible(true);
  };
  const handleDownload = async (record: API.Firmware) => {
    if (record.S3Object?.DownloadUrl) {
      const response = await request(record.S3Object?.DownloadUrl, {
        method: 'GET',
        responseType: 'blob', // set response type to blob to receive file data
      });
      // create a new Blob object from the file data
      const file = new Blob([response]);
      // create a URL for the file using the createObjectURL function
      const fileUrl = URL.createObjectURL(file);

      // create a link element with the URL and click it to initiate download
      const link = document.createElement('a');
      link.href = fileUrl;
      link.setAttribute('download', record.S3Object?.FileName || 'firmware'); // set filename here
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link)
    }
  };

  const moreItems = (record: API.Firmware): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'download',
        icon: (<DownloadOutlined />),
        disabled: !access.canGetOMCFirmware,
        label: (
            <a onClick={() => handleDownload(record)}>
              {intl.formatMessage({ id: 'common.download' })}
            </a>
        ),
      },
      {
        key: 'delete',
        icon: (<DeleteOutlined />),
        disabled: !access.canDeleteOMCFirmware,
        label: (
            <a onClick={() => handleDelete(record)}>
              {intl.formatMessage({ id: 'common.delete' })}
            </a>
        ),
      },
    ];
    return { items };
  };

  const columns: ProColumns[] = [
    {
      title: 'ID',
      dataIndex: 'Id',
      key: 'Id',
      search: false,
      width: 100,
      fixed: 'left' as 'left',
      render: (text: any, record: API.Firmware) => (
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
      title: 'Version',
      dataIndex: 'Version',
      key: 'Version',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Version",
      },
    },

    {
      title: 'File Name',
      dataIndex: ["S3Object", 'FileName'],
      key: 'FileName',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "File Name",
      },
    },
    {
      title: 'File Size',
      dataIndex: ['S3Object', "FileSize"],
      key: 'File Size',
      search: false,
      render: (text: any, record: API.Firmware) => (
        <span>{formatBytes(record.S3Object?.FileSize || -1)}</span>
      ),
    },
    {
      title: 'Uploader',
      dataIndex: 'Uploader',
      key: 'Uploader',
      search: false,
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
        scroll={{ x: 'max-content' }}
        options={{
          density: false,
          fullScreen: true,
          setting: true,
          reload: true
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
      <CreateFirmwareModal
        productType={productType}
        visible={createModalVisible}
        onCancel={() => setCreateModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
      <DeleteFirmwareModal
        visible={deleteModalVisible}
        firmware={selectedFirmware}
        onCancel={() => setDeleteModalVisible(false)}
        onSuccess={() => ref.current?.reload()}
      />
    </>
  );
};

export default FirmwareInfoTable
