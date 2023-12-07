// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List child groups GET /omc/groups */
export async function listGroups(options?: { [key: string]: any }) {
  return request<API.Group[]>('/omc/groups', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Create a group POST /omc/groups */
export async function createGroup(body: API.createGroupBody, options?: { [key: string]: any }) {
  return request<API.Group>('/omc/groups', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get group details GET /omc/groups/${param0} */
export async function getGroup(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getGroupParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Group>(`/omc/groups/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Update group information PUT /omc/groups/${param0} */
export async function updateGroupInfo(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.updateGroupInfoParams,
  body: API.updateGroupInfoBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Group>(`/omc/groups/${param0}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}

/** Delete a group DELETE /omc/groups/${param0} */
export async function deleteGroup(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteGroupParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/groups/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Get child groups GET /omc/groups/${param0}/children */
export async function getGroupChildren(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getGroupChildrenParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Group[]>(`/omc/groups/${param0}/children`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Set group parent PUT /omc/groups/${param0}/parent */
export async function setGroupParent(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.setGroupParentParams,
  body: API.setGroupParentBody,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Group>(`/omc/groups/${param0}/parent`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    params: { ...queryParams },
    data: body,
    ...(options || {}),
  });
}
