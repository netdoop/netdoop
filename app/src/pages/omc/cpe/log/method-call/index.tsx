import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import { DeviceMethodCallTable } from '@/components/omc/method-call';

const DeviceMethordCallPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <DeviceMethodCallTable  productType='cpe'/>
      </ProCard>

    </PageContainer >
  );
};

export default DeviceMethordCallPage;
