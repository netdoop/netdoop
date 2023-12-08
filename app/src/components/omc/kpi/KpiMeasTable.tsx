import React, { useEffect, useRef, useState } from 'react';
import { history, useIntl } from '@umijs/max';
import { ActionType, ProCard, ProColumns, ProTable } from '@ant-design/pro-components';
import { Space, Dropdown, Button, MenuProps, Switch } from 'antd';
import { MoreOutlined } from '@ant-design/icons';
import { formatTimestamp2 } from '@/utils/format';

import { FetchParams, SearchItem, updateSearchItemWithValue } from '@/models/common';
import { SortOrder } from 'antd/es/table/interface';
import { fetchKPIMeasures } from '@/models/kpi_meas';
import { KpiMeasSetTree } from '.';
import { proTableLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';

const KpiMeasTable: React.FC = () => {
  const access = useAccess()
  const intl = useIntl()
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
    Enable?: boolean;
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
    searchItems = updateSearchItemWithValue(searchItems, "enable", params?.Enable)
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

  const handleDetail = async (record: API.KPIMeas) => {
    history.push(history.location.pathname + '/' + record.Id + '/detail')
  };

  // Define the actions for the table
  const moreItems = (record: API.KPIMeas): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'detail',
        disabled: !access.canGetOMCKPIMeasure,
        label: (
            <a onClick={() => handleDetail(record)}>
              {intl.formatMessage({ id: 'common.detail' })}
            </a>
        ),
      },
    ]
    return { items }
  }

  const columns: ProColumns[] = [
    {
      title: 'Measure Type ID',
      dataIndex: 'MeasTypeID',
      key: 'MeasTypeID',
      search: false,
      render: (text: any, record: API.KPIMeas) => (
        <div>
          {record.MeasTypeID}
          <Dropdown menu={moreItems(record)}>
            <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
              <MoreOutlined />
            </Button>
          </Dropdown>
        </div>
      ),
    },
    {
      title: 'Measure Type Set',
      dataIndex: 'MeasTypeSet',
      key: 'MeasTypeSet',
      valueType: 'text',
      formItemProps: {
        label: "",
      },
      fieldProps: {
        placeholder: "Measure Type Set",
        defaultValue: "",
      },
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
      render: (text: any, record: API.Product) => (<Switch checked={record.Enable} />),
    },
    {
      title: 'Default',
      dataIndex: 'Default',
      key: 'Default',
      search: false,
      render: (text: any, record: API.Product) => (<Switch checked={record.Default} />),
    },
    {
      title: 'Update Time',
      dataIndex: 'Updated',
      key: 'Updated',
      search: false,
      render: (text: any, record: API.KPIMeas) => formatTimestamp2(record.Updated),
    },
  ];


  return (
    <ProCard gutter={4} style={{}}>
      <ProCard colSpan={4} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
        <KpiMeasSetTree onSelect={handleSelectSet} />
      </ProCard>
      <ProCard colSpan={20} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
        <ProTable
          {...proTableLayout}
          rowKey="Id"
          params={{}}
          request={handleRequest}
          columns={columns}
          actionRef={ref}
          search={{
            span: 6,
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
              {/* TODO: add create alarm button */}
            </Space>]}
        />    </ProCard>
    </ProCard>
  );
};

export default KpiMeasTable;