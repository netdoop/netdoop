import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import DeviceEventsPage from './event';
import DeviceMethordCallPage from './method-call';
import { Tabs } from 'antd';
import { useParams, history } from '@umijs/max';

const InventoryPage: React.FC = () => {
  const params = useParams();

  const handleTabChange = (key: string) => {
    history.push(`/omc/cpe/log/${key}`);
  };

  const tabItems = [
    {
      label: 'Event',
      key: 'events',
      children: <DeviceEventsPage />,
      onChange: () => handleTabChange("events"),
    },
    {
      label: 'Method Call',
      key: 'method-calls',
      children: <DeviceMethordCallPage />,
      onChange: () => handleTabChange("method-calls"),
    },
  ];
  return (
    <PageContainer
      header={{
        title: '',
      }}
    >
      <Tabs
        items={tabItems}
        defaultActiveKey={params.key || "events"}
        onChange={(v) => handleTabChange(v)}
      />
    </PageContainer>
  );
};

export default InventoryPage;
