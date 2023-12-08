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

const PerfMgmtSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const [items, setItems] = useState<Record<string, any>[]>([]);
  const [selected, setSelected] = useState<string | undefined>(undefined);

  const onFinish = async (data2: Record<string, Record<string, any>>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.FAP.PerfMgmt.';
    const configs = data2;
    Object.entries(configs).forEach(([i, v]) => {
      values[prefixName + 'Config.' + i + '.Alias'] = v.Alias;
      values[prefixName + 'Config.' + i + '.Enable'] = v.Enable;
      values[prefixName + 'Config.' + i + '.Password'] = v.Password;
      values[prefixName + 'Config.' + i + '.Username'] = v.Username;
      values[prefixName + 'Config.' + i + '.URL'] = v.URL;
      values[prefixName + 'Config.' + i + '.PeriodicUploadInterval'] = v.PeriodicUploadInterval;
      values[prefixName + 'Config.' + i + '.PeriodicUploadTime'] = v.PeriodicUploadTime;
    });
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const perfMgmtConfig:Record<string, any> = data?.FAP?.PerfMgmt?.Config || {};
    form.setFieldsValue(perfMgmtConfig);
  };
  useEffect(() => {
    const perfMgmtConfig:Record<string, any> = data?.FAP?.PerfMgmt?.Config || {};
    form.setFieldsValue(perfMgmtConfig);
    let updateItems: Record<string, any>[] = [];
    if (perfMgmtConfig) {
      Object.entries(perfMgmtConfig).forEach((entry) => {
        const [key, value] = entry;
        const item: Record<string, any> = { ...value, key };
        updateItems.push(item);
      });
    }
    setItems(updateItems);
  }, [data]);

  const configListColumns = [
    { title: "#", dataIndex: "key", key: "key" },
    { title: "Alias", dataIndex: "Alias", key: "Alias" },
    {
      title: "Enable", dataIndex: "Enable", key: "Enable",
      render: (text: any, record: Record<string, any>) => {
        if (record.Enable === '0') {
          return ("Disable")
        } else {
          return ("Enable")
        }
      }
    },
    { title: "URL", dataIndex: "URL", key: "URL" },
    { title: "Username", dataIndex: "Username", key: "Username" },
    { title: "Password", dataIndex: "Password", key: "Password" },
    { title: "Periodic Upload Interval", dataIndex: "PeriodicUploadInterval", key: "PeriodicUploadInterval" },
    { title: "Periodic Upload Time", dataIndex: "PeriodicUploadTime", key: "PeriodicUploadTime" },
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
        <Form.Item name={["Config", selected, "Enable"]} label="Upload Enable">
          <Select >
            <Select.Option key="0" value="0">Disable</Select.Option>
            <Select.Option key="1" value="1">Enable</Select.Option>
          </Select>
        </Form.Item>
        <Form.Item name={["Config", selected, "Alias"]} label="Alias">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "URL"]} label="URL">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "Username"]} label="Username">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "Password"]} label="Password">
          <Input />
        </Form.Item>
        <Form.Item name={["Config", selected, "PeriodicUploadInterval"]} label="Periodic Upload Interval">
          <InputNumber addonAfter="sec" />
        </Form.Item>
        <Form.Item name={["Config", selected, "PeriodicUploadTime"]} label="Periodic Upload Time">
          <Input />
        </Form.Item>
      </ProForm>
      )
      }
    </ProCard >
  );
};

export default PerfMgmtSettingForm;
