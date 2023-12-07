import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import { KpiValueCharts } from '@/components/omc/kpi';

const KpiValueChartsPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <KpiValueCharts />
    </PageContainer >
  );
};

export default KpiValueChartsPage;
