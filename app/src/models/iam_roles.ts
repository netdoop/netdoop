import { useEffect, useState } from 'react';
import services from '@/services/netdoop';
import message from 'antd/es/message';
import { FetchParams, getOrderByString, getSearchItemsString } from './common';

export const fetchRoles = async (params: FetchParams): Promise<API.listRolesData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.roles.listRoles({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch roles');
    return { Data: [], Total: 0 };
  }
};

export const useAllRoles = () => {
  const [total, setTotal] = useState<number>(0);
  const [roles, setRoles] = useState<API.Role[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const params ={
        current: 0,
        pageSize: 100000,
      }
      const result = await fetchRoles(params)
      setTotal(result.Total || 0)
      setRoles(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch devices');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, []);
  return { roles, total, loading, reload };
}

export const useRoles = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [roles, setRoles] = useState<API.Role[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchRoles(params)
      setTotal(result.Total || 0)
      setRoles(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch devices');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);
  return { roles, total, loading, reload };
}

export const useRole = (id: number | undefined) => {
  const [data, setData] = useState<API.Role | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      if (id) {
        const data = await services.roles.getRole({ id });
        setData(data);
      }
    } catch (error) {
      console.error('Failed to fetch role');
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

export const createRole = async (params: API.createRoleBody) => {
  try {
    await services.roles.createRole(params);
  } catch (error) {
    message.error('Failed to create role');
  }
};

export const deleteRole = async (record: API.Role) => {
  try {
    if (record.Id) {
      await services.roles.deleteRole({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable role');
  }
};

export const updateRole = async (record: API.Role, params: API.updateRoleBody) => {
  try {
    if (record.Id) {
      await services.roles.updateRole({ id: record.Id }, params);
    }
  } catch (error) {
    message.error('Failed to update role');
  }
};

export const setRoleRules = async (record: API.Role, params: API.setRoleApiRulesBody) => {
  try {
    if (record.Id) {
      await services.roles.setRoleRules({ id: record.Id }, params);
    }
  } catch (error) {
    message.error('Failed to set role rules');
  }
};

export const setRoleEnable = async (record: API.Role) => {
  try {
    if (record.Id) {
      await services.roles.setRoleEnable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to enable role');
  }
};

export const setRoleDisable = async (record: API.Role) => {
  try {
    if (record.Id) {
      await services.roles.setRoleDisable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable role');
  }
};