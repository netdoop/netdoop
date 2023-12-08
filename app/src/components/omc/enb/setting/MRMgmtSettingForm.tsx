import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface Props {
  device: API.Device,
  data: Record<string, any>,
};

const MRMgmtSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const [items, setItems] = useState<Record<string, any>[]>([]);
  const [selected, setSelected] = useState<string | undefined>(undefined);

  const onFinish = async (data2: Record<string, Record<string, any>>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.FAP.MRMgmt.'
    const configs = data2;
    Object.entries(configs).forEach(([i, v]) => {
      values[prefixName + 'Config.' + i + '.MrEnable'] = v.MrEnable;
      values[prefixName + 'Config.' + i + '.MrPassword'] = v.MrPassword;
      values[prefixName + 'Config.' + i + '.MrUsername'] = v.MrUsername;
      values[prefixName + 'Config.' + i + '.MrUrl'] = v.MrUrl;
      values[prefixName + 'Config.' + i + '.UploadPeriod'] = v.UploadPeriod;
      values[prefixName + 'Config.' + i + '.MeasureType'] = v.MeasureType;
      values[prefixName + 'Config.' + i + '.OmcName'] = v.OmcName;
      values[prefixName + 'Config.' + i + '.SampleBeginTime'] = v.SampleBeginTime;
      values[prefixName + 'Config.' + i + '.SampleEndTime'] = v.SampleEndTime;
      values[prefixName + 'Config.' + i + '.SamplePeriod'] = v.SamplePeriod;
      values[prefixName + 'Config.' + i + '.SubFrameNum'] = v.SubFrameNum;
      values[prefixName + 'Config.' + i + '.PrbNum'] = v.PrbNum;
    });

    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const mrMgmtConfig:Record<string, any> = data?.FAP?.MRMgmt?.Config || {};
    form.setFieldsValue(mrMgmtConfig);
  };
  useEffect(() => {
    const mrMgmtConfig:Record<string, any> = data?.FAP?.MRMgmt?.Config || {};
    form.setFieldsValue(mrMgmtConfig);
    let items: Record<string, any>[] = [];
    if (mrMgmtConfig) {
      Object.keys(mrMgmtConfig).forEach(key => {
        const value = mrMgmtConfig[key];
        const item: Record<string, any> = { ...value, key };
        items = [...items, item];
      });
    }
    setItems(items)
  }, [data]);

  const configListColumns = [
    { title: "#", dataIndex: "key", key: "key" },
    { title: "Measure Type", dataIndex: "MeasureType", key: "MeasureType" },
    {
      title: "Enable", dataIndex: "MrEnable", key: "MrEnable",
      render: (text: any, record: Record<string, any>) => {
        if (record.MrEnable === '0') {
          return ("Disable")
        } else {
          return ("Enable")
        }
      }
    },
    { title: "URL", dataIndex: "MrUrl", key: "MrUrl" },
    { title: "Username", dataIndex: "MrUsername", key: "MrUsername" },
    { title: "Password", dataIndex: "MrPassword", key: "MrPassword" },
    { title: "Upload Interval", dataIndex: "UploadPeriod", key: "UploadPeriod" },
    { title: "Sample Interval", dataIndex: "SamplePeriod", key: "SamplePeriod" },
    {
      title: "Edit",
      dataIndex: "key",
      key: "edit",
      render: (text: any, record: Record<string, any>) => (
        <a onClick={() => setSelected(record.key)}>Edit</a>
      ),
    },
  ];
  return (
    <ProCard split="horizontal">
      <ProCard title="Configuration List">
        <ProTable dataSource={items} columns={configListColumns} pagination={false} search={false} options={false} />;
      </ProCard>
      {selected && (<ProForm
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
        <Form.Item name={["Config", selected, "MrEnable"]} label="Upload Enable">
          <Select >
            <Select.Option key="0" value="0">Disable</Select.Option>
            <Select.Option key="1" value="1">Enable</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item name={["Config", selected, "MeasureType"]} label="Measure Type">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "MrUrl"]} label="URL">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "MrUsername"]} label="Username">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "MrPassword"]} label="Password">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "UploadPeriod"]} label="Upload Interval">
          <InputNumber addonAfter="sec" />
        </Form.Item>
        <Form.Item name={["Config", selected, "SamplePeriod"]} label="Sample Interval">
          <InputNumber addonAfter="sec" />
        </Form.Item>
        <Form.Item name={["Config", selected, "SampleBeginTime"]} label="Sample Begin Time">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "SampleEndTime"]} label="Sample End Time">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "SubFrameNum"]} label="Sub Frame Number">
          <InputNumber />
        </Form.Item>
        <Form.Item name={["Config", selected, "PrbNum"]} label="PRB Number">
          <InputNumber />
        </Form.Item>
        <Form.Item name={["Config", selected, "OmcName"]} label="OMC Name">
          <Input />
        </Form.Item>
      </ProForm>
      )
      }
    </ProCard >
  );
};

export default MRMgmtSettingForm;
