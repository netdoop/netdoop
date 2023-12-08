import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import { KpiValuesTable } from '@/components/omc/kpi';

const KpiValuePage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <KpiValuesTable />
    </PageContainer >
  );
};

export default KpiValuePage;
