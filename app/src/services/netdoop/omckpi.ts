// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List kpi measures GET /omc/kpi/measures */
export async function listKpiMeasures(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listKpiMeasuresParams,
  options?: { [key: string]: any },
) {
  return request<API.listKPIMeasuresData>('/omc/kpi/measures', {
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

/** Create kpi measure POST /omc/kpi/measures */
export async function createKpiMeasure(
  body: API.createKPIMeasureBody,
  options?: { [key: string]: any },
) {
  return request<API.KPIMeas>('/omc/kpi/measures', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get kpi measure GET /omc/kpi/measures/${param0} */
export async function getKpiMeasure(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getKpiMeasureParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.KPIMeas>(`/omc/kpi/measures/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update kpi measure info PUT /omc/kpi/measures/${param0} */
export async function updateKpiMeasureInfo(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateKpiMeasureInfoParams,
  body: API.updateKPIMeasureInfoBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.KPIMeas>(`/omc/kpi/measures/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete kpi measure DELETE /omc/kpi/measures/${param0} */
export async function deleteKpiMeasure(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteKpiMeasureParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/kpi/measures/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set kpi measure disable PUT /omc/kpi/measures/${param0}/disable */
export async function setKpiMeasureDisable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setKpiMeasureDisableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.KPIMeas>(`/omc/kpi/measures/${param0}/disable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set kpi measure enable PUT /omc/kpi/measures/${param0}/enable */
export async function setKpiMeasureEnable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setKpiMeasureEnableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.KPIMeas>(`/omc/kpi/measures/${param0}/enable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** List kpi templates GET /omc/kpi/templates */
export async function listKpiTemplates(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listKpiTemplatesParams,
  options?: { [key: string]: any },
) {
  return request<API.listKPITemplatesData>('/omc/kpi/templates', {
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

/** Create kpi template POST /omc/kpi/templates */
export async function createKpiTemplate(
  body: API.createKPITemplateBody,
  options?: { [key: string]: any },
) {
  return request<API.KPITemplate>('/omc/kpi/templates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get kpi template GET /omc/kpi/templates/${param0} */
export async function getKpiTemplate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getKpiTemplateParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.KPITemplate>(`/omc/kpi/templates/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update kpi template PUT /omc/kpi/templates/${param0} */
export async function updateKpiTemplate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateKpiTemplateParams,
  body: API.updateKPITemplateBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.KPITemplate>(`/omc/kpi/templates/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete kpi template DELETE /omc/kpi/templates/${param0} */
export async function deleteKpiTemplate(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteKpiTemplateParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/kpi/templates/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** List kpi template records GET /omc/kpi/templates/${param0}/records */
export async function listKpiTemplateRecords(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listKpiTemplateRecordsParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.listKPITemplateRecordsData>(`/omc/kpi/templates/${param0}/records`, {
    method: 'GET',
    params: {
      // page has a default value: 1
      page: '1',
      // page_size has a default value: 20
      page_size: '20',

      ...queryParams,
    },
    ...(options || {}),
  });
}
