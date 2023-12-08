import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import { AlarmInfoTable } from '@/components/omc/alarm';

const DevicesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <AlarmInfoTable />
      </ProCard>

    </PageContainer >
  );
};

export default DevicesPage;
