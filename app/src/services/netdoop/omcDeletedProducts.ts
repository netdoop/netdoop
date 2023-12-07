// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List deleted products GET /omc/deleted-products */
export async function listDeletedProducts(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDeletedProductsParams,
  options?: { [key: string]: any },
) {
  return request<API.listDeletedProductsData>('/omc/deleted-products', {
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

/** Delete deleted product DELETE /omc/deleted-products/${param0} */
export async function deleteDeletedProduct(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDeletedProductParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/deleted-products/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}
