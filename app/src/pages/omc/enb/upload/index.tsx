import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import DeviceLogFilesPage from './logfile';
import DeviceConfigurationFilesPage from './cfgfile';
import DevicePmFilesPage from './pmfile';
import DeviceNRMFilesPage from './nrmfile';
import { Tabs } from 'antd';
import { useParams, history } from '@umijs/max';

const UploadPage: React.FC = () => {
  const params = useParams();

  const handleTabChange = (key: string) => {
    history.push(`/omc/enb/upload/${key}`);
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
    {
      label: 'PM',
      key: 'PM',
      children: <DevicePmFilesPage />,
      onChange: () => handleTabChange("Event"),
    },
    {
      label: 'NRM',
      key: 'NRM',
      children: <DeviceNRMFilesPage />,
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
