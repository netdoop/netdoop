import { Col, Form, Row, Input, Space, Select } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const typeOptions = [
  { label: 'SLAAC', value: 'SLAAC' },
  { label: 'DHCPv6', value: 'DHCPv6' },
];

interface IPv6LANSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const IPv6LANSettingsForm: React.FC<IPv6LANSettingsFormProps> = ({
  device,
  data,
}) => {
  const prefixName = 'Device.WEB_GUI.IPv6.WAN.';
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {};
    values[prefixName + 'Type'] = data2.Type;
    values[prefixName + 'LocalAddress'] = data2.LocalAddress;
    setDeviceParameterValues(device, values);
  };

  const onReset = () => {
    const ipv6LAN:Record<string, any> = data?.WEB_GUI?.IPv6?.LAN || {};
    form.setFieldsValue(ipv6LAN);
  };
  useEffect(() => {
    const ipv6LAN:Record<string, any> = data?.WEB_GUI?.IPv6?.LAN || {};
    form.setFieldsValue(ipv6LAN);
  }, [data]);

  return (
    <ProForm
      {...layout}
      form={form}
      layout="horizontal"
      onFinish={onFinish}
      onReset={onReset}
      labelWrap
      submitter={{
        render: (props, doms) => {
          return (
            <Row>
              <Col span={24} offset={8}>
                <Space>{doms}</Space>
              </Col>
            </Row>
          );
        },
      }}
    >
      <ProCard split="horizontal">
        <ProCard title="LAN Settings">
          <Form.Item label="LAN Connection Type" name="Type" >
            <Select placeholder="Select LAN Connection Type" >
              {typeOptions.map((item) => (
                <Select.Option key={item.value} value={item.value} >
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item label="IPv6 Link-Local Address" name="LocalAddress">
            <Input readOnly />
          </Form.Item>
        </ProCard>
      </ProCard>
    </ProForm>
  );
};

export default IPv6LANSettingsForm;
