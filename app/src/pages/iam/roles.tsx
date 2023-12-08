import React, { } from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';

import { RoleTable } from '@/components/iam';

const UserPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <RoleTable />
      </ProCard>

    </PageContainer >
  );
};

export default UserPage;
