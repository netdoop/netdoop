// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List routes GET /routes */
export async function listRoutes(options?: { [key: string]: any }) {
  return request<API.listRoutesBody>('/routes', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Get information of system GET /system/info */
export async function getSystemInfo(options?: { [key: string]: any }) {
  return request<API.systemInfoData>('/system/info', {
    method: 'GET',
    ...(options || {}),
  });
}

/** Get current time of system GET /system/time */
export async function getSystemTime(options?: { [key: string]: any }) {
  return request<API.systemTimeData>('/system/time', {
    method: 'GET',
    ...(options || {}),
  });
}
