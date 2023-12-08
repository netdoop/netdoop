import { useEffect, useState } from 'react';
// import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import { FetchParams, getSearchItemsString, getOrderByString, updateSearchItemWithValue, SearchItem } from './common';
import { message } from 'antd';

export const fetchProducts = async (params: FetchParams): Promise<API.listProductsData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcProducts.listProducts({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch products');
    return { Data: [], Total: 0 };
  }
};

export const useProducts = (productType?:string) => {
  const [total, setTotal] = useState<number>(0);
  const [products, setProducts] = useState<API.Product[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      let searchItems: SearchItem[] = []
      searchItems = updateSearchItemWithValue(searchItems, "product_type", productType)
      const result = await fetchProducts({current:0, pageSize:-1, searchItems:searchItems})
      setTotal(result.Total || 0)
      setProducts(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch products');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, []);

  const getProduct = (oui: string | undefined, productClass: string | undefined): API.Product | undefined => {
    let product: API.Product | undefined = undefined
    if (oui && productClass) {
      products?.forEach(item => {
        if (item.Oui === oui && item.ProductClass === productClass) {
          product = item
          return
        }
      });
    }
    return product
  }

  return {
    loading,
    total,
    products,
    reload,
    getProduct,
  };
};

export const setProductEnable = async (record: API.Product) => {
  try {
    if (record.Id) {
      await services.omcProducts.setProductEnable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to enable product');
  }
};

export const setProductDisable = async (record: API.Product) => {
  try {
    if (record.Id) {
      await services.omcProducts.setProductDisable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable product');
  }
};


export const useProductFirmwares = (id?: number) => {
  const [firmwares, setFirmwares] = useState<API.Firmware[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      if (id !== undefined) {
       const result =  await services.omcProducts.listProductFirmwares({ id: id });
       setFirmwares(result)
      }
    } catch (error) {
      message.error('Failed to fetch products');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [id]);
  return {
    loading,
    firmwares,
    reload,
  };
};