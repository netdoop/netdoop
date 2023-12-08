import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { DevicePerfMgmtTable } from '@/components/omc/enb';

const PerfMgmtPage: React.FC = ({
}) => {
  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <ProCard colSpan={24} layout="default" direction="column" bodyStyle={{ paddingInline: 2 }}>
          <DevicePerfMgmtTable groupIds={[]} />
        </ProCard>
      </ProCard>
    </PageContainer >
  );
};

export default PerfMgmtPage;
