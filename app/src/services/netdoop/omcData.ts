// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** Query data series POST /omc/data */
export async function queryData(body: API.TSQueryCommand, options?: { [key: string]: any }) {
  return request<API.TSResult>('/omc/data', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
