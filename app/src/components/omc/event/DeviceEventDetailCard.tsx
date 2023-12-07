import React from 'react';
import { ProCard } from '@ant-design/pro-components';
import { Descriptions, Spin, Divider } from 'antd';
import { formatTimestamp2 } from '@/utils/format';

type Props = {
  record?: API.DeviceEvent,
  loading: boolean,
};

const DeviceEventDetailCard: React.FC<Props> = ({
  record,
  loading,
}) => {
  const events = record?.MetaData as unknown as Record<string, string> || {};

  return (
    <Spin spinning={loading}>
      {record && (
        <>
          <ProCard gutter={8} title={"Event: " + record?.EventType} style={{ marginBlockStart: 8 }}>
            <ProCard colSpan={16} layout="center" direction="column">
              <ProCard title="" type="inner"  >
                <Descriptions title="Network Element" column={1}>
                  <Descriptions.Item label="OUI">{record?.Oui || ''}</Descriptions.Item>
                  <Descriptions.Item label="Product Class">{record?.ProductClass || ''}</Descriptions.Item>
                  <Descriptions.Item label="Serial Number">{record?.SerialNumber || ''}</Descriptions.Item>
                </Descriptions>
              </ProCard>
              <Divider style={{ margin: '0' }} />
              <ProCard title="" type="inner" >
                <Descriptions title="Event Information" column={1}>
                  <Descriptions.Item label="Event Type">{record?.EventType || ''}</Descriptions.Item>
                  <Descriptions.Item label="Event Time">{formatTimestamp2(record?.Time)}</Descriptions.Item>
                  <Descriptions.Item label="Inform Events">
                    {Object.entries(events).map(([key]) => (
                      <span key={key}>{key}</span>
                    ))}
                  </Descriptions.Item>
                  <Descriptions.Item label="Current Time">{formatTimestamp2(record?.CurrentTime)}</Descriptions.Item>
                </Descriptions>
              </ProCard>
            </ProCard>
          </ProCard>
        </>)
      }
      {!loading && !record && (
        <div>Failed to fetch event detail</div>
      )}
    </Spin>
  );
};

export default DeviceEventDetailCard;
