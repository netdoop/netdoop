import { useEffect, useState } from 'react';
import services from '@/services/netdoop';
import message from 'antd/es/message';
import { FetchParams, SearchItem, getOrderByString, getSearchItemsString, updateSearchItemWithRangeValue, updateSearchItemWithValue } from './common';
import dayjs from 'dayjs';

export const fetchKPITemplates = async (params: FetchParams): Promise<API.listKPITemplatesData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omckpi.listKpiTemplates({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch kpi templates');
    return { Data: [], Total: 0 };
  }
};

export const useKPITemplates = () => {
  const [total, setTotal] = useState<number>(0);
  const [kpiTemplates, setKPITemplates] = useState<API.KPITemplate[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchKPITemplates({ current: 0, pageSize: -1 })
      setTotal(result.Total || 0)
      setKPITemplates(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch kpi templates');
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    reload();
  }, []);

  const kpiTemplateById = (id: number | undefined): API.KPITemplate | undefined => {
    let temp: API.KPITemplate | undefined = undefined
    if (id) {
      kpiTemplates?.forEach(item => {
        if (item.Id === id) {
          temp = item
          return
        }
      });
    }
    return temp
  }
  return { kpiTemplates, total, loading, reload, kpiTemplateById };
}

export const useKPITemplate = (id: number | undefined) => {
  const [data, setData] = useState<API.KPITemplate | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      if (id) {
        const data = await services.omckpi.getKpiTemplate({ id });
        setData(data);
      }
    } catch (error) {
      console.error('Failed to fetch KPI template');
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

export const createKPITemplate = async (params: API.createKPITemplateBody) => {
  try {
    await services.omckpi.createKpiTemplate(params);
  } catch (error) {
    message.error('Failed to create KPI template');
  }
};

export const deleteKPITemplate = async (record: API.KPITemplate) => {
  try {
    if (record.Id) {
      await services.omckpi.deleteKpiTemplate({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable KPI template');
  }
};

export const updateKPITemplate = async (record: API.KPITemplate, params: API.updateKPITemplateBody) => {
  try {
    if (record.Id) {
      await services.omckpi.updateKpiTemplate({ id: record.Id }, params);
    }
  } catch (error) {
    message.error('Failed to update KPI template');
  }
};

export const fetchKpiValuesByTemp = async (
  record: API.KPITemplate,
  params: FetchParams,
): Promise<API.listKPITemplateRecordsData> => {
  try {
    if (record.Id) {
      const q = getSearchItemsString(params.searchItems || [])
      const order = getOrderByString(params.sort || {})
      const result = await services.omckpi.listKpiTemplateRecords({
        id: record.Id,
        page: params.current,
        page_size: params.pageSize,
        q: q,
        order_by: order,
      });
      const data = result.Data as Record<string, any>[] || [];
      let updateData:Record<string, any>[] = [];
      data.forEach((item) => {
        let tmp:Record<string, any>=item;
        record.MeasTypeIds?.forEach((v) => {
          const name = v.toLowerCase();
          tmp[name] = item[name] !== undefined && item[name] !== null? parseFloat(item[name]): undefined;
        });
        updateData.push(tmp)
      })
      return {Data:updateData, Total: result.Total};

    } else {
      return { Data: [], Total: 0 };
    }
  } catch (error) {
    message.error('Failed to fetch kpi template records');
    return { Data: [], Total: 0 };
  }
};

export const useAllKpiValuesByTemp = (
  dateRange: [dayjs.Dayjs, dayjs.Dayjs],
  record?: API.KPITemplate,
  searchItems?: SearchItem[],
) => {
  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<Record<string, Record<string, any>[]>>({});
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      if (record) {
        const from = dateRange[0].format('YYYY-MM-DDTHH:mm:ssZ');
        const to = dateRange[1].format('YYYY-MM-DDTHH:mm:ssZ');

        let _searchItems = searchItems || [];
        _searchItems = updateSearchItemWithRangeValue(_searchItems, "time", from, to);
        _searchItems = updateSearchItemWithValue(_searchItems, "product_type", record.ProductType)
        const result = await fetchKpiValuesByTemp(record, { current: 0, pageSize: -1, searchItems: _searchItems, })
        setTotal(result.Total || 0)

        let update: Record<string, Record<string, any>[]> = {};
        record.MeasTypeIds?.forEach((v) => {
          update[v] = [];
        });
        const items = result.Data as Record<string, any>[] || [];
        items.forEach((item) => {
          record.MeasTypeIds?.forEach((v) => {
            const name = v.toLowerCase();
            // const value = parseFloat(item[name]);
            const tmp = {
              device_id: item.device_id,
              product_type: item.product_type,
              oui: item.oui,
              product_class: item.product_class,
              serial_number: item.serial_number,
              time: item.time,
              value: item[name],
            };
            update[v].push(tmp);
          });
        });
        setData(update);
      }
    } catch (error) {
      console.log(error)
      message.error('Failed to fetch kpi template records');
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    reload();
  }, []);

  return { total, loading, reload, data };
};
