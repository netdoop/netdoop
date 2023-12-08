import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { message } from 'antd';
import { FetchParams, getSearchItemsString, getOrderByString } from './common';

export const fetchDeviceMethodCalls = async (params: FetchParams): Promise<API.listDeviceMethodCallsData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDeviceMethodCalls.listDeviceMethodCalls({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch device methodCalls');
    return { Data: [], Total: 0 };
  }
};

export const useDeviceMethodCalls = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [methodCalls, setMethodCalls] = useState<API.DeviceMethodCall[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchDeviceMethodCalls(params)
      setTotal(result.Total || 0)
      setMethodCalls(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch device method calls');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);
  return { methodCalls, total, loading, reload };
}

export const useDeviceMethodCall = (ts:number) => {
  const access = useAccess()

  const [methodCall, setMethodCall] = useState<API.DeviceMethodCall | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchMethodCall = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const methodCall = await services.omcDeviceMethodCalls.getDeviceMethodCall({ ts: ts });
        setMethodCall(methodCall);
      }
    } catch (error) {
      console.error('Failed to fetch method call:', error);
      setMethodCall(undefined);
    } finally {
      setLoading(false);
    }
  };


  useEffect(() => {
    fetchMethodCall();
  }, [ts]);

  return { methodCall, loading };
};


