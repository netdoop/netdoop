import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { message } from 'antd';
import { FetchParams, getSearchItemsString, getOrderByString } from './common';

export const fetchDeviceAlarms = async (params: FetchParams): Promise<API.listDeviceAlarmsData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDeviceAlarms.listDeviceAlarms({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch device alarms');
    return { Data: [], Total: 0 };
  }
};


export const useDeviceAlarms = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [alarms, setAlarms] = useState<API.Device[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchDeviceAlarms(params)
      setTotal(result.Total || 0)
      setAlarms(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch alarms');
    } finally {
      setLoading(false);
    }
  };
  useEffect(() => {
    reload();
  }, [params]);
  return { alarms, total, loading, reload };
};

export const useDeviceAlarm = (ts: number) => {
  const access = useAccess()

  const [alarm, setAlarm] = useState<API.DeviceAlarm | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchAlarm = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const alarm = await services.omcDeviceAlarms.getDeviceAlarm({ ts: ts });
        setAlarm(alarm);
      }
    } catch (error) {
      console.error('Failed to fetch alarm:', error);
      setAlarm(undefined);
    } finally {
      setLoading(false);
    }
  };


  useEffect(() => {
    fetchAlarm();
  }, [ts]);

  return { alarm, loading };
};

