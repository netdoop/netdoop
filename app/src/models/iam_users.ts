import { useEffect, useState } from 'react';
import services from '@/services/netdoop';
import message from 'antd/es/message';
import { FetchParams, getOrderByString, getSearchItemsString } from './common';

export const fetchUsers = async (values: FetchParams): Promise<API.listUsersData> => {
  try {
    const q = getSearchItemsString(values.searchItems || [])
    const order = getOrderByString(values.sort || {})
    const result = await services.users.listUsers({
      page: values.current,
      page_size: values.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch users');
    return { Data: [], Total: 0 };
  }
};

export const useUsers = (values: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [data, setData] = useState<API.User[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchUsers(values)
      setTotal(result.Total || 0)
      setData(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch devices');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [values]);
  return { data, total, loading, reload };
}

export const useUser = (id: number | undefined) => {
  const [data, setData] = useState<API.User | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      if (id) {
        const data = await services.users.getUser({ id });
        setData(data);
      }
    } catch (error) {
      console.error('Failed to fetch user');
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

export const createUser = async (values: API.createUserBody) => {
  try {
    await services.users.createUser(values);
  } catch (error) {
    message.error('Failed to create user');
  }
};

export const deleteUser = async (record: API.User) => {
  try {
    if (record.Id) {
      await services.users.deleteUser({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable user');
  }
};

export const updateUser = async (record: API.User, values: API.updateUserBody) => {
  try {
    if (record.Id) {
      await services.users.updateUser({ id: record.Id }, values);
    }
  } catch (error) {
    message.error('Failed to update user');
  }
};

export const setUserEnable = async (record: API.User) => {
  try {
    if (record.Id) {
      await services.users.setUserEnable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to enable user');
  }
};

export const setUserDisable = async (record: API.User) => {
  try {
    if (record.Id) {
      await services.users.setUserDisable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable user');
  }
};

export const setRolesForUser = async (record: API.User, values: API.setUserRoles) => {
  try {
    if (record.Id) {
      await services.users.setUserRoles({ id: record.Id }, values);
    }
  } catch (error) {
    message.error('Failed to set roles for user');
  }
};

export const changeUserPassword = async (record: API.User, values: API.changePasswordBody) => {
  try {
    if (record.Id) {
      await services.users.changeUserPassword({ id: record.Id }, values);
    }
  } catch (error) {
    message.error('Failed to change user password');
  }
};


export const changeCurrentPassword = async (values: API.changePasswordBody) => {
  try {
      await services.current.changeCurrentPassword(values);
  } catch (error) {
    message.error('Failed to change password');
  }
};