import { Col, Form, InputNumber, Row, Space } from "antd";
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

const SCTPSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.Transport.SCTP.';
    values[prefixName + 'HBInterval'] = data2.HBInterval;
    values[prefixName + 'MaxPathRetransmits'] = data2.MaxPathRetransmits;
    values[prefixName + 'MaxAssociationRetransmits'] = data2.MaxAssociationRetransmits;
    values[prefixName + 'RTOInitial'] = data2.RTOInitial;
    values[prefixName + 'RTOMax'] = data2.RTOMax;
    values[prefixName + 'RTOMin'] = data2.RTOMin;
    values[prefixName + 'ValCookieLife'] = data2.ValCookieLife;
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const sctp:Record<string, any> = data?.Services?.FAPService?.["1"]?.Transport?.SCTP;
    form.setFieldsValue(sctp);
  };
  useEffect(() => {
    const sctp:Record<string, any> = data?.Services?.FAPService?.["1"]?.Transport?.SCTP;
    form.setFieldsValue(sctp);
  }, [data]);

  return (
    <ProCard split="horizontal">
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
        <Form.Item name="HBInterval" label="HBInterval">
          <InputNumber />
        </Form.Item>
        <Form.Item name="MaxPathRetransmits" label="MaxPathRetransmits">
          <InputNumber />
        </Form.Item>
        <Form.Item name="MaxAssociationRetransmits" label="MaxAssociationRetransmits">
          <InputNumber />
        </Form.Item>
        <Form.Item name="RTOInitial" label="RTOInitial">
          <InputNumber />
        </Form.Item>
        <Form.Item name="RTOMax" label="RTOMax">
          <InputNumber />
        </Form.Item>
        <Form.Item name="RTOMin" label="RTOMin">
          <InputNumber />
        </Form.Item>
        <Form.Item name="ValCookieLife" label="ValCookieLife">
          <InputNumber />
        </Form.Item>
      </ProForm>

    </ProCard>

  );
};

export default SCTPSettingForm;
