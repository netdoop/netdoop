// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List products GET /omc/products */
export async function listProducts(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listProductsParams,
  options?: { [key: string]: any },
) {
  return request<API.listProductsData>('/omc/products', {
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

/** Create product POST /omc/products */
export async function createProduct(body: API.createProductBody, options?: { [key: string]: any }) {
  return request<API.Product>('/omc/products', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get product GET /omc/products/${param0} */
export async function getProduct(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getProductParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Product>(`/omc/products/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update product info PUT /omc/products/${param0} */
export async function updateProductInfo(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateProductInfoParams,
  body: API.updateProductInfoBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Product>(`/omc/products/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete product DELETE /omc/products/${param0} */
export async function deleteProduct(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteProductParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/products/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set product disable PUT /omc/products/${param0}/disable */
export async function setProductDisable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setProductDisableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Product>(`/omc/products/${param0}/disable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set product enable PUT /omc/products/${param0}/enable */
export async function setProductEnable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setProductEnableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Product>(`/omc/products/${param0}/enable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** List product firmwares GET /omc/products/${param0}/firmwares */
export async function listProductFirmwares(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listProductFirmwaresParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Firmware[]>(`/omc/products/${param0}/firmwares`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
