import { useState } from 'react';
import { history } from '@umijs/max';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import dayjs from 'dayjs';
import { AppstoreOutlined, DesktopOutlined, WarningOutlined } from '@ant-design/icons';

import { IconButton } from '@/components/common';
import { DevicePmValueChart, DeviceStatusStacked } from '@/components/omc/device';
import { AlarmSeverityStatusStacked, AlarmEventTypePie, AlarmIdentifierPie } from '@/components/omc/alarm';
import { DateRangePick, PageHeaderExtra } from '@/components/omc/common';


const HomePage: React.FC = () => {
  const [timeRange, setTimeRange] = useState<[dayjs.Dayjs, dayjs.Dayjs]>([dayjs(), dayjs()]);
  const handleTimeRangeChange = (value: [dayjs.Dayjs, dayjs.Dayjs]) => {
    setTimeRange(value);
  }

  const alarmStatsTabItems = [
    {
      key: "Alarm ID",
      label: "Alarm ID",
      children: (
        <AlarmIdentifierPie dateRange={timeRange} height={240} radius={80} />
      ),
    },
    {
      key: "Event Type",
      label: "Event Type",
      children: (
        <AlarmEventTypePie dateRange={timeRange} height={240} radius={80} />
      ),
    },
  ];

  return (
    <PageContainer extra={[<PageHeaderExtra key="extra" />]}>
      <ProCard split="horizontal" ghost >
        <ProCard split="vertical">
          <ProCard split="vertical" colSpan="320px" gutter={0}>
            <ProCard ><span style={{ display: 'block', marginTop: 12 }}>ENB</span></ProCard>
            <ProCard > <IconButton icon={<AppstoreOutlined />} text='Device' onClick={() => { history.push('/omc/enb/inventory') }} /></ProCard>
            <ProCard > <IconButton icon={<DesktopOutlined />} text='Monitor' onClick={() => { history.push('/omc/enb/monitor') }} /></ProCard>
            <ProCard > <IconButton icon={<WarningOutlined />} text='Alarm' onClick={() => { history.push('/omc/alarms') }} /></ProCard>
          </ProCard>

          <ProCard split="vertical" colSpan="320px" gutter={0}>
            <ProCard ><span style={{ display: 'block', marginTop: 12 }}>CPE</span></ProCard>
            <ProCard > <IconButton icon={<AppstoreOutlined />} text='Device' onClick={() => { history.push('/omc/cpe/inventory') }} /></ProCard>
            <ProCard > <IconButton icon={<DesktopOutlined />} text='Monitor' onClick={() => { history.push('/omc/cpe/monitor') }} /></ProCard>
            <ProCard > <IconButton icon={<WarningOutlined />} text='Alarm' onClick={() => { history.push('/omc/alarms') }} /></ProCard>
          </ProCard>

        </ProCard>
        <ProCard.Divider style={{ marginBlock: 8 }} />

        <ProCard split="horizontal" ghost >
          <ProCard split="vertical">
            <ProCard colSpan="100%" layout="default">
            <DateRangePick onChange={handleTimeRangeChange} />
            </ProCard>
          </ProCard>
        </ProCard>

        <ProCard split="vertical">
          <ProCard title="Active Alarm" colSpan="60%">
            <AlarmSeverityStatusStacked dateRange={timeRange} height={240} />
          </ProCard>
          <ProCard tabs={{ tabPosition: 'top', type: 'card', items: alarmStatsTabItems }}>

          </ProCard>
        </ProCard>
        {/* <ProCard.Divider style={{ marginBlock: 8 }} /> */}
        <ProCard split="horizontal">
          <ProCard split="vertical">
            <ProCard title="eNB Count" colSpan="50%">
              <DeviceStatusStacked productType="enb" dateRange={timeRange} height={240} />
            </ProCard>
            <ProCard title="CPE Count">
              <DeviceStatusStacked productType="cpe" dateRange={timeRange} height={240} />
            </ProCard>
          </ProCard>
        </ProCard>

        <ProCard split="horizontal">
          <ProCard split="vertical">
            <ProCard title="Throughput" colSpan="50%">
              <DevicePmValueChart names={["PDCP.DataVolumeDL", "PDCP.DataVolumeUL"]}
                productType='enb' func='sum'
                unit='(Mbps)' dateRange={timeRange} height={240} />
            </ProCard>
            <ProCard title="PRB Utilization">
              <DevicePmValueChart names={["RRU.UplinkPRBUtilizationRate", "RRU.DownlinkPRBUtilizationRate"]}
                productType='enb' func='sum'
                unit='(%)' dateRange={timeRange} height={240} />
            </ProCard>
          </ProCard>
        </ProCard>
      </ProCard>
    </PageContainer >
  );
};

export default HomePage;