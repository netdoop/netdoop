import { Line, LineConfig } from '@ant-design/charts';
import { fetchPmValueData } from '@/models/device_pm';
import { useEffect, useState } from 'react';
import { SearchItem, updateSearchItemWithValue } from '@/models/common';
import dayjs from 'dayjs';

type Props = {
  names: string[],
  productType: string,
  device?: API.Device,
  dateRange: [dayjs.Dayjs, dayjs.Dayjs],
  interval?: number,
  width?: number,
  height: number,
  unit?: string,
  func?: 'sum' | 'increase' | 'last',
};

const DevicePmValueChart: React.FC<Props> = ({ 
  names, productType, device, dateRange, interval, width, height, unit, func = 'last' 
}) => {
  const [loading, setLoading] = useState<boolean>(false);
  const [data, setData] = useState<Record<string, any>[]>([]);

  const fetchData = async () => {
    const groupBy = ["name", "time", "value"];
    // const names = ['TotalUpload', 'TotalDownload'];
    setLoading(true)
    try {
      let items: Record<string, any>[] = [];
        await Promise.all(names.map(async (name) => {
          let searchItems: SearchItem[] = [];
          searchItems = updateSearchItemWithValue(searchItems, "product_type", productType);
          searchItems = updateSearchItemWithValue(searchItems, "name", name);
          if (device) {
            searchItems = updateSearchItemWithValue(searchItems, "device_id", device.Id);
          }
          const _items = await fetchPmValueData(dateRange, interval, searchItems, groupBy, func);
          items.push(..._items);
        }));
      setData(items);

    } catch (error) {
      setData([])
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData()
  }, [names, device, dateRange, interval])


  const config: LineConfig = {
    loading: loading,
    data: data || [],
    autoFit: width === undefined,
    width: width,
    height: height,
    xField: "timeText",
    yField: 'value',
    yAxis: {
      label: {
        formatter: (value) => `${value} ${unit||''}`,
      },
    },
    seriesField: 'name',
    tooltip: {
      shared: true,
      showMarkers: false,
    },
    interactions: [{ type: 'active-region' }],
  }
  // const onReady =async (chart:any) => {
  // console.log(chart)
  // }
  return <Line
    {...config}
  //  onReady={onReady} 
  />;
};

export default DevicePmValueChart;
