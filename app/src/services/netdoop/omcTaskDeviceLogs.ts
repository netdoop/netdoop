// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List log device logs GET /omc/device-logs */
export async function listDeviceLogs(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceLogsParams,
  options?: { [key: string]: any },
) {
  return request<API.listTaskDeviceLogsData>('/omc/device-logs', {
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

/** Get log device log GET /omc/device-logs/${param0} */
export async function getDeviceLog(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceLogParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<API.TaskDeviceLog>(`/omc/device-logs/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
