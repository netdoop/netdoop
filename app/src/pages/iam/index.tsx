import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import UsersPage from './users';
import RolesPage from './roles';

const IAMPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}
      tabList={[
        {
          tab: 'Users',
          key: 'users',
          children: <UsersPage />,
        },
        {
          tab: 'Roles',
          key: 'roles',
          children: <RolesPage />,
        },
      ]}
    />
  );
};

export default IAMPage;
