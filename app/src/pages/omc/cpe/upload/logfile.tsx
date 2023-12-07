import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { DeviceUploadFileTable } from '@/components/omc/upload-file';

const DeviceLogFilesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <DeviceUploadFileTable productType="cpe" type="2 Vendor Log File" />
      </ProCard>

    </PageContainer >
  );
};

export default DeviceLogFilesPage;
