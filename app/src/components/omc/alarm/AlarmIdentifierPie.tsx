import { Datum, Pie, PieConfig } from '@ant-design/charts';
import { useAlarmIdentifierStats } from '@/models/alarm';
import { useEffect, useState } from 'react';
import dayjs from 'dayjs';

type Props = {
  dateRange: [dayjs.Dayjs, dayjs.Dayjs],
  height: number,
  radius: number,
};

type item = {
  name: string,
  value: number,
}

const AlarmIdentifierPie: React.FC<Props> = ({ dateRange, height, radius }) => {
  const { data, loading } = useAlarmIdentifierStats(dateRange);
  const [items, setItems] = useState<item[]>([]);

  useEffect(() => {
    let tmp: item[] = []
    for (const key in data) {
      if (data.hasOwnProperty(key)) {
        tmp.push({ name: key, value: data[key] })
      }
    }
    setItems(tmp)
  }, [data]);

  const config: PieConfig = {
    loading:loading,
    data: items,
    autoFit: true,
    radius: radius,
    innerRadius: 0.372,
    angleField: 'value',
    colorField: 'name',
    theme: 'light',
    // color:  ['#e8c1a0', '#f47560', '#f1e15b', '#e8a838', '#61cdbb'],
    pieStyle: {
      stroke: '#FFF',
      lineWidth: 2,
      fontWeight: 'bold',
      textBaseline: 'bottom',
    },
    label: {
      fields: ['name'],
      type: 'spider',
      content: (data: Datum) => { return data.name },
    },

    statistic: {
      title: false,
      content: false,
    }
  }
  return (
    <Pie style={{height:height}}
      {...config}
    />
  );
};
export default AlarmIdentifierPie;