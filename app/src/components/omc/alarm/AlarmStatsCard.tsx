import { ProCard } from '@ant-design/pro-components';
import { useAlarmStats } from '@/models/alarm';
import { Badge } from 'antd';
import { SEVERITY_COLOR_MAP } from '@/models/alarm';
import dayjs from 'dayjs';
import { useEffect, useState } from 'react';

type Props = {
  dateRange: [dayjs.Dayjs, dayjs.Dayjs],
};

const AlarmStatsCard: React.FC<Props> = ({ dateRange }) => {
  const { data } = useAlarmStats(dateRange)
  const [critical, setCritical] = useState<number>(0);
  const [major, setMajor] = useState<number>(0);
  const [minor, setMinor] = useState<number>(0);
  const [warning, setWarning] = useState<number>(0);

  useEffect(() => {
    setCritical(data.Critical || 0);
    setMajor(data.Major || 0);
    setMinor(data.Minor || 0);
    setWarning(data.Warning || 0);
  }, [data]);

  return (
    <ProCard split="vertical" gutter={8} ghost  >
      <ProCard ghost > <Badge key="Critical" color={SEVERITY_COLOR_MAP.Critical} text={"Critical: " + critical} /></ProCard>
      <ProCard ghost> <Badge key="Major" color={SEVERITY_COLOR_MAP.Major} text={"Major: " + major} /></ProCard>
      <ProCard ghost> <Badge key="Minor" color={SEVERITY_COLOR_MAP.Minor} text={"Minor: " + minor} /></ProCard>
      <ProCard ghost> <Badge key="Warning" color={SEVERITY_COLOR_MAP.Warning} text={"Warning: " + warning} /></ProCard>
    </ProCard>
  );
};
export default AlarmStatsCard;