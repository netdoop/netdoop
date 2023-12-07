import { useEffect, useState } from 'react';
// import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { FetchParams, getSearchItemsString, getOrderByString } from './common';
import { message } from 'antd';

export const fetchFirmwares = async (params: FetchParams): Promise<API.listFirmwaresData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcFirmwares.listFirmwares({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch firmwares');
    return { Data: [], Total: 0 };
  }
};

export const useFirmwares = (params: FetchParams) => {
  // const access = useAccess()

  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<API.Firmware[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchFirmwares(params)
      setTotal(result.Total || 0)
      setData(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch firmwares');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);

  return {
    loading,
    total,
    data,
    reload,
  };
};

export const createFirmware = async (params: { Version: string; Products: Array<string>; ProductType: string}, file: File) => {
  try {
    await services.omcFirmwares.createFirmware({
      Version:params.Version,
      Products: params.Products.join(","),
      ProductType: params.ProductType,
    }, file);
  } catch (error) {
    message.error('Failed to create firmware');
  }
};


export const deleteFirmware = async (record: API.Firmware) => {
  try {
    if (record.Id) {
      await services.omcFirmwares.deleteFirmware({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to delete firmware');
  }
};