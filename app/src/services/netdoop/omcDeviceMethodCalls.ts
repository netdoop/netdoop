// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List device method calls GET /omc/device-method-calls */
export async function listDeviceMethodCalls(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceMethodCallsParams,
  options?: { [key: string]: any },
) {
  return request<API.listDeviceMethodCallsData>('/omc/device-method-calls', {
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

/** Get device method call GET /omc/device-method-calls/${param0} */
export async function getDeviceMethodCall(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceMethodCallParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/device-method-calls/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
