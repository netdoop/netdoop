import React from 'react';
import { ProCard } from '@ant-design/pro-components';
import { Descriptions, Spin, Divider } from 'antd';
import { formatTimestamp2 } from '@/utils/format';

type Props = {
  record?: API.DeviceMethodCall,
  loading: boolean,
};

const DeviceMethodCallDetailCard: React.FC<Props> = ({
  record,
  loading,
}) => {

  return (
    <Spin spinning={loading}>
      {record && (
        <>
          <ProCard gutter={8} title={"Event: " + record?.MethodName} style={{ marginBlockStart: 8 }}>
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
                <Descriptions title="Information" column={1}>
                  <Descriptions.Item label="Method Name">{record?.MethodName || ''}</Descriptions.Item>
                  <Descriptions.Item label="Call Time">{formatTimestamp2(record?.Time)}</Descriptions.Item>
                  <Descriptions.Item label="Method State">{record?.State || ''}</Descriptions.Item>

                  <Descriptions.Item label="Fault Code">{record?.FaultCode || ''}</Descriptions.Item>
                  <Descriptions.Item label="Fault String">{record?.FaultString || ''}</Descriptions.Item>
                  {/* <Descriptions.Item label="Request">{record?.RequestValues || ''}</Descriptions.Item> */}
                  {/* <Descriptions.Item label="Response">{record?.ResponseValues || ''}</Descriptions.Item> */}
                </Descriptions>
              </ProCard>
            </ProCard>
          </ProCard>
        </>)
      }
      {!loading && !record && (
        <div>Failed to fetch method call detail</div>
      )}
    </Spin>
  );
};

export default DeviceMethodCallDetailCard;
