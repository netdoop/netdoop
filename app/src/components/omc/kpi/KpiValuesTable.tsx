import React, { useEffect, useRef, useState } from 'react';
import { ActionType, ProCard, ProColumns, ProTable } from '@ant-design/pro-components';
import { formatTimeString } from '@/utils/format';

import { FetchParams, SearchItem, updateSearchItemWithRangeValue, updateSearchItemWithValue } from '@/models/common';
import { fetchKpiValuesByTemp } from '@/models/kpi_temp';
import { KpiTemplateTree } from '.';
import { SortOrder } from 'antd/es/table/interface';
import { proTableLayout } from '@/constants/style';
import { DateRangePick } from '../common';
import dayjs, { Dayjs } from 'dayjs';

const KpiValuesTable: React.FC = () => {
  const ref = useRef<ActionType>();

  const [template, setTemplate] = useState<API.KPITemplate | undefined>(undefined);
  const [columns, setColumns] = useState<ProColumns[]>([])

  const [timeRange, setTimeRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>();

  const handleTimeRangeChange = (value: [Dayjs, Dayjs]) => {
    setTimeRange(value);
  }

  useEffect(() => {
    if (template) {
      let update: ProColumns[] = [
        {
          title: 'Time',
          dataIndex: 'time',
          key: 'time',
          render: (text: any) => (
            <>{formatTimeString(text)}</>
          ),
          formItemProps: {
            label: "",
          },
          renderFormItem: () => (
            <DateRangePick onChange={handleTimeRangeChange} />
          )
        },
        {
          title: 'Serial Number',
          dataIndex: 'serial_number',
          key: 'serial_number',
          valueType: 'text',
          formItemProps: {
            label: "",
          },
          fieldProps: {
            placeholder: "Serial Number",
          },
        },
      ];
      template.MeasTypeIds?.forEach(v => {
        update.push({
          title: v,
          dataIndex: v.replaceAll(".", "_").toLowerCase(),
          key: v.replaceAll(".", "_").toLowerCase(),
          search: false,
          align: 'center',
        });
      })
      setColumns(update)
    } else {
      setColumns([])
    }
    ref.current?.reload();
  }, [template, timeRange]);

  const handleRequest = async (params: {
    serial_number: number | undefined,
  } & {
    pageSize?: number | undefined;
    current?: number | undefined;
    keyword?: string | undefined;
  }, sort: Record<string, SortOrder>) => {
    if (template && timeRange) {
      const from = timeRange[0];
      const to = timeRange[1];

      let searchItems: SearchItem[] = []
      searchItems = updateSearchItemWithValue(searchItems, "product_type", template.ProductType)
      searchItems = updateSearchItemWithValue(searchItems, "serial_number", params.serial_number)
      searchItems = updateSearchItemWithRangeValue(searchItems, "time", from.format("YYYY-MM-DDTHH:mm:ssZ"), to.format("YYYY-MM-DDTHH:mm:ssZ"))

      const fetchParams: FetchParams = {
        pageSize: params.pageSize,
        current: params.current,
        sort: sort,
        searchItems: searchItems,
      };
      const result = await fetchKpiValuesByTemp(template, fetchParams);
      const success = true;
      const data = result.Data;
      const total = result.Total;
      return { data, success, total };
    }
    return { data: [], success: true, total: 0 };
  }

  function handleSelect(template: API.KPITemplate | undefined): void {
    setTemplate(template)
  }

  return (
    <ProCard gutter={4} style={{}}>
      <ProCard colSpan={4} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
        <KpiTemplateTree onSelect={handleSelect} />
      </ProCard>
      <ProCard colSpan={20} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
        <ProTable
          {...proTableLayout}
          rowKey="time"
          request={handleRequest}
          columns={columns}
          actionRef={ref}
          search={{
            span: 8,
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
  );
};

export default KpiValuesTable;

