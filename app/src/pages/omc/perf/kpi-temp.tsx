import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { KpiTemplateTable } from '@/components/omc/kpi';

const KpiMgmtPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <ProCard colSpan={24} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
          <KpiTemplateTable/>
        </ProCard>
      </ProCard>
    </PageContainer >
  );
};

export default KpiMgmtPage;
