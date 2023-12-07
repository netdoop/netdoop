import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface OnlineUpgradeFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const OnlineUpgradeForm: React.FC<OnlineUpgradeFormProps> = ({
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (values: Record<string, any>) => {
    console.log('Form values:', values);
    // Submit the form data to the server
  };

  const onReset = () => {
    const onlineUpgrade:Record<string, any> = data?.WEB_GUI?.System?.OnlineUpgrade || {};
    form.setFieldsValue(onlineUpgrade);
  };
  useEffect(() => {
    const onlineUpgrade:Record<string, any> = data?.WEB_GUI?.System?.OnlineUpgrade || {};
    form.setFieldsValue(onlineUpgrade);
  }, [data]);

  const validateCheckNewFirmwareEvery = (rule: any, value: string) => {
    const intValue = parseInt(value, 10);
    if (isNaN(intValue)) {
      return Promise.reject('Please enter a valid integer');
    } else if (intValue < 1 || intValue > 1440) {
      return Promise.reject('Please enter a value between 1 and 1440');
    } else {
      return Promise.resolve();
    }
  };

  const validateRandomTime = (rule: any, value: string) => {
    const intValue = parseInt(value, 10);
    if (isNaN(intValue)) {
      return Promise.reject('Please enter a valid integer');
    } else if (intValue < 0 || intValue > 1440) {
      return Promise.reject('Please enter a value between 0 and 1440');
    } else {
      return Promise.resolve();
    }
  };


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
      <Form.Item
        name="Enable"
        label="Enable Online Upgrade"
      >
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item
        name="VersionFile"
        label="Version File Path"
        rules={[
          {
            required: true,
            message: 'Please enter a version file path',
          },
        ]}
      >
        <Input placeholder="Enter version file path" />
      </Form.Item>
      <Form.Item
        name="UpgradeFolder"
        label="Upgrade Folder Path"
        rules={[
          {
            required: true,
            message: 'Please enter an upgrade folder path',
          },
        ]}
      >
        <Input placeholder="Enter upgrade folder path" />
      </Form.Item>
      <Form.Item name="Username" label="Username">
        <Input placeholder="Enter username" />
      </Form.Item>
      <Form.Item name="Password" label="Password">
        <Input.Password placeholder="Enter password" />
      </Form.Item>
      <Form.Item
        name="CheckNewFWAfterConnectedEnable"
        label="Check for New Firmware after Connection"
      >
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item
        name="CheckNewFirmwareEvery"
        label="Check for New Firmware Every"
        rules={[
          {
            validator: validateCheckNewFirmwareEvery,
          },
        ]}
      >
        <InputNumber placeholder="Enter time in minutes" addonAfter="min" />
      </Form.Item>
      <Form.Item name="StartTime" label="Start Time">
        <Input placeholder="Select start time" />
      </Form.Item>
      <Form.Item
        name="RandomTime"
        label="Random Time"
        rules={[
          {
            validator: validateRandomTime,
          },
        ]}
      >
        <InputNumber placeholder="Enter random time in minutes" addonAfter="min" />
      </Form.Item>
    </ProForm>
  );
};

export default OnlineUpgradeForm;
