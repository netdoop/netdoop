import { useEffect, useState } from 'react';

import services from '@/services/netdoop';
import { message } from 'antd';

const mapToList = (groups: API.Group[]): API.Group[] => {
  let result: API.Group[] = [];
  groups.forEach(group => {
    result.push(group);
    if (group.Children) {
      result = result.concat(mapToList(group.Children));
    }
  });
  return result;
}

export const fetchGroups = async (): Promise<API.Group[]> => {
  try {
    const result = await services.omcGroups.listGroups();
    return result;
  } catch (error) {
    message.error('Failed to fetch groups');
    return [];
  }
};

export const useGroups = () => {
  const [groups, setGroups] = useState<API.Group[]>([]);
  const [groupsList, setGroupsList] = useState<API.Group[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const groupById = (id: number | undefined): API.Group | undefined => {
    let group: API.Group | undefined = undefined
    if (id) {
      groupsList?.forEach(item => {
        if (item.Id === id) {
          group = item
          return
        }
      });
    }
    return group
  }

  const groupNameById = (id: number | undefined): string => {
    const group = groupById(id)
    if (group?.Name) {
      return group.Name
    }
    return ''
  }

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchGroups()
      setGroups(result)
      const _groupsList = mapToList(result)
      setGroupsList(_groupsList)
    } catch (error) {
      message.error('Failed to fetch groups');
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    reload()
  }, []);

  return {
    reload,
    loading,
    groups,
    groupsList,
    groupById,
    groupNameById,
  };
};

export const createGroup = async (params: API.createGroupBody) => {
  try {
    await services.omcGroups.createGroup(params);
  } catch (error) {
    message.error('Failed to create group');
  }
};

export const deleteGroup = async (record: API.Group) => {
  try {
    if (record.Id) {
      await services.omcGroups.deleteGroup({id:record.Id});
    }
  } catch (error) {
    message.error('Failed to create group');
  }
};


export const setGroupParent = async (record: API.Group, params: API.setGroupParentBody) => {
  try {
    if (record.Id) {
      await services.omcGroups.setGroupParent({id:record.Id}, params);
    }
  } catch (error) {
    message.error('Failed to set group parent');
  }
};