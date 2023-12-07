import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProForm } from '@ant-design/pro-components';
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

const NRMMgmtSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.FAP.NRMMgmt.'

    values[prefixName + 'Enable'] = data2.Enable;
    values[prefixName + 'URL'] = data2.URL;
    values[prefixName + 'PeriodicUploadInterval'] = data2.PeriodicUploadInterval;
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const nrmMgmt:Record<string, any> = data?.FAP?.NRMMgmt;
    form.setFieldsValue(nrmMgmt);
  };
  useEffect(() => {
    const nrmMgmt:Record<string, any> = data?.FAP?.NRMMgmt;
    form.setFieldsValue(nrmMgmt);
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
      <Form.Item name="Enable" label="Upload Enable">
        <Select >
          <Select.Option key="0" value="0">Disable</Select.Option>
          <Select.Option key="1" value="1">Enable</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item name="URL" label="URL">
        <Input />
      </Form.Item>
      <Form.Item name="PeriodicUploadInterval" label="Periodic Upload Interval">
        <InputNumber addonAfter="sec" />
      </Form.Item>
    </ProForm>
  );
};

export default NRMMgmtSettingForm;
