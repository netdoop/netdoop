import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { message } from 'antd';
import { FetchParams, getSearchItemsString, getOrderByString } from './common';

export const fetchDeviceEvents = async (params: FetchParams): Promise<API.listDeviceEventsData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDeviceEvents.listDeviceEvents({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch device events');
    return { Data: [], Total: 0 };
  }
};

export const useDeviceEvents = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [events, setEvents] = useState<API.DeviceEvent[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchDeviceEvents(params)
      setTotal(result.Total || 0)
      setEvents(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch devices');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);
  return { events, total, loading, reload };
}

export const useDeviceEvent = (ts:number) => {
  const access = useAccess()

  const [event, setEvent] = useState<API.DeviceEvent | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchEvent = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const event = await services.omcDeviceEvents.getDeviceEvent({ ts: ts });
        setEvent(event);
      }
    } catch (error) {
      console.error('Failed to fetch event:', error);
      setEvent(undefined);
    } finally {
      setLoading(false);
    }
  };


  useEffect(() => {
    fetchEvent();
  }, [ts]);

  return { event, loading };
};

