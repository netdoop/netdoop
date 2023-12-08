import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { timestampToString } from '@/utils/format';
import { DataItem, caculateInterval } from "./data";
import dayjs from 'dayjs';

const { queryData } = services.omcData;

export type DeviceStatusPoint = {
  time: number;
  total: number;
  active: number;
  online: number;
  offline: number;
};

export const useDeviceCurrentStatusValue = (productType: string, dateRange:[dayjs.Dayjs, dayjs.Dayjs]) => {
  const access = useAccess()
  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<Record<string, number>>({});
  const [loading, setLoading] = useState<boolean>(false);
  
  const fetchData = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const from = dateRange[0].unix();
        const to = dateRange[1].unix();
        const query: API.TSQueryCommand = {
          From: from,
          To: to,
          TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
          Query: {
            Id: "device_status",
            Source: 'device_statuses',
            Interval: to - from,
            Limit: 1,
            Search: "product_type:" + productType,
            Select: [
              {
                Func: 'last',
                Source: 'total',
                Name: 'total',
              },
              {
                Func: 'last',
                Source: 'active',
                Name: 'active',
              },
              {
                Func: 'last',
                Source: 'online',
                Name: 'online',
              },
              {
                Func: 'last',
                Source: 'offline',
                Name: 'offline',
              },
            ],
          }
        };
        const result = await queryData(query);
        const all = result.Data as API.TSSeries[] || [];
        let total = 0
        let stats: Record<string, number> = {}
        all.forEach((series) => {
          const points = series.Points as API.TSPoint[] || []
          const point = points[0] as unknown as DeviceStatusPoint;
          if (point) {
            total = point.total
            stats['active'] = point.active
            stats['offline'] = point.offline
            stats['online'] = point.online
          }
        })
        setData(stats)
        setTotal(total)
      }
    } catch (error) {
      console.error('Failed to fetch device status:', error);
      setData({});
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [productType, dateRange]);

  return { total, data, loading };
};

export const useDeviceStatusValues = (productType: string, dateRange:[dayjs.Dayjs, dayjs.Dayjs], interval?: number) => {
  const access = useAccess()
  const [max, setMax] = useState<number>(-1);
  const [data, setData] = useState<DataItem[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const increseValue = (items: DataItem[], time: number, name: string, value: number) => {
    const item = items.find((item) => item.time === time && item.name === name);
    if (item) {
      item.value = value + (item.value || 0)
    }
  }
  const fetchData = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const from = dateRange[0].unix();
        const to = dateRange[1].unix();
        const _interval = interval || caculateInterval(dateRange)

        const query: API.TSQueryCommand = {
          From: from,
          To: to,
          TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
          Query: {
            Id: "device_status",
            Source: 'device_statuses',
            Interval: _interval,
            Search: "product_type:" + productType,
            Select: [
              {
                Func: 'last',
                Source: 'total',
                Name: 'total',
              },
              {
                Func: 'last',
                Source: 'active',
                Name: 'active',
              },
              {
                Func: 'last',
                Source: 'online',
                Name: 'online',
              },
              {
                Func: 'last',
                Source: 'offline',
                Name: 'offline',
              },
            ],
          }
        };
        const num = Math.floor((to - from) / _interval) + 1;
        let items: DataItem[] = []
        for (let i = 0; i < num; i++) {
          const currentTime = from + i * _interval;
          items.push({ name: "active", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          items.push({ name: "online", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          items.push({ name: "offline", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          items.push({ name: "inactive", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
        }

        const result = await queryData(query);
        const series = result.Data as API.TSSeries[] || [];
        const points = series[0]?.Points as API.TSPoint[] || []

        let max = -1
        points.forEach((point) => {
          const v = point as unknown as DeviceStatusPoint
          if (max < v.total) {
            max = v.total
          }
          // increseValue(items, v.time, "inactive", v.total - v.active)
          increseValue(items, v.time, "active", v.active)
          increseValue(items, v.time, "online", v.online)
          increseValue(items, v.time, "offline", v.offline)
        })

        setData(items)
        setMax(max)
      }
    } catch (error) {
      console.error('Failed to fetch device status:', error);
      setData([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [productType, dateRange, interval]);

  return { max, data, loading };
};

