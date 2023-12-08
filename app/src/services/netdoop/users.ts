// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List users GET /iam/users */
export async function listUsers(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listUsersParams,
  options?: { [key: string]: any },
) {
  return request<API.listUsersData>('/iam/users', {
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

/** Create user POST /iam/users */
export async function createUser(body: API.createUserBody, options?: { [key: string]: any }) {
  return request<API.User>('/iam/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get user GET /iam/users/${param0} */
export async function getUser(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getUserParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update user PUT /iam/users/${param0} */
export async function updateUser(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateUserParams,
  body: API.updateUserBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete user DELETE /iam/users/${param0} */
export async function deleteUser(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteUserParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/iam/users/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Change user password PUT /iam/users/${param0}/change-password */
export async function changeUserPassword(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.changeUserPasswordParams,
  body: API.changePasswordBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}/change-password`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Set user disable PUT /iam/users/${param0}/disable */
export async function setUserDisable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setUserDisableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}/disable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set user enable PUT /iam/users/${param0}/enable */
export async function setUserEnable(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setUserEnableParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}/enable`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Reset user password PUT /iam/users/${param0}/reset-password */
export async function resetUserPassword(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.resetUserPasswordParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Map>(`/iam/users/${param0}/reset-password`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Get roles for user GET /iam/users/${param0}/roles */
export async function getUserRoles(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getUserRolesParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<string[]>(`/iam/users/${param0}/roles`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set roles for user PUT /iam/users/${param0}/roles */
export async function setUserRoles(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setUserRolesParams,
  body: API.setUserRoles,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}/roles`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Add roles for user POST /iam/users/${param0}/roles */
export async function addUserRoles(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.addUserRolesParams,
  body: API.addUserRolesBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.User>(`/iam/users/${param0}/roles`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}
