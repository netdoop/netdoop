// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List device events GET /omc/device-events */
export async function listDeviceEvents(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceEventsParams,
  options?: { [key: string]: any },
) {
  return request<API.listDeviceEventsData>('/omc/device-events', {
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

/** Get device event GET /omc/device-events/${param0} */
export async function getDeviceEvent(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceEventParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<API.DeviceEvent>(`/omc/device-events/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
