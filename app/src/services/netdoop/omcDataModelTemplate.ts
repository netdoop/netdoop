// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** Delete data model template DELETE /omc/datamodels/${param0}/template/${param1} */
export async function deleteDatamodelTemplate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDatamodelTemplateParams,
  options?: { [key: string]: any },
) {
  const { id: param0, template_id: param1, ...queryParams } = params;
  return request<any>(`/omc/datamodels/${param0}/template/${param1}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Create data model template POST /omc/datamodels/${param0}/templates */
export async function createDatamodelTemplate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.createDatamodelTemplateParams,
  body: API.createDataModelTemplateBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DataModelTemplate>(`/omc/datamodels/${param0}/templates`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Get data model template GET /omc/datamodels/${param0}/templates/${param1} */
export async function getDatamodelTemplate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDatamodelTemplateParams,
  options?: { [key: string]: any },
) {
  const { id: param0, template_id: param1, ...queryParams } = params;
  return request<API.DataModelTemplate>(`/omc/datamodels/${param0}/templates/${param1}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
