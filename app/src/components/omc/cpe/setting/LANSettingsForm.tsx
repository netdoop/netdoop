import { Col, Form, Row, Input, Space, InputNumber, Select } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface LANSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const LANSettingsForm: React.FC<LANSettingsFormProps> = ({
  device,
  data,
}) => {
  const prefixName = 'Device.WEB_GUI.Network.LANSettings.';
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {};
    values[prefixName + 'LANHost.IPAddress'] = data2["LANHost.IPAddress"];
    values[prefixName + 'LANHost.SubnetMask'] = data2["LANHost.SubnetMask"];
    values[prefixName + 'DHCP.ServerEnable'] = data2["DHCP.ServerEnable"];
    values[prefixName + 'DHCP.StartIP'] = data2["DHCP.StartIP"];
    values[prefixName + 'DHCP.EndIP'] = data2["DHCP.EndIP"];
    values[prefixName + 'DHCP.LeaseTime'] = data2["DHCP.LeaseTime"];

    setDeviceParameterValues(device, values);
  };

  const onReset = () => {
    const lanSettings:Record<string, any> = data?.WEB_GUI?.Network?.LANSettings || {};
    form.setFieldsValue(lanSettings);
  };
  useEffect(() => {
    const lanSettings:Record<string, any> = data?.WEB_GUI?.Network?.LANSettings || {};
    form.setFieldsValue(lanSettings);
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
          <Form.Item key="IPAddress" name={["LANHost","IPAddress"]} label="IP Address">
            <Input placeholder="Enter IP Address" />
          </Form.Item>
          <Form.Item key="SubnetMask" name={["LANHost","SubnetMask"]} label="Subnet Mask">
            <Input placeholder="Enter Subnet Mask" />
          </Form.Item>
        </ProCard>
        <ProCard title="DHCP Settings">
          <Form.Item key="ServerEnable" name={["DHCP","ServerEnable"]} label="DHCP Server">
            <Select>
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item key="StartIP" name={["DHCP","StartIP"]} label="Start IP Address">
            <Input placeholder="Enter Start IP Address" />
          </Form.Item>
          <Form.Item key="EndIP" name={["DHCP","EndIP"]} label="End IP Address">
            <Input placeholder="Enter End IP Address" />
          </Form.Item>
          <Form.Item key="LeaseTime" name={["DHCP","LeaseTime"]} label="Lease Time">
            <InputNumber placeholder="Enter Lease Time" step={30} min={30} />
          </Form.Item>
        </ProCard>
      </ProCard>
    </ProForm>

  );
};

export default LANSettingsForm;
