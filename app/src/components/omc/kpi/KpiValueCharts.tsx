import React, { useEffect, useState } from 'react';
import { ProCard } from '@ant-design/pro-components';

import { useAllKpiValuesByTemp } from '@/models/kpi_temp';
import { KpiTemplateTree } from '.';
import { DateRangePick } from '../common';
import dayjs, { Dayjs } from 'dayjs';
import { Line } from '@ant-design/charts';
import { timestampToString } from '@/utils/format';
import { caculateInterval } from '@/models/data';


const KpiValueCharts: React.FC = () => {

  const [template, setTemplate] = useState<API.KPITemplate | undefined>(undefined);
  const [timeRange, setTimeRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([dayjs(), dayjs()]);
  const [interval, setInterval] = useState<number>(900)
  const { data, reload, loading } = useAllKpiValuesByTemp(timeRange, template);

  const handleTimeRangeChange = (value: [Dayjs, Dayjs]) => {
    setTimeRange(value);
  }
  function handleSelect(template: API.KPITemplate | undefined): void {
    setTemplate(template)
  }

  useEffect(() => {
    setInterval(caculateInterval(timeRange))
    reload()
  }, [template, timeRange]);


  return (
    <ProCard gutter={4} style={{}}>
      <ProCard colSpan={4} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
        <KpiTemplateTree onSelect={handleSelect} />
      </ProCard>
      <ProCard colSpan={20} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
        <ProCard gutter={4} split="horizontal">
          <ProCard>
            <DateRangePick onChange={handleTimeRangeChange} />
          </ProCard>

          {template?.MeasTypeIds?.map((v) => (
            <ProCard key={v} title={v}>
              <Line
                loading={loading}
                data={data[v] || []}
                autoFit
                height={320}
                xField="time"
                xAxis={{
                  label: {
                    formatter: (value) => `${timestampToString(dayjs(value).unix(), interval)}`,
                  },
                }}
                yField="value"
                yAxis={{
                  label: {
                    formatter: (value) => `${value}`,
                  },
                }}
                seriesField="serial_number"
                tooltip={{
                  shared: true,
                  showMarkers: false,
                }}
                interactions={[{ type: 'active-region' }]}
              />
            </ProCard>
          ))}
        </ProCard>
      </ProCard>
    </ProCard>

  );
};

export default KpiValueCharts;

