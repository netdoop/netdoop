import React from 'react';
import { history } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
import { Tooltip, Button } from 'antd';
import { CloseOutlined } from '@ant-design/icons';
import {useDeviceEvent} from '@/models/device_event';

import { useParams } from '@umijs/max';
import { DeviceEventDetailCard } from '@/components/omc/event';


const EventDetail: React.FC = () => {
  const params = useParams();
  const { event, loading } = useDeviceEvent(Number(params.ts));
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
      <DeviceEventDetailCard record={event} loading={loading} />
    </PageContainer >
  );
};

export default EventDetail;
