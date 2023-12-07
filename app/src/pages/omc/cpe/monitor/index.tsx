import React, { } from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import { DeviceStatusTable } from '@/components/omc/cpe';

const DevicesPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
          <DeviceStatusTable groupIds={[]} />
      </ProCard>

    </PageContainer >
  );
};

export default DevicesPage;
