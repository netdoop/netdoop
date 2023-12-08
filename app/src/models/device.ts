import { useEffect, useState } from 'react';
import { useAccess } from '@umijs/max';
import services from '@/services/netdoop';
import message from 'antd/es/message';
import { FetchParams, getOrderByString, getSearchItemsString } from './common';
import { ParameterValuesObject, parseParameterValues } from './device_params';

export type DeviceMetaData = {
  SerialNumber?: string;
  MACAddress?: string;
  IMSI?: string;
};
export type DeviceProperties = {
  Latitude?: number;
  Longitude?: number;
  Antitude?: number;

  Height?: number;
  Distance?: number;

  LastInformTime?: number;
  FirstInformTime?: number;
};

export const fetchDevices = async (params: FetchParams): Promise<API.listDevicesData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDevices.listDevices({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch devices');
    return { Data: [], Total: 0 };
  }
};

export const useDevices = (params: FetchParams) => {
  const [total, setTotal] = useState<number>(0);
  const [devices, setDevices] = useState<API.Device[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const reload = async () => {
    setLoading(true);
    try {
      const result = await fetchDevices(params)
      setTotal(result.Total || 0)
      setDevices(result.Data || []);
    } catch (error) {
      message.error('Failed to fetch devices');
    } finally {
      setLoading(false);
    }
  }
  useEffect(() => {
    reload();
  }, [params]);
  return { devices, total, loading, reload };
}

export const useDevice = (id: number | undefined) => {
  const access = useAccess()

  const [device, setDevice] = useState<API.Device>({});
  const [deviceProperties, setDeviceProperties] = useState<DeviceProperties>({});
  const [deviceMetaData, setDeviceMetaData] = useState<DeviceMetaData>({});
  const [deviceParameterValues, setDeviceParameterValues] = useState<Record<string, any>>({});
  const [parameterValuesObject, setParameterValuesObject] = useState<ParameterValuesObject>({});
  const [loading, setLoading] = useState<boolean>(false);

  const fetchDevice = async () => {
    setLoading(true);
    try {
      if (access?.isLogin && id) {
        const device = await services.omcDevices.getDevice({ id });
        const properties = device.Properties as unknown as DeviceProperties;
        const metaData = device.MetaData as unknown as DeviceMetaData;
        const values = device.ParameterValues as unknown as Record<string, any>;
        setDeviceProperties(properties);
        setDeviceMetaData(metaData);
        setDeviceParameterValues(values);
        setParameterValuesObject(parseParameterValues(values));
        setDevice(device);
      }
    } catch (error) {
      console.error('Failed to fetch device:', error);
      setDevice({});
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDevice();
  }, [id]);

  return {
    device,
    deviceParameterValues,
    parameterValuesObject,
    deviceProperties,
    deviceMetaData,
    loading,
  };
};

export const createDevice = async (params: API.createDeviceBody) => {
  try {
    await services.omcDevices.createDevice(params);
  } catch (error) {
    message.error('Failed to create device');
  }
};

export const updateDeviceInfo = async (record: API.Device, params: API.updateDeviceInfoBody) => {
  try {
    if (record.Id) {
      await services.omcDevices.updateDeviceInfo({ id: record.Id }, params);
    }
  } catch (error) {
    message.error('Failed to update device info');
  }
};

export const setGroupForDevice = async (record: API.Device, groupId: number) => {
  try {
    if (record.Id) {
      await services.omcDevices.setGroupForDevice({ id: record.Id }, { GroupId: groupId });
    }
  } catch (error) {
    message.error('Failed to set group for device');
  }
};

export const setDeviceEnable = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDevices.setDeviceEnable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to enable device');
  }
};

export const setDeviceDisable = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDevices.setDeviceDisable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to disable device');
  }
};

export const deleteDevice = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDevices.deleteDevice({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to delete device');
  }
};

export const rebootDevice = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDevices.rebootDevice({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to reboot device');
  }
};


export const getDeviceParameterNames = async (record: API.Device, parameterPath: string, nextLevel: boolean) => {
  try {
    if (record.Id) {
      await services.omcDevices.getDeviceParameterNames({ id: record.Id }, {
        ParameterPath: parameterPath,
        NextLevel: nextLevel,
      });
    }
  } catch (error) {
    message.error('Failed to get device parameter values');
  }
};

export const setDeviceParameterValues = async (record: API.Device, values: Record<string, any>) => {
  try {
    if (record.Id) {
      await services.omcDevices.setDeviceParameterValues({ id: record.Id }, { Values: values });
    }
  } catch (error) {
    message.error('Failed to set device parameter values');
  }
};

export const getDeviceParameterValues = async (record: API.Device, names: string[]) => {
  try {
    if (record.Id) {
      await services.omcDevices.getDeviceParameterValues({ id: record.Id }, { Names: names });
    }
  } catch (error) {
    message.error('Failed to get device parameter values');
  }
};

export const addDeviceObject = async (record: API.Device, objectName: string) => {
  try {
    if (record.Id) {
      await services.omcDevices.addDeviceObject({ id: record.Id }, { ObjectName: objectName });
    }
  } catch (error) {
    message.error('Failed to add device object');
  }
};

export const deleteDeviceObject = async (record: API.Device, objectName: string) => {
  try {
    if (record.Id) {
      await services.omcDevices.deleteDeviceObject({ id: record.Id }, { ObjectName: objectName });
    }
  } catch (error) {
    message.error('Failed to delete device object');
  }
};


export const uploadDeviceFile = async (record: API.Device, fileType: string) => {
  try {
    if (record.Id) {
      await services.omcDevices.uploadDeviceFile({ id: record.Id }, { FileType: fileType });
    }
  } catch (error) {
    message.error('Failed to upload device configure file');
  }
};

export const setDevicePerfEnable = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDevices.setDevicePerfEnable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to set device performance upload enable');
  }
};

export const setDevicePerfDisable = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDevices.setDevicePerfDisable({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to set device performance upload disable');
  }
};

export const upgradeDevice = async (record: API.Device, firmwareID: number) => {
  try {
    if (record.Id) {
      await services.omcDevices.ugradeDevice({ id: record.Id }, { FirmwareID: firmwareID });
    }
  } catch (error) {
    message.error('Failed to upgrade device');
  }
};

export const fetchDeletedDevices = async (params: FetchParams): Promise<API.listDevicesData> => {
  try {
    const q = getSearchItemsString(params.searchItems || [])
    const order = getOrderByString(params.sort || {})
    const result = await services.omcDeletedDevices.listDeletedDevices({
      page: params.current,
      page_size: params.pageSize,
      q: q,
      order_by: order,
    });
    return result;
  } catch (error) {
    message.error('Failed to fetch deleted devices');
    return { Data: [], Total: 0 };
  }
};

export const deleteDeletedDevice = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDeletedDevices.deleteDeletedDevice({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to delete deleted device');
  }
};


export const recoverDeletedDevice = async (record: API.Device) => {
  try {
    if (record.Id) {
      await services.omcDeletedDevices.recoverDeletedDevice({ id: record.Id });
    }
  } catch (error) {
    message.error('Failed to recover deleted device');
  }
};