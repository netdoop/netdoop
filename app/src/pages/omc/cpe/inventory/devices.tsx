import React from 'react';
import { PageContainer } from '@ant-design/pro-components';

import { DeviceInfoTable } from '@/components/omc/cpe';

const DevicesPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <DeviceInfoTable />
    </PageContainer >
  );
};

export default DevicesPage;
