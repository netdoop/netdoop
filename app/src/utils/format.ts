import moment from 'moment'
import dayjs from 'dayjs';

export function trim(str: string) {
  return str.trim();
}

export function timeSerial(ts: any) {
  if (ts) {
    let v = ts;
    if (typeof ts === 'string') {
      v = parseInt(ts, 10);
    }
    return Math.round(v / 1000)
  }
  return 0
}

export function formatTimeString(v: string) {
  if (v) {
    const t = moment(v)
    return t.format('YYYY-MM-DD HH:mm:ss');
  }
  return ""
}

export function formatTimestampWithTZ(ts: any) {
  const timeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;

  if (ts) {
    let v = ts;
    if (typeof ts === 'string') {
      v = parseInt(ts, 10);
    }
    const t = moment.unix(v)
    return timeZone + ": " + t.format('YYYY-MM-DD HH:mm:ss');
  }
  return timeZone + ": "
}

export function formatTimestamp(ts: any) {
  if (ts) {
    let v = ts;
    if (typeof ts === 'string') {
      v = parseInt(ts, 10);
    }
    if (v < 0) {
      return ""
    }
    const t = moment.unix(v)
    return t.format('YYYY-MM-DD HH:mm:ss');
  }
  return ""
}

export function formatTimestamp2(ts: any, k?: number) {
  if (ts) {
    let v = ts;
    if (typeof ts === 'string') {
      v = parseInt(ts, 10);
    }
    if (v < 0) {
      return ""
    }
    const t = moment.unix(v / (k || 1000000))
    return t.format('YYYY-MM-DD HH:mm:ss');
  }
  return ""
}

export function timestampToDayjs(ts: number) {
  return dayjs(new Date(ts * 1000))
}

export function timestampToString(ts: number, interval: number) {
  const t = moment.unix(ts)
  if (interval % 86400 === 0) {
    return t.format('MM-DD')
  } else if (interval % 3600 === 0) {
    return t.format('DD HH:mm')
  } else if (interval % 60 === 0) {
    return t.format('HH:mm')
  }
  return t.format('MM-DD HH:mm')
}

export function getDuratinon(secs: any) {
  if (secs){
    let v = secs;
    if (typeof secs === 'string') {
      v = parseInt(secs, 10);
    }
    if (v < 60) {
      return `${v} sec`;
    } else if (v < 3600) {
      return `${Math.floor(v / 60)} minutes`;
    } else if (v < 86400) {
      return `${Math.floor(v / 3600)} hours`;
    } else {
      return `${Math.floor(v / 86400)} days`;
    }
  }

}

export function getBytesUnit(max: number) {
  const units = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
  let unitIndex = 0;
  let v = max;
  let k = 1;
  while (v / 1024 >= 1024 && unitIndex < units.length - 1) {
    v = v / 1024;
    k = k * 1024
    unitIndex++;
  }
  return { k, unit: units[unitIndex] };
}

export function formatBytes(size: number, fixed?: number) {
  if (size < 0) {
    return "-"
  }
  const units = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
  let unitIndex = 0;
  let v = size;
  while (v >= 1024 && unitIndex < units.length - 1) {
    v = v / 1024;
    unitIndex++;
  }
  return `${v.toFixed(fixed || 1)} ${units[unitIndex]}`;

}