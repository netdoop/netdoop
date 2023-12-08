import React from 'react';
import { PageContainer } from '@ant-design/pro-components';

import { DeletedDeviceInfoTable } from '@/components/omc/enb';

const DevicesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <DeletedDeviceInfoTable />
    </PageContainer >
  );
};

export default DevicesPage;
