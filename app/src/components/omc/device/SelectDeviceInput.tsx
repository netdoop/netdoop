import React, { useEffect, useRef, useState } from 'react';
import { ActionType, ProCard, ProColumns, ProTable } from '@ant-design/pro-components';

import { fetchDevices } from '@/models/device';
import { SortOrder } from 'antd/es/table/interface';
import { FetchParams, SearchItem, updateSearchItemWithValue, updateSearchItemWithValues } from '@/models/common';
import { GroupTree } from '../group';
import { proTableLayout } from '@/constants/style';

interface Props {
  productType: string,
  value?: React.Key[];
  onChange?: (value: React.Key[]) => void;
};

const SelectDeviceInput: React.FC<Props> = ({
  productType,
  value,
  onChange,
}) => {
  const ref = useRef<ActionType>();
  const [groupIds, setGroupIds] = useState<number[]>([]);

  useEffect(() => {
    ref.current?.reload();
  }, [groupIds, productType]);

  const handleRequest = async (params: {
    groupIds?: number[];
    SerialNumber?: string;
    Name?: string,
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
    searchItems = updateSearchItemWithValue(searchItems, "name", params.Name)
    searchItems = updateSearchItemWithValues(searchItems, "group", params.groupIds)
    searchItems = updateSearchItemWithValue(searchItems, "serial_number", params.SerialNumber)
    searchItems = updateSearchItemWithValue(searchItems, "enable", true)

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
    const GroupIds: number[] = selectedKeys.map((key: React.Key) => parseInt(key as string));
    setGroupIds(GroupIds)
  }

  const rowSelection = {
    onChange: (selectedRowKeys: React.Key[], selectedRows: API.Device[]) => {
      if (onChange) {
        onChange(selectedRowKeys)
      }
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record: API.Device) => ({
      name: record.Name,
    }),
  };

  const columns: ProColumns[] = [
    {
      title: 'ID',
      dataIndex: 'Id',
      key: 'Id',
      search: false,
      fixed: 'left' as 'left',
      width: 100,
    },
    {
      title: 'Cell Name',
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
  ];

  return (
    <ProCard gutter={4} split="vertical">
      <ProCard colSpan={4} bodyStyle={{ paddingInline: 2 }}>
        <GroupTree onSelect={handleSelectGroups} />
      </ProCard>
      <ProCard colSpan={20} bodyStyle={{ paddingInline: 2 }}>
        <ProTable
          {...proTableLayout}
          rowKey="Id"
          columns={columns}
          rowSelection={{
            type: "checkbox",
            selectedRowKeys: value,
            ...rowSelection,
          }}
          params={{ groupIds: groupIds }}
          request={handleRequest}
          actionRef={ref}
          search={{
            span: 8,
            labelWidth: 0,
          }}
          scroll={{ x: 'max-content' }}
          options={{
            density: false,
            fullScreen: false,
            setting: false,
            reload: false,
          }}
        />
      </ProCard>
    </ProCard>
  );
};

export default SelectDeviceInput;

