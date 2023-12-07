// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List firmwares GET /omc/firmwares */
export async function listFirmwares(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listFirmwaresParams,
  options?: { [key: string]: any },
) {
  return request<API.listFirmwaresData>('/omc/firmwares', {
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

/** Create firmware POST /omc/firmwares */
export async function createFirmware(
  body: {
    /** Version */
    Version: string;
    /** ProductType */
    ProductType: string;
    /** Product list */
    Products: string;
  },
  File?: File,
  options?: { [key: string]: any },
) {
  const formData = new FormData();

  if (File) {
    formData.append('File', File);
  }

  Object.keys(body).forEach((ele) => {
    const item = (body as any)[ele];

    if (item !== undefined && item !== null) {
      if (typeof item === 'object' && !(item instanceof File)) {
        if (item instanceof Array) {
          item.forEach((f) => formData.append(ele, f || ''));
        } else {
          formData.append(ele, JSON.stringify(item));
        }
      } else {
        formData.append(ele, item);
      }
    }
  });

  return request<API.Firmware>('/omc/firmwares', {
    method: 'POST',
    data: formData,
    requestType: 'form',
    ...(options || {}),
  });
}

/** Get firmware GET /omc/firmwares/${param0} */
export async function getFirmware(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getFirmwareParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Firmware>(`/omc/firmwares/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Delete firmware DELETE /omc/firmwares/${param0} */
export async function deleteFirmware(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteFirmwareParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/firmwares/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set firmware products PUT /omc/firmwares/${param0}/products */
export async function setFirmwareProducts(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setFirmwareProductsParams,
  body: API.setFirmwareProductsBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Device>(`/omc/firmwares/${param0}/products`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}
