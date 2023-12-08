// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List devices GET /omc/devices */
export async function listDevices(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDevicesParams,
  options?: { [key: string]: any },
) {
  return request<API.listDevicesData>('/omc/devices', {
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

/** Create a device POST /omc/devices */
export async function createDevice(body: API.createDeviceBody, options?: { [key: string]: any }) {
  return request<API.Device>('/omc/devices', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get a device GET /omc/devices/${param0} */
export async function getDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Device>(`/omc/devices/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update device info PUT /omc/devices/${param0} */
export async function updateDeviceInfo(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateDeviceInfoParams,
  body: API.updateDeviceInfoBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Device>(`/omc/devices/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete a device DELETE /omc/devices/${param0} */
export async function deleteDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDeviceParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/devices/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Add object of a device POST /omc/devices/${param0}/add-device-object */
export async function addDeviceObject(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.addDeviceObjectParams,
  body: API.addDeviceObejectBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/add-device-object`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete object of a device POST /omc/devices/${param0}/delete-device-object */
export async function deleteDeviceObject(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDeviceObjectParams,
  body: API.deleteDeviceObejectBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/delete-device-object`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Set a device disable PUT /omc/devices/${param0}/disable */
export async function setDeviceDisable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setDeviceDisableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Map>(`/omc/devices/${param0}/disable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set a device enable PUT /omc/devices/${param0}/enable */
export async function setDeviceEnable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setDeviceEnableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Map>(`/omc/devices/${param0}/enable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Get parameter names of a device POST /omc/devices/${param0}/get-parameter-names */
export async function getDeviceParameterNames(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceParameterNamesParams,
  body: API.getDeviceParameterNamesBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/get-parameter-names`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Get parameter values of a device POST /omc/devices/${param0}/get-parameter-values */
export async function getDeviceParameterValues(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDeviceParameterValuesParams,
  body: API.getDeviceParameterValuesBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/get-parameter-values`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Set group for device PUT /omc/devices/${param0}/group */
export async function setGroupForDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setGroupForDeviceParams,
  body: API.setGroupForDeviceBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Device>(`/omc/devices/${param0}/group`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Get the methods of a device GET /omc/devices/${param0}/methods */
export async function listDeviceMethods(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceMethodsParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<string[]>(`/omc/devices/${param0}/methods`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Get the parameters of a device GET /omc/devices/${param0}/parameters */
export async function listDeviceParameters(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeviceParametersParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.ParameterValues>(`/omc/devices/${param0}/parameters`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set a device perf disable POST /omc/devices/${param0}/perf-disable */
export async function setDevicePerfDisable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setDevicePerfDisableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Map>(`/omc/devices/${param0}/perf-disable`, {
    method: 'POST',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set a device perf enable POST /omc/devices/${param0}/perf-enable */
export async function setDevicePerfEnable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setDevicePerfEnableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Map>(`/omc/devices/${param0}/perf-enable`, {
    method: 'POST',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Reboot a device POST /omc/devices/${param0}/reboot */
export async function rebootDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rebootDeviceParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/reboot`, {
    method: 'POST',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set parameter values of a device POST /omc/devices/${param0}/set-parameter-values */
export async function setDeviceParameterValues(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setDeviceParameterValuesParams,
  body: API.setDeviceParameterValuesBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/set-parameter-values`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Upgrade device POST /omc/devices/${param0}/upgrade */
export async function ugradeDevice(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.ugradeDeviceParams,
  body: API.upgradeDeviceBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/upgrade`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Upload file of device POST /omc/devices/${param0}/upload-file */
export async function uploadDeviceFile(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.uploadDeviceFileParams,
  body: API.uploadDeviceFileBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DeviceMethodCall>(`/omc/devices/${param0}/upload-file`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}
