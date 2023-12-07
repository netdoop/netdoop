// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List device alarms GET /omc/device-alarms */
export async function listDeviceAlarms(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceAlarmsParams,
  options?: { [key: string]: any },
) {
  return request<API.listDeviceAlarmsData>('/omc/device-alarms', {
    method: 'GET',
    params: {
      // page has a default value: 1
      page: '1',
      // page_size has a default value: 20
      page_size: '20',

      ...params,
    },
    ...(options || {}),
  });
}

/** Get device alarm GET /omc/device-alarms/${param0} */
export async function getDeviceAlarm(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceAlarmParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<API.DeviceAlarm>(`/omc/device-alarms/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
