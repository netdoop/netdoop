import { Col, Form, Input, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
import { useEffect, useState } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface PINManagementFormProps {
  device: API.Device,
  data?: Record<string, any>,
};

const PINManagementForm: React.FC<PINManagementFormProps> = ({ data }) => {
  const [form] = Form.useForm();
  const [verifyPIN, setVerifyPIN] = useState(0);
  const [changePIN, setChangePIN] = useState(0);

  const onFinish = async (values: Record<string, any>) => {
    console.log('Form values:', values);
    // Submit the form data to the server
  };
  const onReset = () => {
    const pinManagement:Record<string, any> = data?.WEB_GUI?.Network?.PINManagement || {};
    form.setFieldsValue(pinManagement);
  };
  useEffect(() => {
    const pinManagement:Record<string, any> = data?.WEB_GUI?.Network?.PINManagement || {};
    form.setFieldsValue(pinManagement);
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
      <Form.Item name="PINStatus" label="PIN Status">
        <span className="ant-form-text">{data?.PINStatus || '-'}</span>
      </Form.Item>

      <Form.Item name="VerifyEnable" label="PIN Verifification" >
        <Select onChange={setVerifyPIN}>
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>

      <Form.Item label="PIN" name="PIN">
        <Input.Password />
      </Form.Item>

      <Form.Item name="PINRemainingAttempts" label="Remaining Attempts">
        <span className="ant-form-text">{data?.PINRemainingAttempts || '-'}</span>
      </Form.Item>

      <Form.Item name="ChangePIN" label="Change PIN" hidden={verifyPIN === 0}>
        <Select onChange={setChangePIN}>
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>

      <Form.Item label="NewPIN" name="New PIN" hidden={changePIN === 0}>
        <Input.Password />
      </Form.Item>

      <Form.Item label="ConfirmPIN" name="Confirm PIN" hidden={changePIN === 0}>
        <Input.Password />
      </Form.Item>


    </ProForm>
  );
};

export default PINManagementForm;
