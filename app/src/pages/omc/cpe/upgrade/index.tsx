import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import FirmwarePage from './firmware';

const InventoryPage: React.FC = () => {
  return (
    <PageContainer
      header={{
        title: '',
      }}
      tabList={[
        // {
        //   tab: 'Upgrade',
        //   key: 'upgrade',
        //   children: <></>,
        // },
        {
          tab: 'Firmware',
          key: 'firmwares',
          children: <FirmwarePage />,
        },
      ]}
    />
  );
};

export default InventoryPage;
