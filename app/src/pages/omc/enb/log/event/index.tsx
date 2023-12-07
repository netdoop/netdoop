import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import { DeviceEventInfoTable } from '@/components/omc/event';

const DeviceEventsPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <DeviceEventInfoTable productType='enb'/>
      </ProCard>

    </PageContainer >
  );
};

export default DeviceEventsPage;
