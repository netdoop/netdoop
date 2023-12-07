import dayjs from "dayjs";

export type DataItem = {
  name: string,
  time: number,
  timeText: string,
  value: number | undefined,
}

export function caculateInterval(dateRange:[dayjs.Dayjs, dayjs.Dayjs]) {
  const from = dateRange[0].unix();
  const to = dateRange[1].unix();
  const diff = to - from
  if (diff <= 3600) {
    return 300
  } else if (diff <= 3600 * 6) {
    return 300
  } else if (diff <= 3600 * 12) {
    return 900
  } else if (diff <= 3600 * 24) {
    return 1800
  } else if (diff <= 86400 * 6) {
    return 3600
  }
  return 86400
}



