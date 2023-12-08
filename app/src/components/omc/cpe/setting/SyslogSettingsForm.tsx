import { Checkbox, Col, Form, Input, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const methodOptions: { label: string, value: string }[] = [
  { label: "Local", value: "local" },
  { label: "Network", value: "network" },
]
const levelOptions: { label: string, value: string }[] = [
  { label: 'Emergency', value: '1' },
  { label: 'Alert', value: '2' },
  { label: 'Critical', value: '3' },
  { label: 'Error', value: '4' },
  { label: 'Warning', value: '5' },
  { label: 'Notice', value: '6' },
  { label: 'Info', value: '7' },
  { label: 'Debug', value: '8' }
];

interface SyslogSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const SyslogSettingsForm: React.FC<SyslogSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [method, setMethod] = useState<string>("local");

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WebGui.System.Syslog.';
    values[prefixName + 'Method'] = data2.Method
    if (data2.Method === "network") {
      values[prefixName + 'ForwardIPAddress'] = data2.ForwardIPAddress
    }
    values[prefixName + 'Level'] = data2.Level
    values[prefixName + 'Clear'] = data2.Clear
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const syslog:Record<string, any> = data?.WEB_GUI?.System?.Syslog || {};
    form.setFieldsValue(syslog);
  };
  useEffect(() => {
    const syslog:Record<string, any> = data?.WEB_GUI?.System?.Syslog || {};
    form.setFieldsValue(syslog);
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
      <Form.Item label="Method" name="Method" >
        <Select placeholder="Select Method" onChange={(v) => { setMethod(v) }}>
          {methodOptions.map((item) => (
            <Select.Option key={item.value} value={item.value}>
              {item.label}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item label="Level" name="Level" >
        <Select placeholder="Select Level">
          {levelOptions.map((item) => (
            <Select.Option key={item.value} value={item.value}>
              {item.label}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
      {method === "network" && (
        <Form.Item label="ForwardIPAddress" name="ForwardIPAddress">
          <Input />
        </Form.Item>
      )}
      <Form.Item label="Clear" name="Clear">
        <Checkbox checked={false} />
      </Form.Item>
    </ProForm >
  );
};

export default SyslogSettingsForm;
