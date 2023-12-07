import React from 'react';
import { history } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
import { Tooltip, Button } from 'antd';
import { CloseOutlined } from '@ant-design/icons';
import { useParams } from '@umijs/max';
import { DeviceMethodCallDetailCard } from '@/components/omc/method-call';
import { useDeviceMethodCall } from '@/models/device_method_call';


const MethodCallDeail: React.FC = () => {
  const params = useParams();
  const { methodCall, loading } = useDeviceMethodCall(Number(params.ts));
  const handleClose = (event: any) => {
    event.preventDefault();
    history.back()
  };
  return (
    <PageContainer
      header={{
        title: '',
      }}
      extra={[
        <Tooltip title="close" key="close">
          <Button type="default" shape="circle" icon={<CloseOutlined />} onClick={handleClose} />
        </Tooltip>,
      ]}
    >
      <DeviceMethodCallDetailCard record={methodCall} loading={loading} />
    </PageContainer >
  );
};

export default MethodCallDeail;
