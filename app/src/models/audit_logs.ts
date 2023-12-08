import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { message } from 'antd';
import { FetchParams, getSearchItemsString, getOrderByString } from './common';

export const fetchAuditLogs = async (params: FetchParams): Promise<API.listAuditLogsData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.auditLogs.listAuditLogs({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch audit logs');
    return { Data: [], Total: 0 };
  }
};

export const useAuditLogs = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [events, setEvents] = useState<API.AuditLog[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchAuditLogs(params)
      setTotal(result.Total || 0)
      setEvents(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch audit logs');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);
  return { events, total, loading, reload };
}

export const useAuditLog = (ts:number) => {
  const access = useAccess()

  const [event, setEvent] = useState<API.AuditLog | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const fetchEvent = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const event = await services.auditLogs.getAuditLog({ ts: ts });
        setEvent(event);
      }
    } catch (error) {
      console.error('Failed to fetch audit log:', error);
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

