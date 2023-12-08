// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List transfer logs GET /omc/transfer-logs */
export async function listDeviceTransferLogs(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceTransferLogsParams,
  options?: { [key: string]: any },
) {
  return request<API.listDeviceTransferLogsData>('/omc/transfer-logs', {
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

/** Get device transfer log GET /omc/transfer-logs/${param0} */
export async function getDeviceTransferLog(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceTransferLogParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<API.DeviceTransferLog>(`/omc/transfer-logs/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Delete device transfer log DELETE /omc/transfer-logs/${param0} */
export async function deleteDeviceTransferLog(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDeviceTransferLogParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<any>(`/omc/transfer-logs/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}
