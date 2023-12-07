import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import { KpiMeasTable } from '@/components/omc/kpi';

const KpiMgmtPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <KpiMeasTable />
    </PageContainer >
  );
};

export default KpiMgmtPage;
