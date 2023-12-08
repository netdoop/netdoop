import { Col, Form, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const networkModeOptions = [
  { label: 'NAT', value: 'nat' },
  { label: 'Bridge', value: 'bridge' },
  { label: 'ROUTER', value: 'router' },
];

interface WANSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const WANSettingsForm: React.FC<WANSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    values['Device.WEB_GUI.Network.WANSettings.NetworkMode'] = data2.NetworkMode
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const wanSettings:Record<string, any> = data?.WEB_GUI?.Network?.WANSettings || {};
    form.setFieldsValue(wanSettings);
  };
  useEffect(() => {
    const wanSettings:Record<string, any> = data?.WEB_GUI?.Network?.WANSettings || {};
    form.setFieldsValue(wanSettings);
  }, [data]);
  
  return (
    <ProCard split="horizontal">
      <ProCard title="Settings">
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
          <Form.Item name="NetworkMode" label="Network Mode">
            <Select placeholder="Select network mode">
              {networkModeOptions.map((option) => (
                <Select.Option key={option.value} value={option.value}>
                  {option.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
        </ProForm>
      </ProCard>
    </ProCard>
  );
};

export default WANSettingsForm;
