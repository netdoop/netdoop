import { Checkbox, Col, Form, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { useEffect, useState } from "react";
import { setDeviceParameterValues } from "@/models/device";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const scanModeOptions = [
  { label: 'Full Band', value: "0" },
  { label: 'Band Lock', value: "1" },
];

// const preferModeOptions = [
//   { label: 'Auto', value: '0' },
//   { label: '5G SA', value: '1' },
//   { label: '5G NSA', value: '1' },
//   { label: '4G Only', value: '1' },
//   { label: '3G Only', value: '1' },
// ];

type BandLockOption = {
  label: string;
  value: string;
}

interface ScanModeFormProps {
  device: API.Device,
  data: Record<string, any>,
}

const ScanModeForm: React.FC<ScanModeFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [bandLockOptions, setBandLockOptions] = useState<BandLockOption[]>([]);
  const [selectMode, setSelectMode] = useState<string | undefined>('0');

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    values['Device.WEB_GUI.Network.ScanMode.PreferMode'] = data2.PreferMode
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const scanMode: Record<string, any> = data?.WEB_GUI?.Network?.ScanMode || {};
    form.setFieldsValue(scanMode);
  };
  useEffect(() => {
    const scanMode: Record<string, any> = data?.WEB_GUI?.Network?.ScanMode || {};
    const updateOptions: BandLockOption[] = [];
    const supportBand: string[] = scanMode?.SuppBand?.split(" ") || [];
    supportBand.forEach((v) => {
      if (v.startsWith("50")) {
        updateOptions.push({ label: "Band N" + v.substring(2), value: v })
      }
    });
    setBandLockOptions(updateOptions);
    form.setFieldsValue(scanMode);
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
        <ProCard title="Settings">
          <Form.Item name="PreferMode" label="Scan Mode">
            <Select
              placeholder="Select scan mode"
              onSelect={(key: string) => { setSelectMode(key) }}
            >
              {scanModeOptions.map((option) => (
                <Select.Option key={option.value} value={option.value}>
                  {option.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
        </ProCard>
        <ProCard title="Band Lock" hidden={selectMode !== '1'}>
          <Form.Item name="LockBand" label="Lock Band">
            <Checkbox.Group options={bandLockOptions} />
          </Form.Item>
        </ProCard>
      </ProCard>
    </ProForm >

  );
};

export default ScanModeForm;
