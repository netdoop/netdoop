import React, { useEffect, useState } from 'react';
import { ProCard, ProList } from '@ant-design/pro-components';
import { Descriptions, Spin, Divider } from 'antd';
import { useDevice } from '@/models/device';
import DevicePmCharts from './DevicePmCharts';
import { useDeviceParameterValues } from '@/models/device_params';

interface Props {
  id: number;
};

const DeviceStatusDetail: React.FC<Props> = ({ id }) => {
  const { device, deviceParameterValues, loading } = useDevice(id);
  const { getObject, getObjectList } = useDeviceParameterValues(deviceParameterValues);

  const [deviceInfo, setDeviceInfo] = useState<Record<string, any>|undefined>({})
  const [moduleInfo, setModuleInfo] = useState<Record<string, any>|undefined>({})
  const [systemInfo, setSystemInfo] = useState<Record<string, any>|undefined>({})
  const [versionInfo, setVersionInfo] = useState<Record<string, any>|undefined>({})

  const [lteStatus, setLteStatus] = useState<Record<string, any>|undefined>({})
  const [nrStatus, setNrStatus] = useState<Record<string, any>|undefined>({})
  const [isValidENBID, setIsValidENBID] = useState<boolean>(false)
  const [isValidGNBID, setIsValidGNBID] = useState<boolean>(false)

  const [internetStatus, setInternetStatus] = useState<Record<string, any>|undefined>({})
  const [lanStatus, setLANStatus] = useState<Record<string, any>|undefined>({})
  const [wanStatus, setWANStatus] = useState<Record<string, any>|undefined>({})
  const [rateState, setRateState] = useState<Record<string, any>|undefined>({})
  const [throughputStatisticsList, setThroughputStatisticsList] = useState<Record<string, Record<string,any>>>({})


  useEffect(()=>{
    setDeviceInfo(getObject('.DeviceInfo.'));
    setModuleInfo(getObject(".WEB_GUI.Overview.ModuleInfo."));
    setSystemInfo(getObject(".WEB_GUI.Overview.SystemInfo."));
    setVersionInfo(getObject(".WEB_GUI.Overview.VersionInfo."));

    setLteStatus(getObject(".WEB_GUI.Network.NR-LTE.Status.LTE."));
    setNrStatus(getObject(".WEB_GUI.Network.NR-LTE.Status.NR."));
    setIsValidENBID(lteStatus?.eNBID !== '--');
    setIsValidGNBID(nrStatus?.gNBID !== '--');

    setInternetStatus(getObject(".WEB_GUI.Overview.InternetStatus."));
    setLANStatus(getObject(".WEB_GUI.Overview.LANStatus."));
    setWANStatus(getObject(".WANDevice.1.WANConnectionDevice.1.WANIPConnection.1."));
    setRateState(getObject(".WEB_GUI.Overview.WANStatus."));
    setThroughputStatisticsList(getObjectList(".WEB_GUI.Overview.ThroughputStatisticsList.")||{});

  }, [device])

  return (
    <>
      <Spin spinning={loading}>
        {device && (
          <>
            <ProCard gutter={8} title="" style={{ marginBlockStart: 8 }}>
              <ProCard colSpan={7} layout="center" direction="column">
                <ProCard title="" type="inner"  >
                  <Descriptions title="System Information" column={1}>
                    <Descriptions.Item label="Connection Total Time">{systemInfo?.ConnectionTotalTime || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Online Time">{systemInfo?.OnlineTime || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Run Time">{systemInfo?.RunTime || '-'}</Descriptions.Item>
                  </Descriptions>
                  <Divider style={{ margin: '0' }} />

                  <Descriptions title="Module Information" column={1}>
                    <Descriptions.Item label="Model">{moduleInfo?.Model || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Software Version">{moduleInfo?.Version || '-'}</Descriptions.Item>
                    <Descriptions.Item label="IMEI">{moduleInfo?.IMEI || '-'}</Descriptions.Item>
                    <Descriptions.Item label="IMSI">{moduleInfo?.IMSI || '-'}</Descriptions.Item>
                  </Descriptions>
                  <Divider style={{ margin: '0' }} />

                  <Descriptions title="Version Information" column={1}>
                    <Descriptions.Item label="Product Name">{versionInfo?.ProductName || deviceInfo?.ModelName || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Product Model">{versionInfo?.ProductModel || deviceInfo?.ProductClass || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Serial Number">{versionInfo?.SerialNumber || deviceInfo?.SerialNumber || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Hardware Version">{versionInfo?.HardVersion || deviceInfo?.HardwareVersion || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Software Version">{versionInfo?.SoftwareVersion || deviceInfo?.SoftwareVersion || '-'}</Descriptions.Item>
                  </Descriptions>

                  {wanStatus && (
                    <>
                      <Descriptions title="WAN Status" column={1}>
                        <Descriptions.Item label="IP Address">{wanStatus?.ExternalIPAddress || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Subnet Mask">{wanStatus?.SubnetMask || '-'}</Descriptions.Item>
                        <Descriptions.Item label="MAC Address">{wanStatus?.MACAddress || '-'}</Descriptions.Item>
                        <Descriptions.Item label="DNS">{wanStatus?.DNSServers || '-'}</Descriptions.Item>
                      </Descriptions>
                    </>
                  )}
                  {lanStatus && (
                    <>
                      <Descriptions title="LAN Status" column={1}>
                        <Descriptions.Item label="IP Address">{lanStatus?.IPAddress || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Subnet Mask">{lanStatus?.SubnetMask || '-'}</Descriptions.Item>
                        <Descriptions.Item label="MAC Address">{lanStatus?.MACAddress || '-'}</Descriptions.Item>
                      </Descriptions>
                    </>
                  )}
                </ProCard>
              </ProCard>
              <ProCard colSpan={7} layout="center" direction="column">
                <ProCard title="" type="inner" >
                  {internetStatus && (
                    <>
                      <Descriptions title="Internet Status" column={1}>
                        <Descriptions.Item label="Mode">{internetStatus?.Mode || '-'}</Descriptions.Item>
                        {/* <Descriptions.Item label="MAC Address">{InternetStatus?.MACAddress || '-'}</Descriptions.Item> */}
                      </Descriptions>
                    </>
                  )}
                  {lteStatus && isValidENBID && (
                    <>
                      <Descriptions title="LTE Status" column={1}>
                        <Descriptions.Item label="Band">{lteStatus?.Band || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Bandwidth">{lteStatus?.Bandwidth || '-'}</Descriptions.Item>
                        <Descriptions.Item label="CQI">{lteStatus?.CQI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Cell ID">{lteStatus?.CellID || '-'}</Descriptions.Item>
                        <Descriptions.Item label="EARFCN">{lteStatus?.EARFCN || '-'}</Descriptions.Item>
                        <Descriptions.Item label="ECGI">{lteStatus?.ECGI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="eNB ID">{lteStatus?.eNBID || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Frequency">{lteStatus?.Frequency || '-'}</Descriptions.Item>
                        <Descriptions.Item label="PCI">{lteStatus?.PCI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="PLMN">{lteStatus?.PLMN || '-'}</Descriptions.Item>
                        <Descriptions.Item label="RSRP">{lteStatus?.RSRP || '-'}</Descriptions.Item>
                        <Descriptions.Item label="RSRP0">{lteStatus?.RSRP0 || '-'}</Descriptions.Item>
                        <Descriptions.Item label="RSRP1">{lteStatus?.RSRP1 || '-'}</Descriptions.Item>
                        <Descriptions.Item label="RSRQ">{lteStatus?.RSRQ || '-'}</Descriptions.Item>
                        <Descriptions.Item label="RSSI">{lteStatus?.RSSI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Rank">{lteStatus?.Rank || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SINR">{lteStatus?.SINR || '-'}</Descriptions.Item>
                        <Descriptions.Item label="TAC">{lteStatus?.TAC || '-'}</Descriptions.Item>
                        <Descriptions.Item label="TM">{lteStatus?.TM || '-'}</Descriptions.Item>
                        <Descriptions.Item label="TX Power">{lteStatus?.TXPower || '-'}</Descriptions.Item>
                      </Descriptions>
                    </>)}
                  {nrStatus && isValidGNBID && (
                    <>
                      <Descriptions title="NR Status" column={1}>
                        <Descriptions.Item label="gNB ID">{nrStatus?.gNBID || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NCGI">{nrStatus?.NCGI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR ARFCN">{nrStatus?.NR_ARFCN || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR Band">{nrStatus?.NR_Band || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR Bandwidth">{nrStatus?.NR_Bandwidth || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR CQI">{nrStatus?.NR_CQI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR Cell ID">{nrStatus?.NR_CellID || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR DL-MCS">{nrStatus?.NR_DLMCS || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR Frequency">{nrStatus?.NR_Frequency || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR PCI">{nrStatus?.NR_PCI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR PLMN">{nrStatus?.NR_PLMN || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR Rank">{nrStatus?.NR_Rank || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR TAC">{nrStatus?.NR_TAC || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR TX Power">{nrStatus?.NR_TXPower || '-'}</Descriptions.Item>
                        <Descriptions.Item label="NR UL-MCS">{nrStatus?.NR_ULMCS || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB Beam ID">{nrStatus?.SSB_BeamID || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB RSRP">{nrStatus?.SSB_RSRP || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB RSRP0">{nrStatus?.SSB_RSRP0 || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB RSRP1">{nrStatus?.SSB_RSRP1 || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB RSRQ">{nrStatus?.SSB_RSRQ || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB RSSI">{nrStatus?.SSB_RSSI || '-'}</Descriptions.Item>
                        <Descriptions.Item label="SSB SINR">{nrStatus?.SSB_SINR || '-'}</Descriptions.Item>
                      </Descriptions>
                    </>)}
                  <Divider style={{ margin: '0' }} />

                </ProCard>
              </ProCard>
              <ProCard colSpan={10} layout="center" direction="column">
                {false && (<ProCard title="" type="inner"  >
                  <ProList
                    headerTitle="Throughput Statistics List"
                    dataSource={Object.values(throughputStatisticsList || [])}
                    split={false}
                    renderItem={(item) => (
                      <Descriptions title={`Throughput Statistics (${item?.index || '-'})`} column={1}>
                        <Descriptions.Item label="Port">{item?.Port || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Received Dropped">{item?.ReceivedDropped || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Received Errors">{item?.ReceivedErrors || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Received Packets">{item?.ReceivedPackets || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Received Total">{item?.ReceivedTotal || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Sent Dropped">{item?.SentDropped || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Sent Errors">{item?.SentErrors || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Sent Packets">{item?.SentPackets || '-'}</Descriptions.Item>
                        <Descriptions.Item label="Sent Total">{item?.SentTotal || '-'}</Descriptions.Item>
                      </Descriptions>
                    )}
                  />
                  <Divider style={{ margin: '0' }} />
                </ProCard>
                )}
                {false && (<ProCard title="" type="inner"  >
                  <Descriptions title="Rate Info" column={1}>
                    <Descriptions.Item label="DL Rate Current">{rateState?.DLRateCurrent || '-'}</Descriptions.Item>
                    <Descriptions.Item label="DL Rate Max">{rateState?.DLRateMax || '-'}</Descriptions.Item>
                    <Descriptions.Item label="DL Rate Min">{rateState?.DLRateMin || '-'}</Descriptions.Item>
                    <Descriptions.Item label="UL Rate Current">{rateState?.ULRateCurrent || '-'}</Descriptions.Item>
                    <Descriptions.Item label="UL Rate Max">{rateState?.ULRateMax || '-'}</Descriptions.Item>
                    <Descriptions.Item label="UL Rate Min">{rateState?.ULRateMin || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Max DL Throughput">{rateState?.MaxDLThroughput || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Max UL Throughput">{rateState?.MaxULThroughput || '-'}</Descriptions.Item>
                  </Descriptions>
                  <Divider style={{ margin: '0' }} />

                </ProCard>
                )}
                <DevicePmCharts device={device}/>
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
