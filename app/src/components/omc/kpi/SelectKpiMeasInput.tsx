import React, { useEffect, useRef, useState } from 'react';
import { ActionType, ProCard, ProColumns, ProTable } from '@ant-design/pro-components';

import { FetchParams, SearchItem, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { fetchKPIMeasures } from '@/models/kpi_meas';
import { KpiMeasSetTree } from '.';
import { proTableLayout } from '@/constants/style';

interface Props {
  value?: React.Key[];
  onChange?: (value: React.Key[]) => void;
}


const SelectKpiMeasInput: React.FC<Props> = ({
  value,
  onChange,
}) => {
  const ref = useRef<ActionType>();

  const [measTypeSet, setMeasTypeSet] = useState<string>('Customize');

  const handleSelectSet = (selectedKeys: React.Key[]) => {
    const sets: string[] = selectedKeys.map((key: React.Key) => (key as string));
    setMeasTypeSet(sets[0])
  }

  useEffect(() => {
    ref.current?.reload();
  }, [measTypeSet]);

  const handleRequest = async (params: {
    MeasTypeID?: string;
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    let searchItems: SearchItem[] = []
    searchItems = updateSearchItemWithValue(searchItems, "meas_type_id", params?.MeasTypeID)
    if (measTypeSet && measTypeSet !== "All") {
      searchItems = updateSearchItemWithValue(searchItems, "meas_type_set", measTypeSet)
    }
    searchItems = updateSearchItemWithValue(searchItems, "enable", true)
    sort["name"] = "ascend"
    const fetchParams: FetchParams = {
      pageSize: params.pageSize,
      current: params.current,
      sort: sort,
      searchItems: searchItems,
    };
    const result = await fetchKPIMeasures(fetchParams);
    const success = true;
    const data = result.Data;
    const total = result.Total;
    return { data, success, total };
  }

  const rowSelection = {
    onChange: (selectedRowKeys: React.Key[], selectedRows: API.KPIMeas[]) => {
      if (onChange) {
        onChange(selectedRowKeys)
      }
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record: API.KPIMeas) => ({
      name: record.Name,
    }),
  };

  const columns: ProColumns[] = [
    {
      title: 'Measure Type ID',
      dataIndex: 'MeasTypeID',
      key: 'MeasTypeID',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Measure Type ID",
      },
    },
    {
      title: 'Measure Type Set',
      dataIndex: 'MeasTypeSet',
      key: 'MeasTypeSet',
      search: false,
    },
    {
      title: 'Name',
      dataIndex: 'Name',
      key: 'Name',
      search: false,
    },
    {
      title: 'Unit',
      dataIndex: 'Unit',
      key: 'Unit',
      search: false,
    },
  ];


  return (
    <ProCard gutter={4} style={{ minHeight: 320 }}>
      <ProCard colSpan={4} bodyStyle={{ paddingInline: 2 }}>
        <KpiMeasSetTree onSelect={handleSelectSet} />
      </ProCard>
      <ProCard colSpan={20} bodyStyle={{ paddingInline: 2 }}>
        <ProTable
          {...proTableLayout}
          rowKey="MeasTypeID"
          params={{}}
          request={handleRequest}
          columns={columns}
          rowSelection={{
            type: "checkbox",
            selectedRowKeys: value,
            ...rowSelection,
          }}
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

export default SelectKpiMeasInput;