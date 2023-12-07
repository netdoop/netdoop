import { Col, Form, Input, Row, Select, Space } from "antd";
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

const RFSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [form2] = Form.useForm();

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.';
    values[prefixName + 'CellConfig.LTE.RAN.RF.DLBandwidth'] = data2["CellConfig.LTE.RAN.RF.DLBandwidth"];
    values[prefixName + 'CellConfig.LTE.RAN.RF.ULBandwidth'] = data2["CellConfig.LTE.RAN.RF.ULBandwidth"];
    values[prefixName + 'FAPControl.LTE.RFTxStatus'] = data2["FAPControl.LTE.RFTxStatus"];
    setDeviceParameterValues(device, values)
  };

  const onFinish2 = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.';
    values[prefixName + 'CellConfig.LTE.RAN.RF.Ante1WithPssConfig'] = data2["CellConfig.LTE.RAN.RF.Ante1WithPssConfig"];
    setDeviceParameterValues(device, values)
  };

  const onFinish3 = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.';
    values[prefixName + 'FAPControl.LTE.SelfConfig.SONConfigParam.PCIOptEnable'] = data2["FAPControl.LTE.SelfConfig.SONConfigParam.PCIOptEnable"];
    values[prefixName + 'FAPControl.LTE.SelfConfig.SONConfigParam.CandidatePCIList'] = data2["FAPControl.LTE.SelfConfig.SONConfigParam.CandidatePCIList"];
    setDeviceParameterValues(device, values)
  };

  const onFinish4 = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.';
    values[prefixName + 'CellConfig.LTE.RAN.RF.EARFCNDL'] = data2["CellConfig.LTE.RAN.RF.EARFCNDL"];
    // values[prefixName + 'CellConfig.LTE.RAN.RF.EARFCNUL'] = data2.CellConfig?.LTE?.RAN?.RF?.EARFCNUL;
    // values[prefixName + 'CellConfig.LTE.RAN.RF.FreqBandIndicator'] = data2.CellConfig?.LTE?.RAN?.RF?.FreqBandIndicator;    
    setDeviceParameterValues(device, values)
  };

  const onReset = () => {
    const fapService:Record<string, any> = data?.Services?.FAPService?.["1"] || {};
    form.setFieldsValue(fapService);
    form2.setFieldsValue(fapService);
  };
  useEffect(() => {
    const fapService:Record<string, any> = data?.Services?.FAPService?.["1"] || {};
    form.setFieldsValue(fapService);
    form2.setFieldsValue(fapService);
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
          <Form.Item name={["CellConfig.LTE.RAN.RF.DLBandwidth"]} label="DL Bandwidth">
            <Select>
              <Select.Option key="n5" value="n5">5MHz</Select.Option>
              <Select.Option key="n10" value="n10">10MHz</Select.Option>
              <Select.Option key="n15" value="n15">15MHz</Select.Option>
              <Select.Option key="n20" value="n20">20MHz</Select.Option>
              <Select.Option key="n25" value="n25">25MHz</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name={["CellConfig.LTE.RAN.RF.ULBandwidth"]} label="UL Bandwidth">
            <Select>
              <Select.Option key="n5" value="n5">5MHz</Select.Option>
              <Select.Option key="n10" value="n10">10MHz</Select.Option>
              <Select.Option key="n15" value="n15">15MHz</Select.Option>
              <Select.Option key="n20" value="n20">20MHz</Select.Option>
              <Select.Option key="n25" value="n25">25MHz</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name={["FAPControl.LTE.RFTxStatus"]} label="RF Switch">
            <Select >
              <Select.Option key="0" value="0">OFF</Select.Option>
              <Select.Option key="1" value="1">ON</Select.Option>
            </Select>
          </Form.Item>
        </ProForm>
      </ProCard>

      <ProCard title="Ante1WithPssConfig">
        <ProForm
          {...layout}
          form={form2}
          layout="horizontal"
          onFinish={onFinish2}
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
          <Form.Item name={["CellConfig.LTE.RAN.RF.Ante1WithPssConfig"]} label="Ante1WithPssConfig">
            <Select >
              <Select.Option key="0" value="0">Close</Select.Option>
              <Select.Option key="1" value="1">Open</Select.Option>
            </Select>
          </Form.Item>
        </ProForm>
      </ProCard>

      <ProCard title="PCI">
        <ProForm
          {...layout}
          form={form}
          layout="horizontal"
          onFinish={onFinish3}
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
          <Form.Item name={["FAPControl.LTE.SelfConfig.SONConfigParam.PCIOptEnable"]} label="PCI Switch">
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name={["FAPControl.LTE.SelfConfig.SONConfigParam.CandidatePCIList"]} label="PCI List">
            <Input />
          </Form.Item>
          <Form.Item name={["CellConfig.LTE.RAN.RF.PhyCellID"]} label="PCI">
            <Input readOnly disabled />
          </Form.Item>
        </ProForm>
      </ProCard>

      <ProCard title="EARFCN Config">
        <ProForm
          {...layout}
          form={form}
          layout="horizontal"
          onFinish={onFinish4}
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
          <Form.Item name={["CellConfig.LTE.RAN.RF.EARFCNDL"]} label="EARFCN_DL">
            <Input />
          </Form.Item>
          <Form.Item name={["CellConfig.LTE.RAN.RF.FreqBandIndicator"]} label="Band">
            <Input readOnly disabled />
          </Form.Item>
        </ProForm>
      </ProCard>
    </ProCard>
  );
};

export default RFSettingForm;
