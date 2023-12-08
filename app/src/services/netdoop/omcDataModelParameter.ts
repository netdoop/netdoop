// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** Delete data model parameter DELETE /omc/datamodels/${param0}/parameter/${param1} */
export async function deleteDatamodelParameter(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteDatamodelParameterParams,
  options?: { [key: string]: any },
) {
  const { id: param0, parameter_id: param1, ...queryParams } = params;
  return request<any>(`/omc/datamodels/${param0}/parameter/${param1}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Create data model parameter POST /omc/datamodels/${param0}/parameters */
export async function createDatamodelParameter(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.createDatamodelParameterParams,
  body: API.createDataModelParameterBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.DataModelParameter>(`/omc/datamodels/${param0}/parameters`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Get data model parameter GET /omc/datamodels/${param0}/parameters/${param1} */
export async function getDatamodelParameter(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getDatamodelParameterParams,
  options?: { [key: string]: any },
) {
  const { id: param0, parameter_id: param1, ...queryParams } = params;
  return request<API.DataModelParameter>(`/omc/datamodels/${param0}/parameters/${param1}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
