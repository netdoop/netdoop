// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List roles GET /iam/roles */
export async function listRoles(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listRolesParams,
  options?: { [key: string]: any },
) {
  return request<API.listRolesData>('/iam/roles', {
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

/** Create a role POST /iam/roles */
export async function createRole(body: API.createRoleBody, options?: { [key: string]: any }) {
  return request<API.Role>('/iam/roles', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get role GET /iam/roles/${param0} */
export async function getRole(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getRoleParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Role>(`/iam/roles/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update role PUT /iam/roles/${param0} */
export async function updateRole(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateRoleParams,
  body: API.updateRoleBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Role>(`/iam/roles/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete role DELETE /iam/roles/${param0} */
export async function deleteRole(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteRoleParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/iam/roles/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set role disable PUT /iam/roles/${param0}/disable */
export async function setRoleDisable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setRoleDisableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Role>(`/iam/roles/${param0}/disable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set role enable PUT /iam/roles/${param0}/enable */
export async function setRoleEnable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setRoleEnableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Role>(`/iam/roles/${param0}/enable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set rules for role PUT /iam/roles/${param0}/rules */
export async function setRoleRules(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setRoleRulesParams,
  body: API.setRoleApiRulesBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/roles/${param0}/rules`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}
