import { ProCard } from '@ant-design/pro-components';
import { Badge } from 'antd';
import { useDeviceCurrentStatusValue } from '@/models/device_status';
import dayjs from 'dayjs';
import { useEffect, useState } from 'react';

type Props = {
  dateRange: [dayjs.Dayjs, dayjs.Dayjs]
};

const StatsCard: React.FC<Props> = ({ dateRange }) => {
  const { data } = useDeviceCurrentStatusValue("cpe", dateRange)
  const { data: data2 } = useDeviceCurrentStatusValue("enb", dateRange)
  const [online, setOnline] = useState<number>(0);
  const [offline, setOffline] = useState<number>(0);
  const [active, setActive] = useState<number>(0);

  useEffect(() => {
    setOnline(data.online || 0 + data2.online || 0);
    setOffline(data.offline || 0 + data2.offline || 0);
    setActive(data.active || 0 + data2.active || 0);
  }, [data, data2]);

  return (
    <ProCard split="vertical" gutter={8} ghost >
      <ProCard ghost > <Badge key="Online" color={'blue'} text={"Online: " +  online } /></ProCard>
      <ProCard ghost> <Badge key="Offline" color={'gray'} text={"Offline: " +  offline } /></ProCard>
      <ProCard ghost> <Badge key="Active" color={'green'} text={"Active: " +  active } /></ProCard>
    </ProCard>
  );
};
export default StatsCard;