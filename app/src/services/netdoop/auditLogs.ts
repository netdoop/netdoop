// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List audit logs GET /iam/audit-logs */
export async function listAuditLogs(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listAuditLogsParams,
  options?: { [key: string]: any },
) {
  return request<API.listAuditLogsData>('/iam/audit-logs', {
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

/** Get audit log GET /iam/audit-logs/${param0} */
export async function getAuditLog(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getAuditLogParams,
  options?: { [key: string]: any },
) {
  const { ts: param0, ...queryParams } = params;
  return request<API.AuditLog>(`/iam/audit-logs/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
