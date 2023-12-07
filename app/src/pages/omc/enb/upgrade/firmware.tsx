import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import FirmwareInfoTable from '@/components/omc/firmware/FirmwareInfoTable';

const DevicesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <FirmwareInfoTable productType='enb' />
      </ProCard>
    </PageContainer >
  );
};

export default DevicesPage;
