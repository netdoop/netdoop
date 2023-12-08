import { Col, Form, Input, Row, Select, Space } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";
import { DefaultOptionType } from "antd/es/select";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface Props {
  device: API.Device,
  data: Record<string, any>,
};

const EPCSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [form2] = Form.useForm();

  const [mmeInfos, setMmeInfos] = useState<Record<string, Record<string, any>>>({});
  const [mmeKeyOptions, setMmeKeyOptions] = useState<DefaultOptionType[]>([])
  const [mmeKeySelected, setMmeKeySelected] = useState<string>("1")

  const [plmnInfos, setPlmnInfos] = useState<Record<string, Record<string, any>>>({});
  const [plmnKeySelected, setPlmnKeySelected] = useState<string | undefined>(undefined)

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.'

    values[prefixName + 'CellConfig.LTE.RAN.Common.CellIdentity'] = data2.CellConfig?.LTE?.RAN?.Common?.CellIdentity;
    values[prefixName + 'CellConfig.LTE.RAN.Common.EnbType'] = data2.CellConfig?.LTE?.RAN?.Common?.EnbType;
    values[prefixName + 'CellConfig.LTE.EPC.TAC'] = data2.CellConfig?.LTE?.EPC?.TAC;

    values[prefixName + 'FAPControl.LTE.Gateway.S1SigLinkPort'] = data2.FAPControl?.LTE?.Gateway?.S1SigLinkPort;

    Object.keys(mmeInfos).forEach(k => {
      const info = data2.FAPControl?.LTE?.Gateway?.MME_Comm_Info?.[k] as Record<string, any>;
      values[prefixName + 'FAPControl.LTE.Gateway.MME_Comm_Info.' + k + '.NumIpAddr'] = info.NumIpAddr;
      values[prefixName + 'FAPControl.LTE.Gateway.MME_Comm_Info.' + k + '.IPAddrMain'] = info.IPAddrMain;
      values[prefixName + 'FAPControl.LTE.Gateway.MME_Comm_Info.' + k + '.IPAddrSpare'] = info.IPAddrSpare;
      values[prefixName + 'FAPControl.LTE.Gateway.MME_Comm_Info.' + k + '.RelOfMME'] = info.RelOfMME;
    });
    setDeviceParameterValues(device, values)
  };

  const onFinish2 = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.Services.FAPService.1.'

    Object.keys(plmnInfos).forEach(k => {
      const info = data2.CellConfig?.LTE?.EPC?.PLMNList?.[k] as Record<string, any>;
      values[prefixName + 'CellConfig.LTE.EPC.PLMNList.' + k + '.PLMNID'] = info.PLMNID;
      values[prefixName + 'CellConfig.LTE.EPC.PLMNList.' + k + '.CellReservedForOperatorUse'] = info.CellReservedForOperatorUse;
    });
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const fapService = data?.Services?.FAPService?.["1"] || {};
    form.setFieldsValue(fapService);
    form2.setFieldsValue(fapService);
  };

  useEffect(() => {
    const fapService = data?.Services?.FAPService?.["1"] || {};
    form.setFieldsValue(fapService);
    form2.setFieldsValue(fapService);

    const keys = Object.keys(fapService?.FAPControl?.LTE?.Gateway?.MME_Comm_Info || {})
    const infos: Record<string, any> = {};
    keys.forEach(k => {
      const info: Record<string, any> = fapService?.FAPControl?.LTE?.Gateway?.MME_Comm_Info?.[k];
      infos[k] = { key: k, index: k, ...info };
    });
    setMmeInfos(infos);

    const options = keys.map(v => ({ label: v, value: v, }))
    setMmeKeyOptions(options)
    if (keys.length > 0) {
      const key = keys[0];
      setMmeKeySelected(key)
    }

    const infos2: Record<string, any> = Object.keys(fapService?.CellConfig?.LTE?.EPC?.PLMNList || {}).map(k => {
      const info: Record<string, any> = fapService?.CellConfig?.LTE?.EPC?.PLMNList?.[k];
      return { key: k, index: k, enable: "0", isPrimary: "0", ...info };
    });
    setPlmnInfos(infos2);
  }, [data]);

  const setMmeIpNum = (key: string, value: string) => {
    const update = mmeInfos;
    update[key].NumIpAddr = value;
    setMmeInfos(update)
  };

  const plmnListColumns = [
    { title: "#", dataIndex: "index", key: "index" },
    {
      title: "Enable", dataIndex: "Enable", key: "Enable",
      render: (text: any, record: Record<string, any>) => {
        if (record.enable === '1') {
          return ("Enable")
        } else {
          return ("Disable")
        }
      }
    },
    {
      title: "IsPrimary", dataIndex: "IsPrimary", key: "IsPrimary",
      render: (text: any, record: Record<string, any>) => {
        if (record.isPrimary === '1') {
          return ("Enable")
        } else {
          return ("Disable")
        }
      }
    },
    { title: "PLMNID", dataIndex: "PLMNID", key: "PLMNID" },
    { title: "CellReservedForOperatorUse", dataIndex: "CellReservedForOperatorUse", key: "CellReservedForOperatorUse" },
    {
      title: "Edit",
      dataIndex: "key",
      key: "edit",
      render: (text: any, record: Record<string, any>) => (
        <a onClick={() => setPlmnKeySelected(record.key)}>Edit</a>
      ),
    },
  ];
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
          <Form.Item name={["CellConfig", "LTE", "RAN", "Common", "CellIdentity"]} label="Cell Identify">
            <Input />
          </Form.Item>
          <Form.Item name={["CellConfig", "LTE", "RAN", "Common", "EnbType"]} label="Cell Type">
            <Input />
          </Form.Item>
          <Form.Item name={["CellConfig", "LTE", "EPC", "TAC"]} label="TAC">
            <Input />
          </Form.Item>
          <Form.Item name="S1-MM" label="S1-MME">
            <Input disabled />
          </Form.Item>
          <Form.Item name={["FAPControl", "LTE", "Gateway", "S1SigLinkPort"]} label="S1 Port">
            <Input />
          </Form.Item>
          <Form.Item name="MMENumber" label="The Number of MME" initialValue={mmeKeySelected}>
            <Select options={mmeKeyOptions} onChange={setMmeKeySelected} />
          </Form.Item>
          {Object.values(mmeInfos).map(v => (
            mmeKeySelected === v.key && (
              <ProCard key={v.key} title={"MME Configuration " + v.index}>
                <Form.Item name={["FAPControl", "LTE", "Gateway", "MME_Comm_Info", v.key, "NumIpAddr"]} label="MME IP Number">
                  <Select placeholder="" onChange={(value) => setMmeIpNum(v.key, value)}>
                    <Select.Option key="1" value="1">1</Select.Option>
                    <Select.Option key="2" value="2">2</Select.Option>
                  </Select>
                </Form.Item>
                <Form.Item name={["FAPControl", "LTE", "Gateway", "MME_Comm_Info", v.key, "IPAddrMain"]} label="MME Primary IP">
                  <Input />
                </Form.Item>
                {v.NumIpAddr === "2" && (
                  <Form.Item name={["FAPControl", "LTE", "Gateway", "MME_Comm_Info", v.key, "IPAddrSpare"]} label="MME Secondary IP">
                    <Input />
                  </Form.Item>
                )}
                <Form.Item name={["FAPControl", "LTE", "Gateway", "MME_Comm_Info", v.key, "RelOfMME"]} label="Rel Of MME">
                  <Input />
                </Form.Item>
              </ProCard>
            )
          ))}
        </ProForm >
      </ProCard>
      <ProCard split="horizontal">
        <ProCard title="PLMN List">
          <ProTable dataSource={Object.values(plmnInfos)} columns={plmnListColumns} pagination={false} search={false} options={false} />
        </ProCard>

        {plmnKeySelected && (
          <ProCard title={"PLMN " + plmnKeySelected}>
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
              <Form.Item name={["CellConfig", "LTE", "EPC", "PLMNList", plmnKeySelected, "PLMNID"]} label="PLMNID">
                <Input />
              </Form.Item>
              <Form.Item name={["CellConfig", "LTE", "EPC", "PLMNList", plmnKeySelected, "CellReservedForOperatorUse"]} label="Cell Reserved For Operator Use">
                <Select >
                  <Select.Option key="0" value="0">Disable</Select.Option>
                  <Select.Option key="1" value="1">Enable</Select.Option>
                </Select>
              </Form.Item>
            </ProForm>
          </ProCard>
        )}
      </ProCard>
    </ProCard>
  );
};

export default EPCSettingForm;
