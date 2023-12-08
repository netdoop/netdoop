import services from '@/services/netdoop';
import { message } from 'antd';
import { FetchParams, getSearchItemsString, getOrderByString } from './common';
import { useAccess } from '@umijs/max';
import { useState, useEffect } from 'react';

export const fetchDeviceTransferLogs = async (params: FetchParams): Promise<API.listDeviceTransferLogsData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDeviceTransferLogs.listDeviceTransferLogs({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch device transfer logs');
    return { Data: [], Total: 0 };
  }
};

export const useDeviceTransferLogs = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [transferLogs, setTransferLogs] = useState<API.DeviceTransferLog[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchDeviceTransferLogs(params)
      setTotal(result.Total || 0)
      setTransferLogs(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch device transfer logs');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);
  return { transferLogs, total, loading, reload };
}

export const useDeviceTransferLog = (ts: number) => {
  const access = useAccess()

  const [transferLog, setTransferLog] = useState<API.DeviceTransferLog | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchTransferLog = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const transferLog = await services.omcDeviceTransferLogs.getDeviceTransferLog({ ts: ts });
        setTransferLog(transferLog);
      }
    } catch (error) {
      console.error('Failed to fetch device transfer log', error);
      setTransferLog(undefined);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTransferLog();
  }, [ts]);

  return { transferLog, loading };
};

export const deleteDeviceTransferLog = async (ts: number) => {
  try {
    await services.omcDeviceTransferLogs.deleteDeviceTransferLog({ ts: ts });
  } catch (error) {
    message.error('Failed to delete device transfer log');
  }
};