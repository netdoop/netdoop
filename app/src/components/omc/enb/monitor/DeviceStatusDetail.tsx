import React, { useEffect, useState } from 'react';
import { ProCard } from '@ant-design/pro-components';
import { Descriptions, Spin, Tabs } from 'antd';
import { useDevice } from '@/models/device';
import { formatTimestamp2 } from '@/utils/format';
import dayjs, { Dayjs } from 'dayjs';
import { DevicePmValueChart } from '../../device';
import { DateRangePick } from '../../common';
import { useDeviceParameterValues } from '@/models/device_params';

interface Props {
  id: number;
};

const DeviceStatusDetail: React.FC<Props> = ({ id }) => {
  const [timeRange, setTimeRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([dayjs(), dayjs()]);
  const handleTimeRangeChange = (value: [Dayjs, Dayjs]) => {
    setTimeRange(value);
  }

  const { device, deviceParameterValues, loading } = useDevice(id);
  const { getObject } = useDeviceParameterValues(deviceParameterValues);

  const [deviceProperties, setDeviceProperties] = useState<{ Latitude: string, Longitude: string, Height: string }>();
  const [deviceInfo, setDeviceInfo] = useState<Record<string, any>|undefined>({})
  const [enbStatus, setEnbStatus] = useState<Record<string, any>|undefined>({})
  const [wanEthernetInterfaceConfig, setWanEthernetInterfaceConfig] = useState<Record<string, any>|undefined>({})
  const [wanIPConnection, setWanIPConnection] = useState<Record<string, any>|undefined>({})
  const [capabilities, setCapabilities] = useState<Record<string, any>|undefined>({})
  const [fapControl, setFapControl] = useState<Record<string, any>|undefined>({})
  const [epc, setEpc] = useState<Record<string, any>|undefined>({})
  const [rf, setRf] = useState<Record<string, any>|undefined>({})
  const [phy, setPhy] = useState<Record<string, any>|undefined>({})
  const [mmePoolConfigParam, setMmePoolConfigParam] = useState<Record<string, any>|undefined>({})

  useEffect(()=>{
    setDeviceProperties(device.Properties as unknown as { Latitude: string, Longitude: string, Height: string })
    setDeviceInfo(getObject('.DeviceInfo.'));
    setEnbStatus(getObject('.DeviceInfo.X_VENDOR_ENODEB_STATUS.'));
    setWanEthernetInterfaceConfig(getObject('.WANDevice.1.WANEthernetInterfaceConfig.'));
    setWanIPConnection(getObject('.WANDevice.1.WANConnectionDevice.1.WANIPConnection.1.'));
    setCapabilities(getObject('.Services.FAPService.1.Capabilities.'));
    setFapControl(getObject('.Services.FAPService.1.FAPControl.'));
    setEpc(getObject('.Services.FAPService.1.CellConfig.LTE.EPC.'));
    setRf(getObject('.Services.FAPService.1.CellConfig.LTE.RAN.RF.'));
    setPhy(getObject('.Services.FAPService.1.CellConfig.LTE.RAN.PHY.'));
    setMmePoolConfigParam(getObject('.Services.FAPService.1.CellConfig.LTE.MmePoolConfigParam.1.'));
  }, [device])

  const chartTabItems = [
    {
      key: "ULDLPRBUtilization",
      label: "UL/DL PRB Utilization",
      children: <DevicePmValueChart names={["RRU.DownlinkPRBUtilizationRate", "RRU.UplinkPRBUtilizationRate"]}
        productType='enb' device={device} unit='(%)'
        dateRange={timeRange} height={320} />
    },
    {
      key: "ERAB.EstablishSuccRate",
      label: "ERAB.EstablishSuccRate",
      children: <DevicePmValueChart names={["ERAB.InitialEstabSuccRate"]}
        productType='enb' device={device} unit='(%)'
        dateRange={timeRange} height={320} />
    },
    {
      key: "HO.InterEnbOutSucc.Rate.S1",
      label: "HO.InterEnbOutS1SuccRate",
      children: <DevicePmValueChart names={["HO.InterEnbOutSucc.Rate.S1"]}
        productType='enb' device={device} unit='(%)'
        dateRange={timeRange} height={320} />
    },
    {
      key: "HO.InterEnbOutSucc.Rate.X2",
      label: "HO.InterEnbOutX2SuccRate",
      children: <DevicePmValueChart names={["HO.InterEnbOutSucc.Rate.X2"]}
        productType='enb' device={device} unit='(%)'
        dateRange={timeRange} height={320} />
    },
    {
      key: "HO.InterEnbOutSucc.Rate",
      label: "HO.InterEnbOutSuccRate",
      children: <DevicePmValueChart names={["HO.InterEnbOutSucc.Rate"]}
        productType='enb' device={device} unit='(%)'
        dateRange={timeRange} height={320} />
    },
    {
      key: "RRC.SetupSuccRate",
      label: "RRC.SetupSuccRate",
      children: <DevicePmValueChart names={["RRC.SetupSuccRate"]}
        productType='enb' device={device} unit='(%)'
        dateRange={timeRange} height={320} />
    },
  ];

  return (
    <>
      <Spin spinning={loading}>
        {device && (
          <>
            <ProCard gutter={8} title="" style={{ marginBlockStart: 8 }}>
              <ProCard colSpan={12} layout="center" direction="column">
                <ProCard title="" type="inner">

                  <ProCard colSpan={24} title="">
                    <Descriptions title="Cell Information" column={2}>
                      <Descriptions.Item label="Cell Name">{device.Name || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Cell Status">{enbStatus?.X_VENDOR_CELL_STATUS || '-'}</Descriptions.Item>
                      <Descriptions.Item label="RF Status">{fapControl?.LTE?.RFTxStatus || '-'}</Descriptions.Item>
                      <Descriptions.Item label="TAC">{epc?.TAC || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Cell ID">{rf?.PhyCellID || '-'}</Descriptions.Item>
                      <Descriptions.Item label="PLMN">{mmePoolConfigParam?.PLMNID || '-'}</Descriptions.Item>
                      <Descriptions.Item label="EARFCN">DL={rf?.EARFCNDL || '-'} | UL={rf?.EARFCNUL || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Bandwidth">DL={rf?.DLBandwidth || '-'} | UL={rf?.ULBandwidth || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Max Tx Power">{capabilities?.MaxTxPower || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Duplex Mode">{wanEthernetInterfaceConfig?.DuplexMode || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Root Sequence Index">{phy?.PRACH?.RootSequenceIndex || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Special Subframe Patterns">{phy?.TDDFrame?.SpecialSubframePatterns || '-'}</Descriptions.Item>
                      <Descriptions.Item label="SubFrame Assignment">{phy?.TDDFrame?.SubFrameAssignment || '-'}</Descriptions.Item>
                    </Descriptions>
                  </ProCard>
                </ProCard>

                <ProCard title="" type="inner">
                  <ProCard colSpan={12} title="">
                    <Descriptions title="Device Information" column={1}>
                      <Descriptions.Item label="Serial Number">{deviceInfo?.SerialNumber || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Model Name">{deviceInfo?.ModelName || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Product Class">{deviceInfo?.ProductClass || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Software Version">{deviceInfo?.SoftwareVersion || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Hardware Version">{deviceInfo?.HardwareVersion || '-'}</Descriptions.Item>

                      <Descriptions.Item label="MAC Address">{wanEthernetInterfaceConfig?.MACAddress || '-'}</Descriptions.Item>

                      <Descriptions.Item label="Up Time">{deviceInfo?.UpTime?.toString() || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Last Inform Time">{formatTimestamp2(device.LastInformTime)}</Descriptions.Item>
                    </Descriptions>


                  </ProCard>
                  <ProCard colSpan={12} title="" >
                    <Descriptions title="Network" column={1}>
                      <Descriptions.Item label="IP Address">{wanIPConnection?.ExternalIPAddress || '-'}</Descriptions.Item>
                    </Descriptions>
                    <Descriptions title="Status" column={1}>
                      <Descriptions.Item label="MME Status">MMEIp1={mmePoolConfigParam?.MMEIp1 || '-'} | MMEIp2={mmePoolConfigParam?.MMEIp2 || '-'}</Descriptions.Item>
                      <Descriptions.Item label="UE Number">{enbStatus?.X_VENDOR_UEACT_NUM || '-'}</Descriptions.Item>
                    </Descriptions>
                    <Descriptions title="Location" column={1}>
                      <Descriptions.Item label="Longitude">{deviceProperties?.Longitude || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Latitude">{deviceProperties?.Latitude || '-'}</Descriptions.Item>
                      <Descriptions.Item label="Height">{deviceProperties?.Height || '-'}</Descriptions.Item>
                    </Descriptions>
                  </ProCard>
                </ProCard>
              </ProCard>
              <ProCard colSpan={12} layout="center" direction="column">
                <ProCard >
                  <DateRangePick onChange={handleTimeRangeChange} />
                </ProCard>
                <ProCard  >
                  <Tabs defaultActiveKey="Network" tabPosition="left" items={chartTabItems} />)
                </ProCard>
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

export default DeviceStatusDetail;
