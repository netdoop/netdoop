import React, { } from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import { AuditLogsTable } from '@/components/iam';

const AuditPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
          <AuditLogsTable/>
      </ProCard>

    </PageContainer >
  );
};

export default AuditPage;
