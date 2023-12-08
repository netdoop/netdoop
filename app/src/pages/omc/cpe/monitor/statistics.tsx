import React, { useEffect, useState } from 'react';
import { history } from '@umijs/max';
import { PageContainer, ProCard, ProList } from '@ant-design/pro-components';
import { Descriptions, Spin, Divider } from 'antd';
import { useDevice } from '@/models/device';
import { useParams } from '@umijs/max';
import { useDeviceParameterValues } from '@/models/device_params';

const DeviceStatistics: React.FC = () => {
  const params = useParams();
  const { device, deviceParameterValues, loading } = useDevice(Number(params.id));
  const { getObject, getObjectList } = useDeviceParameterValues(deviceParameterValues);

  const [cpuUsage, setCPUUsage] = useState<Record<string, any> | undefined>({})
  const [memoryUsage, setMemoryUsage] = useState<Record<string, any> | undefined>({})
  const [rateState, setRateState] = useState<Record<string, any> | undefined>({})
  const [throughputStatisticsList, setThroughputStatisticsList] = useState<Record<string, Record<string, any>>>({})

  useEffect(() => {
    setCPUUsage(getObject('.WEB_GUI.Overview.CPUUsage.'));
    setMemoryUsage(getObject('.WEB_GUI.Overview.MemoryUsage.'));
    setRateState(getObject('.WEB_GUI.Overview.WANStatus.'));
    setThroughputStatisticsList(getObjectList('.WEB_GUI.Overview.ThroughputStatisticsList.') || {});
  }, [device])

  const handleClose = (event: any) => {
    event.preventDefault();
    history.back()
  };

  return (
    <PageContainer
      onBack={handleClose}
      header={{
        title: "Information: " + device?.Oui + "-" + device?.SerialNumber,
      }}
    >
      <Spin spinning={loading}>
        {device && (
          <>
            <ProCard gutter={8} title="" style={{ marginBlockStart: 8 }}>
              <ProCard colSpan={8} layout="center" direction="column">
                <ProCard title="" type="inner"  >
                  <Descriptions title="CPU Usage" column={1}>
                    <Descriptions.Item label="Current">{cpuUsage?.Current || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Max">{cpuUsage?.Max || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Min">{cpuUsage?.Min || '-'}</Descriptions.Item>
                  </Descriptions>
                  <Divider style={{ margin: '0' }} />

                  <Descriptions title="Memory Usage" column={1}>
                    <Descriptions.Item label="Current">{memoryUsage?.Current || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Max">{memoryUsage?.Max || '-'}</Descriptions.Item>
                    <Descriptions.Item label="Min">{memoryUsage?.Min || '-'}</Descriptions.Item>
                  </Descriptions>
                  <Divider style={{ margin: '0' }} />

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
              </ProCard>
              <ProCard colSpan={8} layout="center" direction="column">
                <ProCard title="" type="inner"  >
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
              </ProCard>
              <ProCard colSpan={8} layout="center" direction="column">
                <ProCard >
                  {/* <DateRangePick onChange={handleTimeRangeChange} /> */}
                </ProCard>
                <ProCard  >
                  {/* <DeviceStatsChart dateRange={timeRange} height={200} /> */}
                </ProCard>
              </ProCard>
            </ProCard>
          </>)
        }
        {!loading && !device && (
          <div>Failed to fetch device information</div>
        )}
      </Spin>
    </PageContainer >
  );
};

export default DeviceStatistics;
