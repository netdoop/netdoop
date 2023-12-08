// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List tasks GET /omc/tasks */
export async function listTasks(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listTasksParams,
  options?: { [key: string]: any },
) {
  return request<API.listTasksData>('/omc/tasks', {
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

/** Create task POST /omc/tasks */
export async function createTask(body: API.createTaskBody, options?: { [key: string]: any }) {
  return request<API.Task>('/omc/tasks', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

/** Get task GET /omc/tasks/${param0} */
export async function getTask(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getTaskParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<API.Task>(`/omc/tasks/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Delete task DELETE /omc/tasks/${param0} */
export async function deleteTask(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteTaskParams,
  options?: { [key: string]: any },
) {
  const { id: param0, ...queryParams } = params;
  return request<any>(`/omc/tasks/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}
