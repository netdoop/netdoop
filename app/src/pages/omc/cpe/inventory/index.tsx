import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import DevicesPage from './devices';
import DeletedDevicesPage from './deleted-devices';

const InventoryPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}
      tabList={[
        {
          tab: 'Devices',
          key: 'devices',
          children: <DevicesPage />,
        },
        {
          tab: 'Deleted Devices',
          key: 'deleted-devices',
          children: <DeletedDevicesPage  />,
        },
      ]}
    />
  );
};

export default InventoryPage;
