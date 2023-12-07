
import { caculateInterval } from "./data";
import { SearchItem, getSearchItemsString } from "./common";
import { timestampToString } from "@/utils/format";
import services from '@/services/netdoop';
import dayjs from "dayjs";

export type DevicePmValuePoint = {
  time: number;
  value: number;
};

export const fetchPmValueData = async (
  dateRange:[dayjs.Dayjs, dayjs.Dayjs], 
  interval?: number,
  searchItems?: SearchItem[],
  groupBy?: string[],
  func: 'sum'|'increase'|'last' ='last',
  offset?:number,
  limit?:number,
) => {
  try {
    const from = dateRange[0].unix();
    const to = dateRange[1].unix();
    const _interval = interval || caculateInterval(dateRange)
    let selects: API.TSQuerySelect[] = [
      { Func: 'text', Source: 'name', Name: 'name' },
      { Func: func, Source: 'value', Name: 'value' },
    ]
    const query: API.TSQueryCommand = {
      From: from,
      To: to,
      TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
      Query: {
        Id: "device_performance_values",
        Source: 'device_performance_values',
        Interval: _interval,
        Search: getSearchItemsString(searchItems ||[]),
        GroupBy: groupBy||[],
        Select: selects,
        Offset: offset,
        Limit: limit,
      }
    };
    const result = await services.omcData.queryData(query);
    const series = result.Data as API.TSSeries[] || [];
    
    const num = Math.floor((to - from) / _interval) + 1;
    let items: Record<string, any>[] = []
    for (let i = 0; i < num; i++) {
      const currentTime = from + i * _interval;
      series.forEach((serie) => {
        items.push({ name: serie.GroupBy?.name, time: currentTime, timeText: timestampToString(currentTime, _interval), value: undefined })
      });
    };

    series.forEach((serie) => {
      const points = serie.Points as API.TSPoint[] || [];
      points.forEach((point) => {
        const v = point as unknown as DevicePmValuePoint
        const item = items.find((item) => item.time === v.time && item.name === serie.GroupBy?.name);
        if (item) {
          item.value = v.value
        }
      })
    });

    return items
  } catch (error) {
    console.log(error)
  }
  return []
};

// export const fetchPmValueIncrData = async (
//   searchItems: SearchItem[],
//   groupBy: string[],
//   from: number,
//   to: number,
//   interval?: number,
// ) => {
//   try {
//     const _interval = interval || caculateInterval(from, to)
//     let selects: API.TSQuerySelect[] = [
//       { Func: 'text', Source: 'name', Name: 'name' },
//       { Func: 'increase', Source: 'value', Name: 'value' },
//     ]
//     const query: API.TSQueryCommand = {
//       From: from,
//       To: to,
//       TimeZone: Intl.DateTimeFormat().resolvedOptions().timeZone,
//       Query: {
//         Id: "device_performance_values",
//         Source: 'device_performance_values',
//         Interval: _interval,
//         Search: getSearchItemsString(searchItems),
//         GroupBy: groupBy,
//         Select: selects,
//       }
//     };
//     const result = await services.omcData.queryData(query);
//     const series = result.Data as API.TSSeries[] || [];

//     const num = Math.floor((to - from) / _interval) + 1;
//     let items: Record<string, any>[] = []
//     for (let i = 0; i < num; i++) {
//       const currentTime = from + i * _interval;
//       series.forEach((serie) => {
//         items.push({ name: serie.GroupBy?.name, time: currentTime, timeText: timestampToString(currentTime, _interval), value: undefined })
//       });
//     };

//     series.forEach((serie) => {
//       const points = serie.Points as API.TSPoint[] || [];
//       points.forEach((point) => {
//         const v = point as unknown as DevicePmValuePoint
//         const item = items.find((item) => item.time === v.time && item.name === serie.GroupBy?.name);
//         if (item) {
//           item.value = v.value
//         }
//       })
//     });

//     return items
//   } catch (error) {
//     console.log(error)
//   }
//   return []
// };
