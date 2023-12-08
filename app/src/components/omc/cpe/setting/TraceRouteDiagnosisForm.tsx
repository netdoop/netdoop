import { Col, Descriptions, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";

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

interface routeHopsItem {
  Key: string,
  HopHost?: string,
  HopHostAddress?: string,
  HopErrorCode?: number,
  HopRTTimes?: string,
};

interface TraceRouteDiagnosisFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const TraceRouteDiagnosisForm: React.FC<TraceRouteDiagnosisFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [routeHopsItems, setRouteHopsItems] = useState<routeHopsItem[]>([]);

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.TraceRouteDiagnostics.';
    values[prefixName + 'DiagnosticsState'] = data2.DiagnosticsState;
    values[prefixName + 'Interface'] = data2.Interface;

    values[prefixName + 'Host'] = data2.Host;
    values[prefixName + 'NumberOfRepetitions'] = data2.NumberOfTries;
    values[prefixName + 'Timeout'] = data2.Timeout;
    values[prefixName + 'DataBlockSize'] = data2.DataBlockSize;
    values[prefixName + 'MaxHopCount'] = data2.MaxHopCount;

    setDeviceParameterValues(device, values)
  };

  const onReset = () => {
    const traceRouteDiagnostics: Record<string, any> = data?.TraceRouteDiagnostics || {};
    form.setFieldsValue(traceRouteDiagnostics);
  };
  useEffect(() => {
    const traceRouteDiagnostics: Record<string, any> = data?.TraceRouteDiagnostics || {};
    const routeHops: Record<string, Record<string, any>> = data?.TraceRouteDiagnostics?.RouteHops || {};
    let items: routeHopsItem[] = [];
    Object.entries(routeHops).forEach(([k, v]) => {
      items.push({
        Key: k,
        HopHost: v.HopHost,
        HopHostAddress: v.HopHostAddress,
        HopErrorCode: v.HopErrorCode,
        HopRTTimes: v.HopRTTimes,
      })
    });
    setRouteHopsItems(items);
    form.setFieldsValue(traceRouteDiagnostics);
  }, [data]);

  const routeHopsColumns = [
    { title: "#", dataIndex: "Index", key: "Index" },
    { title: "HopHost", dataIndex: "HopHost", key: "HopHost" },
    { title: "HopHostAddress", dataIndex: "HopHostAddress", key: "HopHostAddress" },
    { title: "HopErrorCode", dataIndex: "HopErrorCode", key: "HopErrorCode" },
    { title: "HopRTTimes", dataIndex: "HopRTTimes", key: "HopRTTimes" },
  ]
  return (
    <ProCard split="horizontal">
      <ProCard title="TraceRoute">
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
          <Form.Item label="Interface" name="Interface">
            <Input placeholder="Enter Interface" />
          </Form.Item>
          <Form.Item label="Target IP/Domain" name="Host">
            <Input placeholder="Enter Target IP/Domain" />
          </Form.Item>
          <Form.Item label="Maximum Hops" name="MaxHopCount" >
            <InputNumber placeholder="Enter Maximum Hops" step={1} min={1} max={64} />
          </Form.Item>
          <Form.Item label="Timeout" name="Timeout" >
            <InputNumber placeholder="Enter Timeout" step={1} min={1} />
          </Form.Item>
          <Form.Item label="Packet Size" name="DataBlockSize">
            <InputNumber placeholder="Enter Packet Size" step={1} min={1} max={65535} />
          </Form.Item>

          <Form.Item label="Count" name="NumberOfTries" >
            <InputNumber placeholder="Enter Count" step={1} min={1} max={3} />
          </Form.Item>
          <Form.Item label="DSCP" name="DSCP" >
            <InputNumber placeholder="Enter DSCP" step={1} min={0} max={63} />
          </Form.Item>
        </ProForm >
      </ProCard>
      <ProCard title="Status">
        <Descriptions title="" column={1}>
          <Descriptions.Item label="State">{data?.TraceRouteDiagnostics?.DiagnosticsState || '-'}</Descriptions.Item>
          <Descriptions.Item label="Response Time">{data?.TraceRouteDiagnostics?.ResponseTime || '-'}</Descriptions.Item>
        </Descriptions>
      </ProCard>
      <ProCard title="Route Hops">
        <ProTable
          dataSource={routeHopsItems}
          columns={routeHopsColumns}
          pagination={false}
          search={false}
          options={false}
        />;
      </ProCard>
    </ProCard>
  );
};

export default TraceRouteDiagnosisForm;
