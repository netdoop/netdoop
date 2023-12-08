// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List deleted devices GET /omc/deleted-devices */
export async function listDeletedDevices(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeletedDevicesParams,
  options?: { [key: string]: any },
) {
  return request<API.listDeletedDevicesData>('/omc/deleted-devices', {
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

/** Delete a device permanently DELETE /omc/deleted-devices/${param0} */
export async function deleteDeletedDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDeletedDeviceParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/deleted-devices/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Recover a deleted device POST /omc/deleted-devices/${param0}/recover */
export async function recoverDeletedDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.recoverDeletedDeviceParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/deleted-devices/${param0}/recover`, {
    method: 'POST',
    params: { ...queryParams },
    ...(options || {}),
  });
}
