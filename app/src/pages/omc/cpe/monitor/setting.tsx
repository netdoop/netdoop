import React, { } from 'react';
import { history } from '@umijs/max';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { Spin, Tabs } from 'antd';
import { useParams } from '@umijs/max';

import { useDevice } from '@/models/device';
import {
  OnlineUpgradeForm,
  WANSettingsForm,
  NRLTESettingsForm,
  ScanModeForm,
  APNManagementForm,
  LANSettingsForm,
  DMZSettingsForm,
  ManagementServerForm,
  StaticRouteForm,
  VPNSettingsForm,
  IPv6StatusForm,
  IPv6WANSettingsForm,
  IPv6LANSettingsForm,
  RebootSchedulerForm,
  DateAndTimeSettingsForm,
  DDNSSettingsForm,
  IPPingDiagnosisForm,
  TraceRouteDiagnosisForm,
  SyslogSettingsForm,
} from '@/components/omc/cpe';

const DeviceSetting: React.FC = () => {
  const params = useParams();
  const { device, loading, parameterValuesObject } = useDevice(Number(params.id));

  const handleClose = (event: any) => {
    event.preventDefault();
    history.push('/omc/cpe/monitor')
  };

  const handleTabChange = (key: string, subKey?: string) => {
    if (subKey) {
      history.push(`/omc/cpe/monitor/${params.id}/setting/${key}/${subKey}`);
    } else {
      if (key === 'Network') {
        history.push(`/omc/cpe/monitor/${params.id}/setting/${key}/WANSetting`);
      } else if (key === 'Upgrade') {
        history.push(`/omc/cpe/monitor/${params.id}/setting/${key}/OnlineUpgradeSetting`);
      } else if (key === 'VPN') {
        history.push(`/omc/cpe/monitor/${params.id}/setting/${key}/VPNSetting`);
      } else if (key === 'IPv6') {
        history.push(`/omc/cpe/monitor/${params.id}/setting/${key}/IPv6Status`);
      } else if (key === 'System') {
        history.push(`/omc/cpe/monitor/${params.id}/setting/${key}/TR069`);
      }
    }
  };


  const networkTabItems = [
    {
      key: "WANSetting",
      label: "WAN Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <WANSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "NRLTESetting",
      label: "NR/LTE Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <NRLTESettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "ScanMode",
      label: "Scan Mode",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <ScanModeForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "APNManagement",
      label: "APN Management",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <APNManagementForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    // {
    //   key: "PIN Management",
    //   label: "PIN Management",
    //   children: (
    //     <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
    //       <ProCard colSpan='60%' direction="column">
    //         <PINManagementForm data={parameterValuesObject} />
    //       </ProCard>
    //     </ProCard>
    //   ),
    // },
    // {
    //   key: "SIM Lock",
    //   label: "SIM Lock",
    //   children: (
    //     <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
    //       <ProCard colSpan='60%' direction="column">
    //         <SIMLockForm data={parameterValuesObject} />
    //       </ProCard>
    //     </ProCard>
    //   ),
    // },
    {
      key: "LAN Setting",
      label: "LAN Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <LANSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "Static Route",
      label: "Static Route",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <StaticRouteForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "DMZ Setting",
      label: "DMZ Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <DMZSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const upgradeTabItems = [
    {
      key: "OnlineUpgradeSetting",
      label: "Online Upgrade Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <OnlineUpgradeForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const vpnTabItems = [
    {
      key: "VPNSetting",
      label: "VPN Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <VPNSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const ipv6TabItems = [
    {
      key: "IPv6Status",
      label: "IPv6 Status",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <IPv6StatusForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "IPv6WANSettings",
      label: "IPv6 WAN Settings",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <IPv6WANSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "IPv6LANSettings",
      label: "IPv6 LAN Settings",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <IPv6LANSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const systemTabItems = [
    {
      key: "ScheduleReboot",
      label: "Schedule Reboot",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <RebootSchedulerForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "TR069",
      label: "TR069",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <ManagementServerForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "DateAndTime",
      label: "Date & Time",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <DateAndTimeSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "DDNS",
      label: "DDNS",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <DDNSSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "Syslog",
      label: "Syslog",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <SyslogSettingsForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "PingDiagnosis",
      label: "Ping Diagnosis",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <IPPingDiagnosisForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "TraceRouteDiagnosis",
      label: "TraceRoute Diagnosis",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <TraceRouteDiagnosisForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const tabItems = [
    {
      key: "Network",
      label: "Network",
      children: <Tabs defaultActiveKey={params.sub || "WANSettings"} tabPosition="top" items={networkTabItems} onChange={(v) => handleTabChange("Network", v)} />
    },
    {
      key: "Upgrade",
      label: "Upgrade",
      children: <Tabs defaultActiveKey={params.sub || "OnlineUpgradeSettings"} tabPosition="top" items={upgradeTabItems} onChange={(v) => handleTabChange("Upgrade", v)} />
    },
    {
      key: "VPN",
      label: "VPN",
      children: <Tabs defaultActiveKey={params.sub || "VPNSettings"} tabPosition="top" items={vpnTabItems} onChange={(v) => handleTabChange("VPN", v)} />
    },
    {
      key: "IPv6",
      label: "IPv6",
      children: <Tabs defaultActiveKey={params.sub || "IPv6Status"} tabPosition="top" items={ipv6TabItems} onChange={(v) => handleTabChange("IPv6", v)} />
    },
    {
      key: "System",
      label: "System",
      children: <Tabs defaultActiveKey={params.sub || "TR069"} tabPosition="top" items={systemTabItems} onChange={(v) => handleTabChange("System", v)} />
    },
  ];

  return (
    <PageContainer
      onBack={handleClose}
      header={{
        title: "Setting: " + device?.Oui + "-" + device?.SerialNumber,
      }}
    >
      <Spin spinning={loading}>
        {device && (<Tabs
          tabPosition="left"
          items={tabItems}
          defaultActiveKey={params.key || "Network"}
          onChange={(v) => handleTabChange(v)}
        />)}
        {!loading && !device && (
          <div>Failed to fetch device information</div>
        )}
      </Spin>
    </PageContainer >
  );
};



export default DeviceSetting;
