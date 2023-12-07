import { useState, useEffect } from 'react';
import { ProCard } from '@ant-design/pro-components';
import { Typography } from 'antd';
const { Text } = Typography;
import { AlarmStatsCard } from '../alarm';
import { DeviceStatsCard } from '../device';
import { formatTimestampWithTZ } from '@/utils/format';
import dayjs from 'dayjs';
import { useSystemTime } from '@/models/system';

const PageHeaderExtra: React.FC = () => {
  const { systemTimeDiff } = useSystemTime();
  const [currentTime, setCurrentTime] = useState(0);
  const [dateRange, setDateRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([dayjs().add(-3600, "second"), dayjs()]);

  const reload = () => {
    const now = dayjs().add(systemTimeDiff, "second");
    const ts = now.unix()
    if (ts % 60 === 0) {
      setDateRange([dateRange[0], now])
    }
    setCurrentTime(ts);
    setTimeout(reload, 1000);
  };

  useEffect(() => {
    const intervalId = setTimeout(reload, 1000);
    return () => {
      clearTimeout(intervalId);
    };
  }, []);

  return (
    <ProCard split="vertical" ghost gutter={16} >
      <ProCard ghost>
        <DeviceStatsCard dateRange={dateRange} />
      </ProCard>
      <ProCard ghost>
        <AlarmStatsCard dateRange={dateRange} />
      </ProCard>
      <ProCard ghost >
        <Text type="secondary" style={{ width: 220, display: 'block' }}>{formatTimestampWithTZ(currentTime)}</Text>
      </ProCard>
    </ProCard>
  );
};

export default PageHeaderExtra;
