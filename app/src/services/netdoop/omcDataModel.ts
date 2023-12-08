// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List data models GET /omc/datamodels */
export async function listDatamodels(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDatamodelsParams,
  options?: { [key: string]: any },
) {
  return request<API.listDataModelData>('/omc/datamodels', {
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

/** Create data model POST /omc/datamodels */
export async function createDatamodel(
  body: API.createDataModelBody,
  options?: { [key: string]: any },
) {
  return request<API.DataModel>('/omc/datamodels', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get data model GET /omc/datamodels/${param0} */
export async function getDatamodel(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDatamodelParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DataModel>(`/omc/datamodels/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Delete data model DELETE /omc/datamodels/${param0} */
export async function deleteDatamodel(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDatamodelParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/datamodels/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** List the parameters of a datamodel GET /omc/datamodels/${param0}/parameters */
export async function listDatamodelParameters(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDatamodelParametersParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.ParameterValues>(`/omc/datamodels/${param0}/parameters`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** List the templates of a datamodel GET /omc/datamodels/${param0}/templates */
export async function listDatamodelTemplates(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listDatamodelTemplatesParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.ParameterValues>(`/omc/datamodels/${param0}/templates`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
