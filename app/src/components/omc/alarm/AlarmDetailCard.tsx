import React from 'react';
import { ProCard } from '@ant-design/pro-components';
import { Descriptions, Spin, Divider, Tag } from 'antd';
import { SEVERITY_COLOR_MAP } from '@/models/alarm';
import { formatTimestamp2 } from '@/utils/format';

type Props = {
  record?: API.DeviceAlarm,
  loading: boolean,
};

const AlarmDetailCard: React.FC<Props> = ({
  record,
  loading,
}) => {
  const severityTag = record?.PerceivedSeverity ? (
    <Tag color={SEVERITY_COLOR_MAP[record.PerceivedSeverity]}>
      {record.PerceivedSeverity}
    </Tag>
  ) : undefined;
  return (
    <Spin spinning={loading}>
      {record && (
        <>
          <ProCard gutter={8} title={"Alarm: " + record?.AlarmIdentifier} style={{ marginBlockStart: 8 }}>
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
                <Descriptions title="Alarm Information" column={1}>
                  <Descriptions.Item label="Alarm Identifier">{record?.AlarmIdentifier || ''}</Descriptions.Item>
                  <Descriptions.Item label="Perceived Severity">{severityTag || ''}</Descriptions.Item>
                  <Descriptions.Item label="Event Type">{record?.EventType || ''}</Descriptions.Item>
                  <Descriptions.Item label="Alarm Raised Time">{formatTimestamp2(record?.AlarmRaisedTime)}</Descriptions.Item>
                  <Descriptions.Item label="Alarm Changed Time">{formatTimestamp2(record?.AlarmChangedTime)}</Descriptions.Item>
                  <Descriptions.Item label="Alarm Cleared">{record?.AlarmCleared ? 'Yes' : 'No'}</Descriptions.Item>
                  <Descriptions.Item label="Alarm Cleared Time">{formatTimestamp2(record?.AlarmClearedTime)}</Descriptions.Item>
                  <Descriptions.Item label="Alarm Confirmed">{record?.AlarmConfirmed ? 'Yes' : 'No'}</Descriptions.Item>
                  <Descriptions.Item label="Alarm Confirmed Time">{formatTimestamp2(record?.AlarmConfirmedTime)}</Descriptions.Item>
                  <Descriptions.Item label="Probable Cause">{record?.ProbableCause || ''}</Descriptions.Item>
                  <Descriptions.Item label="Additional Information">{record?.AdditionalInformation || ''}</Descriptions.Item>
                  <Descriptions.Item label="Additional Text">{record?.AdditionalText || ''}</Descriptions.Item>
                  <Descriptions.Item label="Managed Object Instance">{record?.ManagedObjectInstance || ''}</Descriptions.Item>
                </Descriptions>
              </ProCard>
            </ProCard>
          </ProCard>
        </>)
      }
      {!loading && !record && (
        <div>Failed to fetch alarm detail</div>
      )}
    </Spin>
  );
};

export default AlarmDetailCard;
