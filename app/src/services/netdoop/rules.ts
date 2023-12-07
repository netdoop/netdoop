// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List rules GET /iam/rules */
export async function listRules(options?: { [key: string]: any }) {
  return request<API.listRulesBody>('/iam/rules', {
    method: 'GET',
    ...(options || {}),
  });
}
