import React, { } from 'react';
import { history } from '@umijs/max';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { Spin, Tabs } from 'antd';
import { useParams } from '@umijs/max';

import { useDevice } from '@/models/device';
import { EPCSettingForm, IPSecSettingForm, MRMgmtSettingForm, ManagementServerForm, NRMMgmtSettingForm, PerfMgmtSettingForm, RFSettingForm, SCTPSettingForm, WANSettingForm } from '@/components/omc/enb';

const DeviceSetting: React.FC = () => {
  const params = useParams();
  const { device, loading, parameterValuesObject } = useDevice(Number(params.id));

  const handleClose = (event: any) => {
    event.preventDefault();
    history.push('/omc/enb/monitor')
  };

  const handleTabChange = (key: string, subKey?: string) => {
    if (subKey) {
      history.push(`/omc/enb/monitor/${params.id}/setting/${key}/${subKey}`);
    } else {
      if (key === 'Network') {
        history.push(`/omc/enb/monitor/${params.id}/setting/${key}/WANSetting`);
      } else if (key === 'HeMS') {
        history.push(`/omc/enb/monitor/${params.id}/setting/${key}/HeMSSetting`);
      } else if (key === 'General') {
        history.push(`/omc/enb/monitor/${params.id}/setting/${key}/EPCSetting`);
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
            <WANSettingForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    // {
    //   key: "LANSetting",
    //   label: "LAN Setting",
    //   children: (
    //     <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
    //       <ProCard colSpan='60%' direction="column">
    //         {/* <WANSettingsForm device={device}  data={parameterValuesObject}/> */}
    //       </ProCard>
    //     </ProCard>
    //   ),
    // },
    {
      key: "IPSecSetting",
      label: "IPSec Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <IPSecSettingForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const hemsTabItems = [
    {
      key: "HeMSSetting",
      label: "HeMS Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <ManagementServerForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "PMSetting",
      label: "PM Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='100%' direction="column">
            <PerfMgmtSettingForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "NRM Setting",
      label: "NRM Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <NRMMgmtSettingForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "MRSetting",
      label: "MR Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <MRMgmtSettingForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const generalTabItems = [
    {
      key: "SyncSetting",
      label: "Sync Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            {/* <IPPingDiagnosisForm device={device} data={IPPingDiagnostics} /> */}
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "EPCSetting",
      label: "EPC Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <EPCSettingForm device={device} data={parameterValuesObject}/>
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "SCTPSetting",
      label: "SCTP Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <SCTPSettingForm device={device} data={parameterValuesObject} />
          </ProCard>
        </ProCard>
      ),
    },
    {
      key: "RFSetting",
      label: "RF Setting",
      children: (
        <ProCard gutter={8} style={{ marginBlockStart: 8 }}>
          <ProCard colSpan='60%' direction="column">
            <RFSettingForm device={device}  data={parameterValuesObject}/>
          </ProCard>
        </ProCard>
      ),
    },
  ];
  const tabItems = [
    {
      key: "Network",
      label: "Network",
      children: <Tabs
        tabPosition="top"
        items={networkTabItems}
        defaultActiveKey={params.sub || "WANSetting"}
        onChange={(v) => handleTabChange("Network", v)}
      />
    },
    {
      key: "HeMS",
      label: "HeMS",
      children: <Tabs
        tabPosition="top"
        items={hemsTabItems}
        defaultActiveKey={params.sub || "HeMSSetting"}
        onChange={(v) => handleTabChange("HeMS", v)}
      />
    },

    {
      key: "General",
      label: "General",
      children: <Tabs
        tabPosition="top"
        items={generalTabItems}
        defaultActiveKey={params.sub || "EPCSetting"}
        onChange={(v) => handleTabChange("General", v)}
      />
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
