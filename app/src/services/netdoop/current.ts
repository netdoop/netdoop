// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** Get current user information GET /iam/current */
export async function getCurrent(options?: { [key: string]: any }) {
  return request<API.User>('/iam/current', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Change current user password PUT /iam/current/change-password */
export async function changeCurrentPassword(
  body: API.changePasswordBody,
  options?: { [key: string]: any },
) {
  return request<any>('/iam/current/change-password', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
