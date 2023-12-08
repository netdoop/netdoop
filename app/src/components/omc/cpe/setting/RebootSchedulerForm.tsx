import { Col, Form, Input, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface RebootSchedulerProps {
  device: API.Device,
  data: Record<string, any>,
};

const RebootSchedulerForm: React.FC<RebootSchedulerProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WebGui.System.ScheduleReboot.';
    values[prefixName + 'Enable'] = data2.Enable
    values[prefixName + 'Time'] = data2.Time
    values[prefixName + 'DateToReboot'] = data2.DateToReboot
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const scheduleReboot:Record<string, any> = data?.WEB_GUI?.System?.ScheduleReboot || {};
    form.setFieldsValue(scheduleReboot);
  };
  useEffect(() => {
    const scheduleReboot:Record<string, any> = data?.WEB_GUI?.System?.ScheduleReboot || {};
    form.setFieldsValue(scheduleReboot);
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
      <Form.Item label="Enable Reboot Scheduler" name="Enable" >
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item label="Date To Reboot" name="DateToReboot">
        <Input />
      </Form.Item>
      <Form.Item label="Time of Day to Reboot" name="Time">
        <Input />
      </Form.Item>
    </ProForm >
  );
};

export default RebootSchedulerForm;
