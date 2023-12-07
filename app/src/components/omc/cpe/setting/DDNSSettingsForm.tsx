import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const serviceProviderOptions: { label: string, value: string }[] = [
  { label: 'WWW.DYNDNS.ORG', value: 'WWW.DYNDNS.ORG' },
  // { label: 'www.dyndns.org (custom)', value: 'WWW.DYNDNS.ORG(CUSTOM)' },
  // { label: 'www.dyndns.org (static)', value: 'WWW.DYNDNS.ORG(STATIC)' },
  { label: 'WWW.TZO.COM', value: 'WWW.TZO.COM' },
  { label: 'WWW.ZONEEDIT.COM', value: 'WWW.ZONEEDIT.COM' },
  { label: 'WWW.JUSTLINUX.COM', value: 'WWW.JUSTLINUX.COM' },
  { label: 'WWW.EASYDNS.COM', value: 'WWW.EASYDNS.COM' },
  { label: 'WWW.DNSOMATIC.COM', value: 'WWW.DNSOMATIC.COM' },
  { label: 'WWW.TUNNELBROKER.NET', value: 'WWW.TUNNELBROKER.NET' },
  { label: 'WWW.NO-IP.COM', value: 'WWW.NO-IP.COM' },
  { label: 'WWW.NAMECHEAP.COM', value: 'WWW.NAMECHEAP.COM' },
  { label: 'WWW.SELFHOST.DE', value: 'WWW.SELFHOST.DE' },
  { label: 'DOMAINS.GOOGLE.COM', value: 'DOMAINS.GOOGLE.COM' },
  { label: 'WWW.ORAY.COM', value: 'WWW.ORAY.COM' },
];


interface DDNSSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const DDNSSettingsForm: React.FC<DDNSSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WebGui.System.DDNS.';
    values[prefixName + 'Enable'] = data2.Enable
    values[prefixName + 'ServiceProvider'] = data2.ServiceProvider
    values[prefixName + 'Domain'] = data2.Domain
    values[prefixName + 'Username'] = data2.Username
    values[prefixName + 'Password'] = data2.Password
    values[prefixName + 'Refresh'] = data2.Refresh
    values[prefixName + 'CheckEvery'] = data2.CheckEvery
    values[prefixName + 'Wildcard'] = data2.Wildcard
    values[prefixName + 'Verification'] = data2.Verification

    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const ddns: Record<string, any> = data?.WEB_GUI?.System?.DDNS || {};
    form.setFieldsValue(ddns);
  };
  useEffect(() => {
    const ddns: Record<string, any> = data?.WEB_GUI?.System?.DDNS || {};
    form.setFieldsValue(ddns);
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
      <Form.Item label="DDNS" name="Enable">
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item label="Service Provider" name="ServiceProvider" >
        <Select placeholder="Select Time Zone">
          {serviceProviderOptions.map((item) => (
            <Select.Option key={item.value} value={item.value}>
              {item.label}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item label="Domain" name="Domain">
        <Input />
      </Form.Item>
      <Form.Item label="Username" name="Username">
        <Input />
      </Form.Item>
      <Form.Item label="Password" name="Password">
        <Input.Password />
      </Form.Item>
      <Form.Item label="Refresh" name="Refresh" >
        <InputNumber placeholder="Enter Refresh" step={1} min={1} max={30} />
      </Form.Item>
      <Form.Item label="Enable Wildcard" name="Wildcard">
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item label="WAN IP and domain verification" name="Verification">
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item label="Check Every" name="CheckEvery">
        <InputNumber placeholder="Enter Check Every" step={1} min={30} max={1440} />
      </Form.Item>
      <Form.Item label="Status" name="Status">
        <Input readOnly />
      </Form.Item>
    </ProForm >
  );
};

export default DDNSSettingsForm;
