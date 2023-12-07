import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { timestampToString } from '@/utils/format';
const { queryData } = services.omcData;
import { DataItem, caculateInterval } from "./data";
import dayjs from 'dayjs';

export const SEVERITY_COLOR_MAP: { [key: string]: string } = {
  Critical: 'red',
  Major: 'orange',
  Minor: 'yellow',
  Warning: 'cyan',
  Indeterminate: 'purple',
  Cleared: 'green',
  Unknown: 'gray',
};

type AlarmCountPoint = {
  time: number;
  count: number;
};

export const useAlarmStats = (dateRange: [dayjs.Dayjs, dayjs.Dayjs]) => {
  const access = useAccess()
  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<Record<string, number>>({});
  const [loading, setLoading] = useState<boolean>(false);
  const reload = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const from = 0;
        const to = dateRange[1].unix();
        const query: API.TSQueryCommand = {
          From: from,
          To: to,
          TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
          Query: {
            Id: "device_alarm_status",
            Source: 'device_alarms',
            Interval: to - from,
            GroupBy: ['perceived_severity'],
            Search: 'alarm_cleared:false',
            Select: [
              {
                Func: 'text',
                Source: 'perceived_severity',
                Name: 'perceived_severity',
              },
              {
                Func: 'count',
                Source: 'time',
                Name: 'count',
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
          points.forEach((point) => {
            const v = point as unknown as AlarmCountPoint;
            const name = series.GroupBy?.perceived_severity || 'Unknow'
            const count = stats[name] || 0
            stats[name] = count + v.count
            total += count
          })
        })
        setData(stats)
        setTotal(total)
      }
    } catch (error) {
      console.error('Failed to fetch alarm status:', error);
      setData({});
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    reload();
  }, [dateRange]);

  return { total, data, loading, reload };
};

export const useAlarmEventTypeStats = (dateRange: [dayjs.Dayjs, dayjs.Dayjs]) => {
  const access = useAccess()
  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<Record<string, number>>({});
  const [loading, setLoading] = useState<boolean>(false);
  const fetchData = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const from = 0;
        const to = dateRange[1].unix();
        const query: API.TSQueryCommand = {
          From: from,
          To: to,
          TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
          Query: {
            Id: "device_alarm_status",
            Source: 'device_alarms',
            Interval: to - from,
            GroupBy: ['event_type'],
            Search: 'alarm_cleared:false',
            Select: [
              {
                Func: 'text',
                Source: 'event_type',
                Name: 'event_type',
              },
              {
                Func: 'count',
                Source: 'time',
                Name: 'count',
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
          points.forEach((point) => {
            const v = point as unknown as AlarmCountPoint;
            const name = series.GroupBy?.event_type || 'Unknow'
            const count = stats[name] || 0
            stats[name] = count + v.count
            total += count
          })
        })
        console.log({all: all, stats: stats})
        setData(stats)
        setTotal(total)
      }
    } catch (error) {
      console.error('Failed to fetch alarm status:', error);
      setData({});
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [dateRange]);

  return { total, data, loading };
};

