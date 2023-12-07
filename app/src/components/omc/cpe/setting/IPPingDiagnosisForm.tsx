import { Col, Descriptions, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const diagnosticsStateOptions: { label: string, value: string }[] = [
  { label: 'None', value: 'NONE' },
  { label: 'Requested', value: 'REQUESTED' },
  { label: 'Complete', value: 'COMPLETE' },
  { label: 'Error_CannotResolveHostName', value: 'ERROR_CANNOT_RESOLVE_HOSTNAME' },
  { label: 'Error_Internal', value: 'ERROR_INTERNAL' },
  { label: 'Error_Other', value: 'ERROR_OTHER' }
];

interface IPPingDiagnosisFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const IPPingDiagnosisForm: React.FC<IPPingDiagnosisFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.IPPingDiagnostics.';
    values[prefixName + 'DiagnosticsState'] = data2.DiagnosticsState;
    values[prefixName + 'Host'] = data2.Host;
    values[prefixName + 'NumberOfRepetitions'] = data2.NumberOfRepetitions;
    values[prefixName + 'Timeout'] = data2.Timeout;
    values[prefixName + 'DataBlockSize'] = data2.DataBlockSize;
    values[prefixName + 'DSCP'] = data2.DSCP;

    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const ipPingDiagnostics:Record<string, any> = data?.IPPingDiagnostics || {};
    form.setFieldsValue(ipPingDiagnostics);
  };
  useEffect(() => {
    const ipPingDiagnostics:Record<string, any> = data?.IPPingDiagnostics || {};
    form.setFieldsValue(ipPingDiagnostics);
  }, [data]);
  return (
    <ProCard split="horizontal">
      <ProCard title="Ping">
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
          <Form.Item label="Diagnostics State" name="DiagnosticsState" >
            <Select placeholder="Select Diagnostics State">
              {diagnosticsStateOptions.map((item) => (
                <Select.Option key={item.value} value={item.value}>
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item label="Target IP/Domain" name="Host">
            <Input placeholder="Enter Target IP/Domain" />
          </Form.Item>
          <Form.Item label="Packet Size" name="DataBlockSize">
            <InputNumber placeholder="Enter Packet Size" step={1} min={1} max={65535} />
          </Form.Item>
          <Form.Item label="Timeout" name="Timeout" >
            <InputNumber placeholder="Enter Timeout" step={1} min={1} />
          </Form.Item>
          <Form.Item label="Count" name="NumberOfRepetitions" >
            <InputNumber placeholder="Enter Count" step={1} min={1} />
          </Form.Item>
          <Form.Item label="DSCP" name="DSCP" >
            <InputNumber placeholder="Enter DSCP" step={1} min={0} max={63} />
          </Form.Item>
        </ProForm >
      </ProCard>
      <ProCard title="Result">
        <Descriptions title="" column={1}>
          <Descriptions.Item label="State">{data?.IPPingDiagnostics?.DiagnosticsState || '-'}</Descriptions.Item>
          <Descriptions.Item label="Success Count">{data?.IPPingDiagnostics?.SuccessCount || '-'}</Descriptions.Item>
          <Descriptions.Item label="Failure Count">{data?.IPPingDiagnostics?.FailureCount || '-'}</Descriptions.Item>
          <Descriptions.Item label="Average Response Time">{data?.IPPingDiagnostics?.AverageResponseTime || '-'}</Descriptions.Item>
          <Descriptions.Item label="Maximum Response Time">{data?.IPPingDiagnostics?.MaximumResponseTime || '-'}</Descriptions.Item>
          <Descriptions.Item label="Minimum Response Time">{data?.IPPingDiagnostics?.MinimumResponseTime || '-'}</Descriptions.Item>
        </Descriptions>
      </ProCard>
    </ProCard>
  );
};

export default IPPingDiagnosisForm;
