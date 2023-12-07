// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** List objects GET /s3/objects */
export async function listObjects(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.listObjectsParams,
  options?: { [key: string]: any },
) {
  return request<API.listS3ObjectsData>('/s3/objects', {
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

/** Download an object from S3 GET /s3/objects/${param0}/${param1} */
export async function downloadObject(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.downloadObjectParams,
  options?: { [key: string]: any },
) {
  const { bucket: param0, key: param1, ...queryParams } = params;
  return request<any>(`/s3/objects/${param0}/${param1}`, {
    method: 'GET',
    params: {
      ...queryParams,
    },
    ...(options || {}),
  });
}

/** Put an object to S3 POST /s3/objects/${param0}/${param1} */
export async function putObject(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.putObjectParams,
  body: {},
  file?: File,
  options?: { [key: string]: any },
) {
  const { bucket: param0, key: param1, ...queryParams } = params;
  const formData = new FormData();

  if (file) {
    formData.append('file', file);
  }

  Object.keys(body).forEach((ele) => {
    const item = (body as any)[ele];

    if (item !== undefined && item !== null) {
      if (typeof item === 'object' && !(item instanceof File)) {
        if (item instanceof Array) {
          item.forEach((f) => formData.append(ele, f || ''));
        } else {
          formData.append(ele, JSON.stringify(item));
        }
      } else {
        formData.append(ele, item);
      }
    }
  });

  return request<API.S3Object>(`/s3/objects/${param0}/${param1}`, {
    method: 'POST',
    params: { ...queryParams },
    data: formData,
    requestType: 'form',
    ...(options || {}),
  });
}

/** Delete an object from S3 DELETE /s3/objects/${param0}/${param1} */
export async function deleteObject(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.deleteObjectParams,
  options?: { [key: string]: any },
) {
  const { bucket: param0, key: param1, ...queryParams } = params;
  return request<any>(`/s3/objects/${param0}/${param1}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {}),
  });
}

/** Get information for an object from S3 GET /s3/objects/${param0}/${param1}/info */
export async function getObjectInfo(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.getObjectInfoParams,
  options?: { [key: string]: any },
) {
  const { bucket: param0, key: param1, ...queryParams } = params;
  return request<API.S3Object>(`/s3/objects/${param0}/${param1}/info`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {}),
  });
}
