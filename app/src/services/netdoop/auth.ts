// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** Authenticate a user POST /auth */
export async function postAuth(body: API.authBody, options?: { [key: string]: any }) {
  return request<API.authResponseBody>('/auth', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
