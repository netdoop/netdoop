import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import DeviceLogFilesPage from './logfile';
import DeviceConfigurationFilesPage from './cfgfile';
import { Tabs } from 'antd';
import { useParams, history } from '@umijs/max';

const UploadPage: React.FC = () => {
  const params = useParams();

  const handleTabChange = (key: string) => {
    history.push(`/omc/cpe/upload/${key}`);
  };
  const tabItems = [
    {
      label: 'Configuraions',
      key: 'Configuraions',
      children: <DeviceConfigurationFilesPage />,
      onChange: () => handleTabChange("Event"),
    },
    {
      label: 'Logs',
      key: 'Logs',
      children: <DeviceLogFilesPage />,
      onChange: () => handleTabChange("Event"),
    },
  ];
  return (
    <PageContainer
      header={{
        title: '',
      }}
    >
      <Tabs
        items={tabItems}
        defaultActiveKey={params.key || "Configuraions"}
        onChange={(v) => handleTabChange(v)}
      />
    </PageContainer>
  );
};

export default UploadPage;
