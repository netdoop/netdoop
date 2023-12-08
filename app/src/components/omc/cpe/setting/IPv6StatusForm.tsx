import { Descriptions } from "antd";
import { ProCard } from '@ant-design/pro-components';

interface IPv6StatusFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const IPv6StatusForm: React.FC<IPv6StatusFormProps> = ({
  data,
}) => {
  const ipv6Status:Record<string, any> = data?.WEB_GUI?.IPv6?.Status || {};
  return (
    <ProCard split="horizontal">
      <ProCard title="IPv6 Information">
        <Descriptions title="" column={1}>
          <Descriptions.Item label="IPv6 Status">{ipv6Status?.Information?.Status || '-'}</Descriptions.Item>
          <Descriptions.Item label="WAN Connection Type">{ipv6Status?.Information?.Type || '-'}</Descriptions.Item>
          <Descriptions.Item label="IPv6 MGMT Global Address">{ipv6Status?.Information?.GlobalAddress || '-'}</Descriptions.Item>
        </Descriptions>
      </ProCard>
      <ProCard title="LAN Address">
        <Descriptions title="" column={1}>
          <Descriptions.Item label="IPv6 DATA Global Address">{ipv6Status?.LANAddress?.GlobalAddress || '-'}</Descriptions.Item>
          <Descriptions.Item label="IPv6 Link-Local Address">{ipv6Status?.LANAddress?.LocalAddress || '-'}</Descriptions.Item>
          <Descriptions.Item label="IPv6 Status">{ipv6Status?.LANAddress?.Type || '-'}</Descriptions.Item>
        </Descriptions>
      </ProCard>
    </ProCard>
  );
};

export default IPv6StatusForm;
