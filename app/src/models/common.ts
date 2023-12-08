import { SortOrder } from "antd/es/table/interface";
import { useEffect, useState } from "react";

export type SearchItemValue = {
  value?: any,
  rangeValue?: any[2],
  symbol?: string,
}
export type SearchItem = {
  name: string,
  values: SearchItemValue[],
};

export type FetchParams = {
  current?: number,
  pageSize?: number,
  sort?: Record<string, SortOrder>,
  keyword?: string,
  searchItems?: SearchItem[],
};

export const getOrderByString = (data: Record<string, SortOrder>) => {
  const orders: string[] = []
  Object.entries(data).forEach(([k, v]) => {
    const order = v === 'descend' ? k + '-' : k;
    orders.push(order);
  })
  if (orders.length === 0) {
    return undefined
  }
  return orders.join(',')
}

export const getSearchItemsString = (data: SearchItem[]) => {
  const q: string[] = []
  data.forEach((item) => {
    let values: string[] = [];
    item.values.forEach((v) => {
      if (v.value) {
        let tmp = "";
        if (v.symbol && v.symbol !== "=") {
          tmp += v.symbol;
        }
        if (typeof v.value === 'string') {
          tmp += "'" + v.value + "'"
        } else if (typeof v.value === 'number') {
          tmp += v.value
        } else {
          tmp += v.value
        }
        values.push(tmp)
      }else if (v.rangeValue) {
        let tmp = ""+v.rangeValue[0] +".."+v.rangeValue[1];
        values.push(tmp)
      }
    });
    q.push(item.name + ":" + values.join(','));
  })
  if (q.length === 0) {
    return undefined
  }
  return q.join(' ')
}

export const updateSearchItem = (data: SearchItem[], item: SearchItem): SearchItem[] => {
  let update: SearchItem[] = data
  const index = update.findIndex((v) => v.name === item.name)
  if (index < 0) {
    update.push(item)
  } else {
    update[index] = item
  }
  return update
}

export const removeSearchItem = (data: SearchItem[], name: string): SearchItem[] => {
  let update = data.filter((item) => item.name !== name);
  return update
}

export const updateSearchItemWithValue = (data: SearchItem[], name: string, value?: any, symbol?: string): SearchItem[] => {
  if (value && value !== '') {
    return updateSearchItem(data, { name: name, values: [{ value: value, symbol: symbol }] })
  } else {
    return removeSearchItem(data, name)
  }
}

export const updateSearchItemWithValues = (data: SearchItem[], name: string, values?: any[], symbol?: string): SearchItem[] => {
  if (values && values.length > 0) {
    const item: SearchItem = { name: name, values: [] }
    values.forEach((v) => {
      item.values.push({ value: v, symbol: symbol })
    })
    return updateSearchItem(data, item)
  } else {
    return removeSearchItem(data, name)
  }
}

export const updateSearchItemWithRangeValue = (data: SearchItem[], name: string, from?: any, to?: any, symbol?: string): SearchItem[] => {
  if (from && to) {
    return updateSearchItem(data, { name: name, values: [{ rangeValue: [from, to], symbol: symbol }] })
  } else {
    return removeSearchItem(data, name)
  }
}


export const useSearch = (value: string) => {
  const [items, setItems] = useState<SearchItem[]>([]);
  const [remain, setRemain] = useState<string>('');

  const clearItem = (key: number) => {
    const updatedSearchItems = items.filter((item, index) => index !== key);
    setItems(updatedSearchItems);
  }
  const clearItems = () => {
    setItems([]);
  }

  const handleUpdate = () => {
    const pattern = /^([$]?[a-zA-Z0-9_.-]+):((?:(?:=|>|>=|<|<=|!=){0,1}["']?(?:[$]?[a-zA-Z0-9_.-:*]+)["']?)?(?:,(?:(?:=|>|>=|<|<=|!=){0,1}["']?(?:[$]?[a-zA-Z0-9_.-:*]+)["']?)?)*)$/;
    const pattern2 = /^(?:(=|>|>=|<|<=|!=){0,1}(["']?(?:[$]?[a-zA-Z0-9_.-:*]+)["']?))$/;

    const parts = value.split(' ')
    const updatedItems = [...items];
    const updatedRemain: string[] = []
    for (let i = 0; i < parts.length - 1; i += 1) {
      const matches = parts[i].match(pattern)
      if (matches && matches.length > 1) {
        const key = matches[1];
        let item: SearchItem = { name: key, values: [] }
        const parts2 = matches[2].split(',');
        for (let i = 0; i < parts2.length; i += 1) {
          const matches2 = parts2[i].trim().match(pattern2)
          if (matches2 && matches2.length > 1) {
            item.values.push({ symbol: matches2[1] || "=", value: matches2[2] })
          }
        }
        updatedItems.push(item)
      } else {
        updatedRemain.push(parts[i])
      }
    }
    updatedRemain.push(parts[parts.length - 1])

    setItems(updatedItems);
    setRemain(updatedRemain.join(' '))
  }
  useEffect(() => {
    handleUpdate();
  }, [value]);

  return { items, remain, clearItem, clearItems };
}