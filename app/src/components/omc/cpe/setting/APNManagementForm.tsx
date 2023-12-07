import { Col, Form, Input, Row, Select, Space } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { useEffect, useState } from "react";
import { setDeviceParameterValues } from "@/models/device";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const authTypeOptions = [
  { label: 'None', value: 0 },
  { label: 'PAP', value: 1 },
  { label: 'CHAP', value: 2 },
];

const pdnTypeOptions = [
  { label: 'IPv4', value: "0" },
  { label: 'IPv6', value: "1" },
  { label: 'IPv4v6', value: "2" },
];

interface APNManagementFormProps {
  device: API.Device,
  data: Record<string, any>,
};

interface apnItem {
  key: string,
  label: string,
  profileName: string,
  enable: string,
  defaultGateway: string,
}

const APNManagementForm: React.FC<APNManagementFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [apnItems, setApnItems] = useState<apnItem[]>([]);
  const [selectedApn, setSelectedApn] = useState<string>("1");

  useEffect(() => {
    const apnManagement:Record<string, any> = data?.WEB_GUI?.Network?.APNManagement || {};
    const apnList:Record<string, Record<string, any>> = data?.WEB_GUI?.Network?.APNManagement?.APNList || {};

    let items: apnItem[] = []
    if (apnList) {
      Object.entries(apnList).forEach(([k, v]) => {
        items.push({
          key: k,
          label: k,
          profileName: v.Label || "-",
          enable: v.Enable ? 'Enable' : 'Disabled' || '-',
          defaultGateway: apnManagement?.DefaultGateway === k ? 'Enable' : '-',
        })
      });
    }
    setApnItems(items)
  }, [data]);

  const onFinish = async (data2: Record<string, any> & { DefaultGateway: boolean, Number: string }) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WEB_GUI.Network.APNManagement.APNList.' + data2.Number + "."
    values[prefixName + 'Enable'] = data2.Enable
    values[prefixName + 'Lable'] = data2.Label
    values[prefixName + 'APNName'] = data2.APNName
    values[prefixName + 'AuthType'] = data2.AuthType
    values[prefixName + 'Username'] = data2.Username
    values[prefixName + 'Password'] = data2.Password
    values[prefixName + 'PDNType'] = data2.PDNType
    setDeviceParameterValues(device, values)
  };
  const updateForm = () => {
    const apnManagement:Record<string, any> = data?.WEB_GUI?.Network?.APNManagement || {};
    const apnList:Record<string, Record<string, any>> = data?.WEB_GUI?.Network?.APNManagement?.APNList || {};
    if (apnList) {
      const apn = apnList[selectedApn];
      const values = { ...apn, DefaultGateway: apnManagement?.DefaultGateway === selectedApn, Number: selectedApn }
      form.setFieldsValue(values);
    }
  }

  const onReset = () => {
    updateForm();
  };

  useEffect(() => {
    updateForm();
  }, [selectedApn, data]);

  const apnListColumns = [
    { title: "#", dataIndex: "key", key: "key" },
    { title: "Profile Name", dataIndex: "profileName", key: "profileName" },
    { title: "Enable", dataIndex: "enable", key: "enable" },
    { title: "Default Gateway", dataIndex: "defaultGateway", key: "defaultGateway" },
  ]

  return (
    <ProCard split="horizontal">
      <ProCard title="APN List">
        <ProTable dataSource={apnItems} columns={apnListColumns} pagination={false} search={false} options={false} />;
      </ProCard>
      <ProCard title="APN Settings">
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
          <Form.Item key="Number" name="Number" label="APN Number">
            <Select placeholder="Select APN Number" onChange={(v) => { setSelectedApn(v) }}>
              {apnItems.map((item) => (
                <Select.Option key={item.key} value={item.key} >
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item name="Enable" label="Enable" >
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item key="Lable" name="Lable" label="Profile Name">
            <Input placeholder="Enter profile name" />
          </Form.Item>
          <Form.Item key="APNName" name="APNName" label="APN Name">
            <Input placeholder="Enter APN name" />
          </Form.Item>
          <Form.Item key="AuthType" name="AuthType" label="Authentication Type">
            <Select placeholder="Select authentication type">
              {authTypeOptions.map((item) => (
                <Select.Option key={item.value} value={item.value}>
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item key="PDNType" name="PDNType" label="PDN Type">
            <Select placeholder="Select PDN type">
              {pdnTypeOptions.map((item) => (
                <Select.Option key={item.value} value={item.value}>
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item key="Username" name="Username" label="Username">
            <Input placeholder="Enter username" />
          </Form.Item>
          <Form.Item key="Password" name="Password" label="Password">
            <Input.Password placeholder="Enter password" />
          </Form.Item>
          <Form.Item key="DefaultGateway" name="DefaultGateway" label="Default Gateway" >
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
        </ProForm >
      </ProCard>
    </ProCard>
  );
};

export default APNManagementForm;
