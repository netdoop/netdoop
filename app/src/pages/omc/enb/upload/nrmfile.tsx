import React from 'react';
import { PageContainer, ProCard } from '@ant-design/pro-components';
import { DeviceUploadFileTable } from '@/components/omc/upload-file';

const DeviceNRMFilesPage: React.FC = () => {

  return (
    <PageContainer
      header={{
        title: '',
      }}>
      <ProCard gutter={4} style={{}}>
        <DeviceUploadFileTable productType="enb" type="NrmFile" />
      </ProCard>

    </PageContainer >
  );
};

export default DeviceNRMFilesPage;
