import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { DeviceUploadFileTable } from '@/components/omc/upload-file';

const DeviceConfigurationFilesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <DeviceUploadFileTable productType="cpe" type="1 Vendor Configuration File" />
      </ProCard>
    </PageContainer >
  );
};

export default DeviceConfigurationFilesPage;
