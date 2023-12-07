import { Col, Form, Row, Input, Space, Select } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const typeOptions = [
  { label: 'AutoConfiguration', value: 'AutoConfiguration' },
];

const dnsFormOptions = [
  { label: 'DHCPv6', value: 'DHCPv6' },
];

interface IPv6WANSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const IPv6WANSettingsForm: React.FC<IPv6WANSettingsFormProps> = ({
  device,
  data,
}) => {
  const prefixName = 'Device.WEB_GUI.IPv6.WAN.';
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {};
    values[prefixName + 'Enable'] = data2.Enable;
    values[prefixName + 'DNSForm'] = data2.DNSForm;
    values[prefixName + 'Type'] = data2.Type;
    setDeviceParameterValues(device, values);
  };

  const onReset = () => {
    const ipv6WAN:Record<string, any> = data?.WEB_GUI?.IPv6?.WAN || {};
    form.setFieldsValue(ipv6WAN);
  };
  useEffect(() => {
    const ipv6WAN:Record<string, any> = data?.WEB_GUI?.IPv6?.WAN || {};
    form.setFieldsValue(ipv6WAN);
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
          <Form.Item label="IPv6 Enable" name="Enable">
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label="WAN Connection Type" name="Type" >
            <Select placeholder="Select WAN Connection Type" >
              {typeOptions.map((item) => (
                <Select.Option key={item.value} value={item.value} >
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item label="IPv6 MGMT Global Address" name="GlobalAddress">
            <Input readOnly />
          </Form.Item>
          <Form.Item label="DNS Form" name="DNSForm">
            <Select placeholder="Select DNS Form" >
              {dnsFormOptions.map((item) => (
                <Select.Option key={item.value} value={item.value} >
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>

          {/* <Form.Item label="IPv6 Pre" name="IPv6Pre">
            <Input />
          </Form.Item>
          <Form.Item label="Gateway" name="Gateway">
            <Input />
          </Form.Item>

          <Form.Item label="DNS1" name="DNS1" >
            <Input />
          </Form.Item>

          <Form.Item label="DNS2" name="DNS2">
            <Input />
          </Form.Item>
          <Form.Item label="PD Enable" name="PDEnable" >
         <Select >
            <Select.Option key="0" value="0">Disable</Select.Option>
            <Select.Option key="1" value="1">Enable</Select.Option>
          </Select>
          </Form.Item> */}
        </ProCard>
      </ProCard>
    </ProForm>

  );
};

export default IPv6WANSettingsForm;
