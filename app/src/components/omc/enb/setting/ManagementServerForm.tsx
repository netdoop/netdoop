import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface ManagementServerFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const ManagementServerForm: React.FC<ManagementServerFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [form2] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.ManagementServer.'
    // values['Device.ManagementServer.URL'] = data2.URL
    values[prefixName + 'PeriodicInformInterval'] = data2.PeriodicInformInterval

    setDeviceParameterValues(device, values)
  };

  const onFinish2 = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.ManagementServer.'
    values[prefixName + 'STUNEnable'] = data2.STUNEnable;
    values[prefixName + 'STUNServerAddress'] = data2.STUNServerAddress;
    values[prefixName + 'STUNServerPort'] = data2.STUNServerPort;
    values[prefixName + 'STUNUsername'] = data2.STUNUsername;
    values[prefixName + 'STUNPassword'] = data2.STUNPassword;
    values[prefixName + 'STUNMaximumKeepAlivePeriod'] = data2.STUNMaximumKeepAlivePeriod;
    values[prefixName + 'STUNMinimumKeepAlivePeriod'] = data2.STUNMinimumKeepAlivePeriod;
    setDeviceParameterValues(device, values)
  };

  const onReset = () => {
    const managementServer: Record<string, any> = data?.ManagementServer || {}
    form.setFieldsValue(managementServer);
    form2.setFieldsValue(managementServer);
  };
  useEffect(() => {
    const managementServer: Record<string, any> = data?.ManagementServer || {}
    form.setFieldsValue(managementServer);
    form2.setFieldsValue(managementServer);
  }, [data]);

  return (
    <ProCard split="horizontal">
      <ProCard>
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
          {/* <Form.Item name="EnableCWMP" label="Enable TR069">
          <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item> */}
          <Form.Item name="URL" label="URL">
            <Input disabled />
          </Form.Item>
          <Form.Item name="Username" label="Username">
            <Input disabled />
          </Form.Item>
          <Form.Item name="Password" label="Password">
            <Input.Password disabled />
          </Form.Item>

          <Form.Item name="PeriodicInformEnable" label="Periodic Inform Enable" >
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="PeriodicInformInterval" label="Periodic Inform Interval">
            <InputNumber placeholder="Enter random time in minutes" addonAfter="sec" />
          </Form.Item>
          {/* <Form.Item name="Periodic Inform Time" label="Periodic Inform Time">
        <Input disabled/>
      </Form.Item> */}

          <Form.Item name="ConnectionRequestURL" label="Connection Request URL">
            <Input disabled />
          </Form.Item>
          <Form.Item name="ConnectionRequestUsername" label="Connection Request Username">
            <Input disabled />
          </Form.Item>
          <Form.Item name="ConnectionRequestPassword" label="Connection Request Password">
            <Input.Password disabled />
          </Form.Item>
        </ProForm>
      </ProCard>

      <ProCard title="STUN">
        <ProForm
          {...layout}
          form={form2}
          layout="horizontal"
          onFinish={onFinish2}
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
          <Form.Item name="STUNEnable" label="STUN Enable" >
            <Select>
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="STUNServerAddress" label="STUN Server Address">
            <Input />
          </Form.Item>
          <Form.Item name="STUNServerPort" label="STUN Server Port">
            <InputNumber />
          </Form.Item>
          <Form.Item name="STUNUsername" label="STUN Username">
            <Input />
          </Form.Item>
          <Form.Item name="STUNPassword" label="STUN Password">
            <Input.Password />
          </Form.Item>
          <Form.Item name="STUNMaximumKeepAlivePeriod" label="STUN Maximum Keep Alive Period">
            <InputNumber />
          </Form.Item>
          <Form.Item name="STUNMinimumKeepAlivePeriod" label="STUN Minimum Keep Alive Period">
            <InputNumber />
          </Form.Item>
          <Form.Item name="NATDetected" label="NAT Detected">
            <Input disabled readOnly />
          </Form.Item>
          <Form.Item name="UDPConnectionRequestAddress" label="UDP Connection Request Address">
            <Input disabled readOnly />
          </Form.Item>
        </ProForm>
      </ProCard>
    </ProCard>

  );
};

export default ManagementServerForm;
