import { toNumber } from "lodash";
import { useState, useEffect } from "react";


export const useDeviceParameterValues = (deviceParameterValues: true | Record<string, any> | undefined) => {
  const [parameterValues, setParameterValues] = useState<Record<string, any>>({});
  const [parameterValuesUpdated, setParameterValuesUpdated] = useState<number>(0);
  useEffect(() => {
    const values = deviceParameterValues as unknown as Record<string, any> || {};
    setParameterValues(values)
    setParameterValuesUpdated(parameterValuesUpdated + 1);
  }, [deviceParameterValues])

  const getParameterValue = (name: string): any => {
    if (name.startsWith(".")) {
      const obj = getParameterValue("Device" + name) || getParameterValue("InternetGatewayDevice" + name);
      return obj;
    }
    const keys = Object.keys(parameterValues).filter(key => key === name);
    if (keys.length === 1) {
      return parameterValues[keys[0]];
    }
    return undefined;
  };

  const getObject = (name: string): Record<string, any> | undefined => {
    if (!name.endsWith(".")) {
      return;
    }
    if (name.startsWith(".")) {
      const obj = getObject("Device" + name);
      if (obj) {
        return obj;
      }
      return getObject("InternetGatewayDevice" + name);
    }

    const keys = Object.keys(parameterValues).filter(key => key.startsWith(name));
    const obj: Record<string, any> = {};
    keys.forEach(key => {
      obj[key.slice(name.length)] = parameterValues[key];
    });
    return obj;
  };

  const getObjectList = (name: string): Record<string, Record<string, any>> | undefined => {
    if (!name.endsWith(".")) {
      return undefined
    }
    if (name.startsWith(".")) {
      const objs = getObjectList("Device" + name);
      if (objs && Object.keys(objs).length > 0) {
        return objs;
      }
      return getObjectList("InternetGatewayDevice" + name);
    }

    const objectNames: Record<string, string> = {};
    Object.keys(parameterValues).forEach(key => {
      if (key.startsWith(name)) {
        const parts = key.slice(name.length).split(".");
        if (parts.length > 1 && toNumber(parts[0]) > 0) {
          objectNames[parts[0]] = name + parts[0] + ".";
        }
      }
    })
    const objs: Record<string, Record<string, any>> = {};
    Object.keys(objectNames).forEach(key => {
      const obj = getObject(objectNames[key]);
      if (obj) {
        objs[key] = obj;
      }
    })
    return objs;
  };

  // export const getObjectFromData = (data: Record<string, any>, name: string): Record<string, any> | undefined => {
  //   if (!name.endsWith(".")) {
  //     return;
  //   }
  //   const keys = Object.keys(data).filter(key => key.startsWith(name));
  //   const obj: Record<string, any> = {};
  //   keys.forEach(key => {
  //     obj[key.slice(name.length)] = data[key];
  //   });
  //   return obj;
  // };

  // export const getObjectListFromData = (data: Record<string, any>, name: string): Record<string, Record<string, any>> | undefined => {
  //   if (!name.endsWith(".")) {
  //     return undefined
  //   }
  //   const objectNames: Record<string, string> = {};
  //   Object.keys(data).forEach(key => {
  //     if (key.startsWith(name)) {
  //       const parts = key.slice(name.length).split(".");
  //       if (parts.length > 1 && toNumber(parts[0]) > 0) {
  //         objectNames[parts[0]] = name + parts[0] + ".";
  //       }
  //     }
  //   })
  //   const objs: Record<string, Record<string, any>> = {};
  //   Object.keys(objectNames).forEach(key => {
  //     const obj = getObjectFromData(data, objectNames[key]);
  //     if (obj) {
  //       objs[key] = obj;
  //     }
  //   })
  //   return objs;
  // };

  return {
    parameterValuesUpdated,
    getParameterValue,
    getObject,
    getObjectList,
  };
};

export type ParameterValuesObject = Record<string, any>;

export function parseParameterValues(values: Record<string, any>): ParameterValuesObject {
  const obj: ParameterValuesObject = {};

  Object.keys(values).forEach(key => {
    const path = key.split('.');
    let current = obj;
    if (path.length > 1 && (path[0] === 'Device' || path[0] === 'InternetGatewayDevice')) {
      for (let i = 1; i < path.length - 1; i++) {
        const part = path[i];
        if (!current[part]) {
          current[part] = {};
        }
        current = current[part];
      }
      current[path[path.length - 1]] = values[key];
    }
  });

  return obj;
}
