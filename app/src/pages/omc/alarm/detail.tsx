import React from 'react';
import { history } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
import { Tooltip, Button } from 'antd';
import { CloseOutlined } from '@ant-design/icons';
import { useDeviceAlarm } from '@/models/device_alarm';
import { useParams } from '@umijs/max';
import { AlarmDetailCard } from '@/components/omc/alarm';

const AlarmDetail: React.FC = () => {
  const params = useParams();
  const { alarm, loading } = useDeviceAlarm(Number(params.ts));
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
      <AlarmDetailCard record={alarm} loading={loading} />
    </PageContainer >
  );
};

export default AlarmDetail;
