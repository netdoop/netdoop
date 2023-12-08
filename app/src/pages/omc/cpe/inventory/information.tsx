import React from 'react';
import { history } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
import { useParams } from '@umijs/max';
import { DeviceInfoDetail } from '@/components/omc/cpe';

const DeviceInformationPage: React.FC = () => {
  const params = useParams();

  const handleClose = (event: any) => {
    event.preventDefault();
    history.back()
  };
  return (
    <PageContainer
      title="Information"
      onBack={handleClose}
    >
      <DeviceInfoDetail id={Number(params.id)} />
    </PageContainer >
  );
};

export default DeviceInformationPage;
