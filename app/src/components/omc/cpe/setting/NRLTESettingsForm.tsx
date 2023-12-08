import { Col, Descriptions, Form, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const connectModeOptions = [
  { label: 'Auto', value: 'auto_select' },
  { label: 'Manual', value: 'manual_select' },
];

interface NRLTENetworkSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const NRLTENetworkSettingsForm: React.FC<NRLTENetworkSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [nrlte, setNRLTE] = useState<Record<string, any>>({})
  const [mode, setMode] = useState<string>("");

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    values['Device.WEB_GUI.Network.NR-LTE.ConnectMethod'] = data2.ConnectMethod
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    form.setFieldsValue(nrlte);
  };
  useEffect(() => {
    const nrlte2: Record<string, any> = data?.WEB_GUI?.Network?.['NR-LTE'] || {};
    const mode2: string = data?.WEB_GUI?.Overview?.InternetStatus?.Mode|| "";
    setNRLTE(nrlte2);
    setMode(mode2);
    form.setFieldsValue(nrlte2);
  }, [data]);

  const renderStatus = () => {
    if (mode?.startsWith("5G")) {
      return (
        <ProCard title="Status">
          <ProCard split="vertical">
            <ProCard>
              <Descriptions title="" column={1}>
                <Descriptions.Item label="Band">{nrlte?.Status?.NR?.NR_Band || '-'}</Descriptions.Item>
                <Descriptions.Item label="NR-ARFCN">{nrlte?.Status?.NR?.NR_ARFCN || '-'}</Descriptions.Item>
                <Descriptions.Item label="Bandwidth">{nrlte?.Status?.NR?.NR_Bandwidth || '-'}</Descriptions.Item>
                <Descriptions.Item label="RSSI">{nrlte?.Status?.NR?.SSB_RSSI || '-'}</Descriptions.Item>
                <Descriptions.Item label="RSRP0">{nrlte?.Status?.NR?.SSB_RSRP || '-'}</Descriptions.Item>
                <Descriptions.Item label="RSRQ">{nrlte?.Status?.NR?.SSB_RSRQ || '-'}</Descriptions.Item>
                <Descriptions.Item label="SINR">{nrlte?.Status?.NR?.SSB_SINR || '-'}</Descriptions.Item>
                <Descriptions.Item label="TX Power">{nrlte?.Status?.NR?.NR_TXPower || '-'}</Descriptions.Item>
              </Descriptions>
            </ProCard>
            <ProCard>
              <Descriptions title="" column={1}>
                <Descriptions.Item label="PCI">{nrlte?.Status?.NR?.NR_PCI || '-'}</Descriptions.Item>
                <Descriptions.Item label="NCGI">{nrlte?.Status?.NR?.NCGI || '-'}</Descriptions.Item>
                <Descriptions.Item label="gNBID">{nrlte?.Status?.NR?.GNBID || '-'}</Descriptions.Item>
                <Descriptions.Item label="Cell ID">{nrlte?.Status?.NR?.NR_CellID || '-'}</Descriptions.Item>
                <Descriptions.Item label="MCC">{nrlte?.Status?.PCC?.MCC || '-'}</Descriptions.Item>
                <Descriptions.Item label="MNC">{nrlte?.Status?.PCC?.MNC || '-'}</Descriptions.Item>
                <Descriptions.Item label="CQI">{nrlte?.Status?.NR?.NR_CQI || '-'}</Descriptions.Item>
                <Descriptions.Item label="Rank">{nrlte?.Status?.NR?.NR_Rank || '-'}</Descriptions.Item>
                <Descriptions.Item label="SSB BeamID">{nrlte?.Status?.NR?.SSB_BeamID || '-'}</Descriptions.Item>
              </Descriptions>
            </ProCard>
          </ProCard>
        </ProCard>
      )
    } else {
      return (
        <ProCard title="Status">
          <ProCard split="vertical">
            <ProCard>
              <Descriptions title="" column={1}>
                <Descriptions.Item label="Band">{nrlte?.Status?.LTE?.Band || '-'}</Descriptions.Item>
                <Descriptions.Item label="EARFCN">{nrlte?.Status?.LTE?.EARFCN || '-'}</Descriptions.Item>
                <Descriptions.Item label="Bandwidth">{nrlte?.Status?.LTE?.Bandwidth || '-'}</Descriptions.Item>
                <Descriptions.Item label="RSSI">{nrlte?.Status?.LTE?.RSSI || '-'}</Descriptions.Item>
                <Descriptions.Item label="RSRP0">{nrlte?.Status?.LTE?.RSRP || '-'}</Descriptions.Item>
                <Descriptions.Item label="RSRQ">{nrlte?.Status?.LTE?.RSRQ || '-'}</Descriptions.Item>
                <Descriptions.Item label="SINR">{nrlte?.Status?.LTE?.SINR || '-'}</Descriptions.Item>
                <Descriptions.Item label="TX Power">{nrlte?.Status?.LTE?.TXPower || '-'}</Descriptions.Item>
              </Descriptions>
            </ProCard>
            <ProCard>
              <Descriptions title="" column={1}>
                <Descriptions.Item label="PCI">{nrlte?.Status?.LTE?.PCI || '-'}</Descriptions.Item>
                <Descriptions.Item label="ECGI">{nrlte?.Status?.LTE?.ECGI || '-'}</Descriptions.Item>
                <Descriptions.Item label="eNBID">{nrlte?.Status?.LTE?.ENBID || '-'}</Descriptions.Item>
                <Descriptions.Item label="Cell ID">{nrlte?.Status?.LTE?.CellID || '-'}</Descriptions.Item>
                <Descriptions.Item label="MCC">{nrlte?.Status?.PCC?.MCC || '-'}</Descriptions.Item>
                <Descriptions.Item label="MNC">{nrlte?.Status?.PCC?.MNC || '-'}</Descriptions.Item>
                <Descriptions.Item label="CQI">{nrlte?.Status?.LTE?.CQI || '-'}</Descriptions.Item>
                <Descriptions.Item label="Rank">{nrlte?.Status?.LTE?.Rank || '-'}</Descriptions.Item>
              </Descriptions>
            </ProCard>
          </ProCard>
        </ProCard>
      )
    }

  }

  return (
    <ProCard split="horizontal">
      <ProCard title="Settings">
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
          <Form.Item name="ConnectMethod" label="Connect Mode">
            <Select placeholder="Select network mode">
              {connectModeOptions.map((option) => (
                <Select.Option key={option.value} value={option.value}>
                  {option.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
        </ProForm>
      </ProCard>
      {renderStatus()}
    </ProCard>


  );
};

export default NRLTENetworkSettingsForm;

