
import { Col, Form, Row, Input, Space, Select, MenuProps, Button, Dropdown } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { setDeviceParameterValues, addDeviceObject, deleteDeviceObject } from "@/models/device";
import { useEffect, useState } from "react";
import { EditOutlined, MoreOutlined, PlusOutlined } from "@ant-design/icons";
import { useIntl } from "@umijs/max";
import { proTableLayout } from "@/constants/style";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const intefaceOptions = [
  { label: 'LAN', value: 'LAN' },
  { label: 'APN1', value: 'APN1' },
  { label: 'APN2', value: 'APN2' },
  { label: 'APN3', value: 'APN3' },
  { label: 'APN4', value: 'APN4' },
];

interface StaticRouteFormProps {
  device: API.Device,
  data?: Record<string, any>,
};

interface staticRouteItem {
  Key: string,
  Index?: number,
  DestIPAddress?: string,
  SubnetMask?: string,
  Interface?: string,
  Gateway?: string,
  Status?: string,
};

const StaticRouteForm: React.FC<StaticRouteFormProps> = ({
  device,
  data,
}) => {
  const intl = useIntl()
  const [form] = Form.useForm();
  const [selectedStaticRoute, setSelectedStaticRoute] = useState<number | undefined>(undefined);
  const [staticRouteItems, setStaticRouteItems] = useState<staticRouteItem[]>([]);
  const prefixName = 'Device.WEB_GUI.Network.StaticRoute.';

  useEffect(() => {
    const list:Record<string, Record<string, any>> = data?.WEB_GUI?.Network?.StaticRoute?.List || {};
    let items: staticRouteItem[] = []
      Object.entries(list).forEach(([k, v]) => {
        items.push({
          Key: k,
          Index: v.index,
          DestIPAddress: v.DestIPAddress,
          SubnetMask: v.SubnetMask,
          Interface: v.Interface,
          Gateway: v.Gateway,
          Status: v.Status,
        })
      });
    setStaticRouteItems(items)
  }, [data])

  const updateForm = () => {
    if (selectedStaticRoute === undefined) {
      form.resetFields()
    } else {
      form.setFieldsValue(staticRouteItems[selectedStaticRoute]);
    }
  }

  const onFinish = async () => {
    let values: Record<string, any> = {};
    // const prefixName = 'Device.WEB_GUI.Network.StaticRoutes.';
    // values[prefixName+'LANHost.IPAddress'] = data2.LANHost?.IPAddress;
    // values[prefixName+'LANHost.SubnetMask'] = data2.LANHost?.SubnetMask;

    // values[prefixName+'DHCP.ServerEnable'] = data2.DHCP?.ServerEnable;
    // values[prefixName+'DHCP.StartIP'] = data2.DHCP?.StartIP;
    // values[prefixName+'DHCP.EndIP'] = data2.DHCP?.EndIP;
    // values[prefixName+'DHCP.LeaseTime'] = data2.DHCP?.LeaseTime;

    setDeviceParameterValues(device, values);
  };
  const onReset = () => {
    updateForm()
  };

  const handleAdd = async () => {
    addDeviceObject(device, prefixName + "List.")
  };

  const handleEdit = async (record: staticRouteItem) => {
    setSelectedStaticRoute(Number(record.Key))
  };

  const handleDelete = async (record: staticRouteItem) => {
    deleteDeviceObject(device, prefixName + "List." + record.Key + ".")
  };

  const moreItems = (record: staticRouteItem): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'edit',
        icon: (<EditOutlined />),
        label: (
          <a onClick={() => handleEdit(record)}>
            {intl.formatMessage({ id: 'common.edit' })}
          </a>
        ),
      },
      {
        key: 'delete',
        icon: (<EditOutlined />),
        label: (
          <a onClick={() => handleDelete(record)}>
            {intl.formatMessage({ id: 'common.delete' })}
          </a>
        ),
      },
    ]
    return { items }
  }
  const staticRouteListColumns = [
    {
      title: "#", dataIndex: "Index", key: "Index",
      render: (text: any, record: staticRouteItem) => (
        <div>
          {text}
          <Dropdown menu={moreItems(record)}>
            <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
              <MoreOutlined />
            </Button>
          </Dropdown>
        </div>
      ),
    },
    { title: "DestIPAddress", dataIndex: "DestIPAddress", key: "DestIPAddress" },
    { title: "SubnetMask", dataIndex: "SubnetMask", key: "SubnetMask" },
    { title: "Interface", dataIndex: "Interface", key: "Interface" },
    { title: "Gateway", dataIndex: "Gateway", key: "Gateway" },
    { title: "Status", dataIndex: "Status", key: "Status" },
  ]

  return (
    <ProCard split="horizontal">
      <ProCard title="Static Route List">
        <ProTable
          {...proTableLayout}
          dataSource={staticRouteItems}
          columns={staticRouteListColumns}
          pagination={false}
          search={false}
          options={false}
          toolBarRender={() => [
            <Space key="custom-options">
              <PlusOutlined
                className="ant-pro-table-toolbar-item-iconButton"
                onClick={() => handleAdd()}
                style={{ fontSize: '12px' }}
              />
            </Space>
          ]}
        />;
      </ProCard>
      <ProCard title="Static Route Settings">
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
          <Form.Item key="DestIPAddress" name="DestIPAddress" label="Destination IP">
            <Input placeholder="Enter Destination IP Address" />
          </Form.Item>
          <Form.Item key="SubnetMask" name="SubnetMask" label="Subnet Mask">
            <Input placeholder="Enter Subnet Mask" />
          </Form.Item>
          <Form.Item name="Interface" label="Interface">
            <Select placeholder="Select Interface">
              {intefaceOptions.map((option) => (
                <Select.Option key={option.value} value={option.value}>
                  {option.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item key="Gateway" name="Gateway" label="Gateway">
            <Input placeholder="Enter Gateway" />
          </Form.Item>
        </ProForm>
      </ProCard>
    </ProCard>
  );
};

export default StaticRouteForm;

