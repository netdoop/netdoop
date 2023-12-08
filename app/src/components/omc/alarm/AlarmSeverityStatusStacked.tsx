import { Column, ColumnConfig } from '@ant-design/charts';
import { useDeviceAlarmStatusValues, SEVERITY_COLOR_MAP } from '@/models/alarm';
import dayjs from 'dayjs';

type DeviceStatusStackedProps = {
  dateRange: [dayjs.Dayjs, dayjs.Dayjs],
  interval?: number,
  width?: number,
  height: number,
};

const DeviceStatusStacked: React.FC<DeviceStatusStackedProps> = ({ dateRange, interval, width, height }) => {
  const { data, max } = useDeviceAlarmStatusValues("", dateRange, interval);
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
    color: [
      SEVERITY_COLOR_MAP.Warning,
      SEVERITY_COLOR_MAP.Minor,
      SEVERITY_COLOR_MAP.Major,
      SEVERITY_COLOR_MAP.Critical,
    ],
    meta: {
      value: {
        min: 0,
        max: max > 0 ? Math.ceil(max * 1.2) : undefined,
      },
      name: {
        values: ['Warning', 'Minor', 'Major', 'Critical'],
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
