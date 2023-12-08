import { useEffect, useState } from 'react';
// import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { FetchParams, getSearchItemsString, getOrderByString, updateSearchItemWithValue, SearchItem } from './common';
import { message } from 'antd';

export const fetchKPIMeasures = async (params: FetchParams): Promise<API.listKPIMeasuresData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omckpi.listKpiMeasures({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch kpi measures');
    return { Data: [], Total: 0 };
  }
};

export const useKPIMeasures = (productType?: string) => {
  const [total, setTotal] = useState<number>(0);
  const [measures, setMeasures] = useState<API.KPIMeas[]>([]);
  const [measuresSets, setMeasuresSets] = useState<Record<string, API.KPIMeas[]>>({});

  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      let searchItems: SearchItem[] = []
      searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
      const result = await fetchKPIMeasures({ current: 0, pageSize: -1, searchItems: searchItems })
      const updateMeasures = result.Data || [];
      setTotal(result.Total || 0)
      setMeasures(updateMeasures);

      let sets:Record<string, API.KPIMeas[]> = {};
      updateMeasures.forEach(item => {
        if (item.MeasTypeSet) {
          let set: API.KPIMeas[] = sets[item.MeasTypeSet] || [];
          set.push(item)
          sets[item.MeasTypeSet] = set
        }
      })
      setMeasuresSets(sets)
    } catch (error) {
      message.error('Failed to fetch kpi measures');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, []);

  const measureByMeasTypeId = (measTypeId: string | undefined): API.KPIMeas | undefined => {
    let temp: API.KPIMeas | undefined = undefined
    if (measTypeId) {
      measures?.forEach(item => {
        if (item.MeasTypeID === measTypeId) {
          temp = item
          return
        }
      });
    }
    return temp
  }
  return {
    loading,
    total,
    measures,
    measuresSets,
    reload,
    measureByMeasTypeId,
  };
};

export const setKPIMeasureEnable = async (record: API.KPIMeas) => {
  try {
    if (record.Id) {
      await services.omckpi.setKpiMeasureEnable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to enable product');
  }
};

export const setKPIMeasureDisable = async (record: API.KPIMeas) => {
  try {
    if (record.Id) {
      await services.omckpi.setKpiMeasureDisable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable product');
  }
};