export const useAlarmIdentifierStats = (dateRange: [dayjs.Dayjs, dayjs.Dayjs]) => {
  const access = useAccess()
  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<Record<string, number>>({});
  const [loading, setLoading] = useState<boolean>(false);
  const fetchData = async () => {
    setLoading(true);
    try {
      if (access?.isLogin) {
        const from = 0;
        const to = dateRange[1].unix();
        const query: API.TSQueryCommand = {
          From: from,
          To: to,
          TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
          Query: {
            Id: "device_alarm_status",
            Source: 'device_alarms',
            Interval: to - from,
            GroupBy: ['alarm_identifier'],
            Search: 'alarm_cleared:false',
            Select: [
              {
                Func: 'text',
                Source: 'alarm_identifier',
                Name: 'alarm_identifier',
              },
              {
                Func: 'count',
                Source: 'time',
                Name: 'count',
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
          points.forEach((point) => {
            const v = point as unknown as AlarmCountPoint;
            const name = series.GroupBy?.alarm_identifier || 'Unknow'
            const count = stats[name] || 0
            stats[name] = count + v.count
            total += count
          })
        })
        setData(stats)
        setTotal(total)
      }
    } catch (error) {
      console.error('Failed to fetch alarm status:', error);
      setData({});
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [dateRange]);

  return { total, data, loading };
};

export const useAlarmSeverityStatus = (dateRange: [dayjs.Dayjs, dayjs.Dayjs], interval?: number) => {
  const access = useAccess()
  const [max, setMax] = useState<number>(-1);
  const [data, setData] = useState<DataItem[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
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
            Id: "device_alarm_status",
            Source: 'device_alarms',
            Interval: _interval,
            GroupBy: ['perceived_severity'],
            Select: [
              {
                Func: 'text',
                Source: 'perceived_severity',
                Name: 'perceived_severity',
              },
              {
                Func: 'count',
                Source: 'time',
                Name: 'count',
              },
            ],
          }
        };
        const result = await queryData(query);
        const all = result.Data as API.TSSeries[] || [];
        let itemsMap: Record<string, DataItem[]> = {
          Critical: [],
          Major: [],
          Minor: [],
          Warning: [],
          Cleared: [],
          Unknow: [],
        };
        const num = Math.floor((to - from) / _interval) + 1;
        for (let i = 0; i < num; i++) {
          const currentTime = from + i * _interval;
          itemsMap.Critical.push({ name: "Critical", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          itemsMap.Major.push({ name: "Major", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          itemsMap.Minor.push({ name: "Minor", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          itemsMap.Warning.push({ name: "Warning", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          // itemsMap.Cleared.push({ name:"Cleared" ,time: currentTime, timeText:timestampToString(currentTime, _interval), value: 0})
          // itemsMap.Unknow.push({ name:"Unknow" ,time: currentTime, timeText:timestampToString(currentTime, _interval), value: 0})
        }

        let counts: Record<string, number> = {};
        all.forEach((series) => {
          const points = series.Points as API.TSPoint[] || []
          points.forEach((point) => {
            const v = point as unknown as AlarmCountPoint;
            const count = counts.timeText || 0;
            counts.timeText = count + v.count;

            let items = itemsMap[series.GroupBy?.perceived_severity || 'Unknow'];
            const current = items.find((item) => item.time === v.time);
            if (current) {
              current.value = v.count + (current.value || 0)
            }
          })
        })

        let items = [
          ...itemsMap.Critical,
          ...itemsMap.Major,
          ...itemsMap.Minor,
          ...itemsMap.Warning,
          ...itemsMap.Cleared,
          ...itemsMap.Unknow,
        ]

        let max = 0;
        for (const value of Object.values(counts)) {
          if (value > max) {
            max = value;
          }
        }

        setData(items)
        setMax(max)
      }
    } catch (error) {
      console.error('Failed to fetch alarm status:', error);
      setData([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [dateRange, interval]);

  return { max, data, loading };
};

export type DeviceAlarmStatusPoint = {
  time: number;
  total: number;
  critical: number;
  major: number;
  minor: number;
  warning: number;

};


export const useDeviceAlarmStatusValues = (productType: string, dateRange: [dayjs.Dayjs, dayjs.Dayjs], interval?: number) => {
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
  const search = productType !== "" ? "product_type:" + productType : "";
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
            Id: "device_alarm_status",
            Source: 'device_alarm_statuses',
            Interval: _interval,
            Search: search,
            GroupBy: ['product_type'],
            Select: [
              {
                Func: 'last',
                Source: 'critical',
                Name: 'critical',
              },
              {
                Func: 'last',
                Source: 'major',
                Name: 'major',
              },
              {
                Func: 'last',
                Source: 'minor',
                Name: 'minor',
              },
              {
                Func: 'last',
                Source: 'warning',
                Name: 'warning',
              },
            ],
          }
        };
        const num = Math.floor((to - from) / _interval) + 1;
        let items: DataItem[] = []
        for (let i = 0; i < num; i++) {
          const currentTime = from + i * _interval;
          items.push({ name: "Critical", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          items.push({ name: "Major", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          items.push({ name: "Minor", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
          items.push({ name: "Warning", time: currentTime, timeText: timestampToString(currentTime, _interval), value: 0 })
        }

        const result = await queryData(query);
        const series = result.Data as API.TSSeries[] || [];
        const points = series[0]?.Points as API.TSPoint[] || [];

        let max = -1
        points.forEach((point) => {
          const v = point as unknown as DeviceAlarmStatusPoint
          const total = v.critical + v.major + v.minor + v.warning;
          if (max < total) {
            max = total
          }
          increseValue(items, v.time, "Critical", v.critical)
          increseValue(items, v.time, "Major", v.major)
          increseValue(items, v.time, "Minor", v.minor)
          increseValue(items, v.time, "Warning", v.warning)
        })

        setData(items)
        setMax(max)
      }
    } catch (error) {
      console.error('Failed to fetch device alarm status:', error);
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