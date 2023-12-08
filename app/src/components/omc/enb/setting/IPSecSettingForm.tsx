import { Col, Form, Input, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
interface Props {
  device: API.Device,
  data: Record<string, any>,
};

const IPSecSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.IPsec.'
    values[prefixName + 'Enable'] = data2.Enable
    values[prefixName + 'MyKeyMode'] = data2.MyKeyMode
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const ipsec:Record<string, any> = data?.IPsec || {};
    form.setFieldsValue(ipsec);
  };
  useEffect(() => {
    const ipsec:Record<string, any> = data?.IPsec || {};
    form.setFieldsValue(ipsec);
  }, [data]);

  return (
    <ProForm
      {...layout}
      form={form}
      layout="horizontal"
      onFinish={onFinish}
      onReset={onReset}
      labelWrap
      disabled
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
        <Form.Item name="Enable" label="IPSec Enable">
          <Select >
            <Select.Option key="0" value="0">Disable</Select.Option>
            <Select.Option key="1" value="1">Enable</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item name="MyKeyMode" label="Authentication Type">
          <Input />
        </Form.Item>
      </ProCard>
    </ProForm>
  );
};

export default IPSecSettingForm;
