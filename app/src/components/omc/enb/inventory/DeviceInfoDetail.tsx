import React, { useEffect, useState } from 'react';
import { ProCard } from '@ant-design/pro-components';
import { Descriptions, Spin, Divider } from 'antd';
import { useDevice } from '@/models/device';
import { useDeviceParameterValues } from '@/models/device_params';

interface Props {
  id: number;
}

const DeviceInfoDetail: React.FC<Props> = ({id}) => {
  const { device, deviceParameterValues, loading } = useDevice(id);
  const { getObject, getParameterValue } = useDeviceParameterValues(deviceParameterValues);

  const [deviceInfo, setDeviceInfo] = useState<Record<string, any>|undefined>({})
  const [deviceSummary, setDeviceSummary] = useState<string|undefined>()

  useEffect(()=>{
    setDeviceInfo(getObject('.DeviceInfo.'));
    setDeviceSummary(getParameterValue('.DeviceSummary'))
  }, [device])

  return (
    <>
      <Spin spinning={loading}>
        {device && (
          <>
            <ProCard gutter={8} title={"Information: " + device.Oui + "-" + device.SerialNumber} style={{ marginBlockStart: 8 }}>
              <ProCard colSpan={8} layout="center" direction="column">
                <ProCard title="" type="inner"  >
                  <Descriptions title="Device Information" column={1}>
                    <Descriptions.Item label="Serial Number">{deviceInfo?.SerialNumber || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Hardware Version">{deviceInfo?.HardwareVersion || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Software Version">{deviceInfo?.SoftwareVersion || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Modem Firmware Version">{deviceInfo?.ModemFirmwareVersion || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Module Version">{deviceInfo?.ModuleVersion || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Spec Version">{deviceInfo?.SpecVersion || '-'}</Descriptions.Item>

                    <Descriptions.Item label="Manufacturer">{deviceInfo?.Manufacturer || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Manufacturer OUI">{deviceInfo?.ManufacturerOUI || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Model Name">{deviceInfo?.ModelName || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Product Class">{deviceInfo?.ProductClass || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Provisioning Code">{deviceInfo?.ProvisioningCode || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Device Summary">{deviceSummary || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Description">{deviceInfo?.Description || '-'}</Descriptions.Item>
                    <Descriptions.Item label="First Use Date">{deviceInfo?.FirstUseDate || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Up Time">{deviceInfo?.UpTime || '-'}</Descriptions.Item>
                  </Descriptions>
                </ProCard>
                <Divider style={{ margin: '0' }} />
                <ProCard title="" type="inner" >

                </ProCard>
              </ProCard>
              <ProCard colSpan={8} layout="center" >

              </ProCard>
              <ProCard colSpan={8} layout="center" >
              </ProCard>
            </ProCard>
          </>)
        }
        {!loading && !device && (
          <div>Failed to fetch device information</div>
        )}
      </Spin>
    </ >
  );
};

export default DeviceInfoDetail;
