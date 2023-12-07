import { useEffect, useState } from 'react';
import services from '@/services/netdoop';
import message from 'antd/es/message';

export const fetchRules = async (): Promise<API.listRulesBody> => {
  try {
    const result = await services.iamRules.listRules();
    return result;
  } catch (error) {
    message.error('Failed to fetch roles');
    return { Data: [], Total: 0 };
  }
};

export const useRules = () => {
  const [total, setTotal] = useState<number>(0);
  const [rules, setRules] = useState<string[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchRules()
      setTotal(result.Total || 0)
      setRules(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch iam rules');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, []);
  return { rules, total, loading, reload };
}