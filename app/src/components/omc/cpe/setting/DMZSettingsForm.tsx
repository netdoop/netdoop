import { Col, Form, Input, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface DMZSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const DMZSettingsForm: React.FC<DMZSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  
  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WebGui.Network.DMZ.';
    values[prefixName + 'Enable'] = data2.Enable
    values[prefixName + 'HostAddress'] = data2.HostAddress
    values[prefixName + 'ICMPRedirectEnable'] = data2.ICMPRedirectEnable
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const dmz:Record<string, any> = data?.WEB_GUI?.Network?.DMZ || {};
    form.setFieldsValue(dmz);
  };
  useEffect(() => {
    const dmz:Record<string, any> = data?.WEB_GUI?.Network?.DMZ || {};
    form.setFieldsValue(dmz);
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
      <Form.Item name="Enable" label="DMZ" >
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item name="ICMPRedirectEnable" label="ICMP Redirect">
         <Select >
            <Select.Option key="0" value="0">Disable</Select.Option>
            <Select.Option key="1" value="1">Enable</Select.Option>
          </Select>
      </Form.Item>
      <Form.Item name="HostAddress" label="Host Address">
        <Input />
      </Form.Item>
    </ProForm>
  );
};

export default DMZSettingsForm;
