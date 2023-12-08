import { Column, ColumnConfig } from '@ant-design/charts';
import { useDeviceStatusValues } from '@/models/device_status';
import dayjs from 'dayjs';

type DeviceStatusStackedProps = {
  productType: string,
  dateRange: [dayjs.Dayjs, dayjs.Dayjs],
  interval?: number,
  width?: number,
  height: number,
};

const DeviceStatusStacked: React.FC<DeviceStatusStackedProps> = ({ productType, dateRange, interval, width, height }) => {
  const { data, max } = useDeviceStatusValues(productType, dateRange, interval);
  const config: ColumnConfig = {
    data,
    autoFit: width === undefined,
    width: width,
    height: height,
    xField: "timeText",
    yField: "value",
    seriesField: "name",
    tooltip: {
      shared: true,
      showMarkers: false,
    },
    isStack: true,
    interactions: [{ type: 'active-region' }],
    padding: "auto",
    color: ['#8c8c8c', '#ffc53d', '#73d13d','#4096ff'],
    meta: {
      value: {
        min: 0,
        max: max > 0 ? Math.ceil(max * 1.2) : undefined,
      },
      name: {
       values: ['offline','inactive',  'online', 'active'], 
      },
    }
  }
  // const onReady =async (chart:any) => {
  // console.log(chart)
  // }
  return <Column
    {...config}
  //  onReady={onReady} 
  />;
};

export default DeviceStatusStacked;
