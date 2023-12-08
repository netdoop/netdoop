import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import { ProductTable } from '@/components/omc/product';

const DevicesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProductTable />
    </PageContainer >
  );
};

export default DevicesPage;
