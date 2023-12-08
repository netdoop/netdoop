import { useEffect, useState } from 'react';
import services from '@/services/netdoop';
import message from 'antd/es/message';
import { FetchParams, getOrderByString, getSearchItemsString } from './common';

export const fetchDataModels = async (params: FetchParams): Promise<API.listDataModelData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDataModel.listDatamodels({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch data models');
    return { Data: [], Total: 0 };
  }
};

export const useDataModels = () => {
  const [total, setTotal] = useState<number>(0);
  const [dataModels, setDataModels] = useState<API.DataModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchDataModels({ current: 0, pageSize: -1 })
      setTotal(result.Total || 0)
      setDataModels(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch data models');
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    reload();
  }, []);

  const dataModelById = (id: number | undefined): API.DataModel | undefined => {
    let temp: API.DataModel | undefined = undefined
    if (id) {
      dataModels?.forEach(item => {
        if (item.Id === id) {
          temp = item
          return
        }
      });
    }
    return temp
  }
  return { dataModels, total, loading, reload, dataModelById };
}

export const useDataModel = (id: number | undefined) => {
  const [data, setData] = useState<API.DataModel | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      if (id) {
        const data = await services.omcDataModel.getDatamodel({ id });
        setData(data);
      }
    } catch (error) {
      console.error('Failed to fetch data model');
      setData(undefined);
    } finally {
      setLoading(false);
    }
  };
  useEffect(() => {
    reload();
  }, [id]);

  return { data, reload, loading };
};

export const createDataModel = async (params: API.createDataModelBody) => {
  try {
    await services.omcDataModel.createDatamodel(params);
  } catch (error) {
    message.error('Failed to create data model');
  }
};

export const deleteDataModel = async (record: API.DataModel) => {
  try {
    if (record.Id) {
      await services.omcDataModel.deleteDatamodel({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable data model');
  }
};