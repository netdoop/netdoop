import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
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

const WANSettingForm: React.FC<Props> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();

  const [vlanType, setVlanType] = useState<string>("0")
  const [ipAddrType, setIpAddrType] = useState<string>("0")
  const [vlan1IpAddrType, setVlan1IpAddrType] = useState<string>("0")
  const [vlan2IpAddrType, setVlan2IpAddrType] = useState<string>("0")

  const [extraWanIp1Enable, setExtraWanIp1Enable] = useState<string>('0')
  const [extraWanIp2Enable, setExtraWanIp2Enable] = useState<string>('0')
  const [extraWanIp3Enable, setExtraWanIp3Enable] = useState<string>('0')
  const [extraWanIp4Enable, setExtraWanIp4Enable] = useState<string>('0')

  const [extraWanIp1AddrType, setExtraWanIp1AddrType] = useState<string>("0")
  const [extraWanIp2AddrType, setExtraWanIp2AddrType] = useState<string>("0")
  const [extraWanIp3AddrType, setExtraWanIp3AddrType] = useState<string>("0")
  const [extraWanIp4AddrType, setExtraWanIp4AddrType] = useState<string>("0")

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WAN.'

    values[prefixName + 'X_VENDOR_WAN_VLAN_TYPE'] = data2.WAN.X_VENDOR_WAN_VLAN_TYPE
    if (vlanType === "0") {
      values[prefixName + 'X_VENDOR_IPADDR_METHOD'] = data2.WAN.X_VENDOR_IPADDR_METHOD
      values[prefixName + 'X_VENDOR_IPADDR_STATIC'] = data2.WAN.X_VENDOR_IPADDR_STATIC
      values[prefixName + 'X_VENDOR_IPADDR_MASK'] = data2.WAN.X_VENDOR_IPADDR_MASK
      values[prefixName + 'X_VENDOR_DEFAULT_GW'] = data2.WAN.X_VENDOR_DEFAULT_GW
    } else if (vlanType === "1") {
      values[prefixName + 'X_VENDOR_VLAN1_IPADDR_METHOD'] = data2.WAN.X_VENDOR_VLAN1_IPADDR_METHOD
      values[prefixName + 'X_VENDOR_VLAN1_IPADDR_STATIC'] = data2.WAN.X_VENDOR_VLAN1_IPADDR_STATIC
      values[prefixName + 'X_VENDOR_VLAN1_IPADDR_MASK'] = data2.WAN.X_VENDOR_VLAN1_IPADDR_MASK
      values[prefixName + 'X_VENDOR_VLAN1_DEFAULT_GW'] = data2.WAN.X_VENDOR_VLAN1_DEFAULT_GW
    } else if (vlanType === "2") {
      values[prefixName + 'X_VENDOR_VLAN1_IPADDR_METHOD'] = data2.WAN.X_VENDOR_VLAN1_IPADDR_METHOD
      values[prefixName + 'X_VENDOR_VLAN1_IPADDR_STATIC'] = data2.WAN.X_VENDOR_VLAN1_IPADDR_STATIC
      values[prefixName + 'X_VENDOR_VLAN1_IPADDR_MASK'] = data2.WAN.X_VENDOR_VLAN1_IPADDR_MASK
      values[prefixName + 'X_VENDOR_VLAN1_DEFAULT_GW'] = data2.WAN.X_VENDOR_VLAN1_DEFAULT_GW
      values[prefixName + 'X_VENDOR_VLAN2_IPADDR_METHOD'] = data2.WAN.X_VENDOR_VLAN2_IPADDR_METHOD
      values[prefixName + 'X_VENDOR_VLAN2_IPADDR_STATIC'] = data2.WAN.X_VENDOR_VLAN2_IPADDR_STATIC
      values[prefixName + 'X_VENDOR_VLAN2_IPADDR_MASK'] = data2.WAN.X_VENDOR_VLAN2_IPADDR_MASK
      values[prefixName + 'X_VENDOR_VLAN2_DEFAULT_GW'] = data2.WAN.X_VENDOR_VLAN2_DEFAULT_GW
    }
    values[prefixName + 'X_VENDOR_WAN_MODE_CONFIG'] = data2.WAN.X_VENDOR_WAN_MODE_CONFIG;
    values[prefixName + 'X_VENDOR_WAN_MTU_CONFIG'] = data2.WAN.X_VENDOR_WAN_MTU_CONFIG;

    const extras:Record<string, Record<string, any>> = data2.WANDevice?.WAN?.EXTRA_IP_LIST || {};
    Object.entries(extras).forEach(([i, v]) => {
      values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_CONFIG'] = v.X_VENDOR_EXTRA_IP_CONFIG;
      if (v.X_VENDOR_EXTRA_IP_CONFIG === '1') {
        values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_VLAN1_VID'] = v.X_VENDOR_EXTRA_IP_VLAN1_VID;
        values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG'] = v.X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG;
        values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_IPADDR_METHOD'] = v.X_VENDOR_EXTRA_IP_IPADDR_METHOD;
        if (v.X_VENDOR_EXTRA_IP_IPADDR_METHOD === '1') {
          values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_IPADDR_STATIC'] = v.X_VENDOR_EXTRA_IP_IPADDR_STATIC;
          values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_IPADDR_MASK'] = v.X_VENDOR_EXTRA_IP_IPADDR_MASK;
          values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_DEFAULT_GW'] = v.X_VENDOR_EXTRA_IP_DEFAULT_GW;
          values[prefixName + 'EXTRA_IP_LIST.' + i + '.X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP'] = v.X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP;
        }
      }
    });

    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const wan:Record<string, any> = data?.WAN || {};
    form.setFieldsValue(wan);
  };
  useEffect(() => {
    const wan:Record<string, any> = data?.WAN || {};
    const extraIpList:Record<string,Record<string, any>> = data?.WAN?.EXTRA_IP_LIST || {};
    form.setFieldsValue(wan);
    setVlanType(wan.X_VENDOR_WAN_VLAN_TYPE || '0')
    setIpAddrType(wan.X_VENDOR_IPADDR_METHOD || '0')
    setVlan1IpAddrType(wan.X_VENDOR_VLAN1_IPADDR_METHOD || '0')
    setVlan2IpAddrType(wan.X_VENDOR_VLAN2_IPADDR_METHOD || '0')

    setExtraWanIp1Enable(extraIpList["1"]?.X_VENDOR_EXTRA_IP_CONFIG || "0")
    setExtraWanIp2Enable(extraIpList["2"]?.X_VENDOR_EXTRA_IP_CONFIG || "0")
    setExtraWanIp3Enable(extraIpList["3"]?.X_VENDOR_EXTRA_IP_CONFIG || "0")
    setExtraWanIp4Enable(extraIpList["4"]?.X_VENDOR_EXTRA_IP_CONFIG || "0")

    setExtraWanIp1AddrType(extraIpList["1"]?.X_VENDOR_EXTRA_IP_IPADDR_METHOD || "0")
    setExtraWanIp2AddrType(extraIpList["2"]?.X_VENDOR_EXTRA_IP_IPADDR_METHOD || "0")
    setExtraWanIp3AddrType(extraIpList["3"]?.X_VENDOR_EXTRA_IP_IPADDR_METHOD || "0")
    setExtraWanIp4AddrType(extraIpList["4"]?.X_VENDOR_EXTRA_IP_IPADDR_METHOD || "0")
  }, [data]);

  return (
    <ProForm
      {...layout}
      form={form}
      layout="horizontal"
      onFinish={onFinish}
      onReset={onReset}
      labelWrap
      disabled
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
        <Form.Item name="X_VENDOR_WAN_VLAN_TYPE" label="VLAN Type">
          <Select placeholder="Select VLAN type" onChange={setVlanType}>
            <Select.Option key="0" value="0">Disable</Select.Option>
            <Select.Option key="1" value="1">Single</Select.Option>
            <Select.Option key="2" value="2">Dual</Select.Option>
          </Select>
        </Form.Item>
        {vlanType === "0" && (
          <>
            <Form.Item name="X_VENDOR_IPADDR_METHOD" label="IP Address Type">
              <Select placeholder="Select IP address type" onChange={setIpAddrType}>
                <Select.Option key="1" value="1">DHCP</Select.Option>
                <Select.Option key="0" value="0">Static</Select.Option>
              </Select>
            </Form.Item>
            {ipAddrType === "0" && (
              <>
                <Form.Item name="X_VENDOR_IPADDR_STATIC" label="IP Address">
                  <Input />
                </Form.Item>
                <Form.Item name="X_VENDOR_IPADDR_MASK" label="Subnet Mask">
                  <Input />
                </Form.Item>
                <Form.Item name="X_VENDOR_DEFAULT_GW" label="Default Gateway">
                  <Input />
                </Form.Item>
              </>
            )}
          </>
        )}
        {vlanType === "1" && (
          <>
            <ProCard title="Single VLAN">
              <Form.Item name="X_VENDOR_VLAN1_VID" label="VLAN ID">
                <Input />
              </Form.Item>
              <Form.Item name="X_VENDOR_VLAN1_IPADDR_METHOD" label="VLAN IP Address Type">
                <Select placeholder="Select IP address type" onChange={setVlan1IpAddrType}>
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {vlan1IpAddrType === "1" && (
                <>
                  <Form.Item name="X_VENDOR_VLAN1_IPADDR_STATIC" label="VLAN IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name="X_VENDOR_VLAN1_IPADDR_MASK" label="VLAN Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name="X_VENDOR_VLAN1_DEFAULT_GW" label="VLAN Default Gateway">
                    <Input />
                  </Form.Item>
                </>
              )}
            </ProCard>
          </>
        )}
        {vlanType === "2" && (
          <>
            <ProCard title="Bussiness VLAN">
              <Form.Item name="X_VENDOR_VLAN1_VID" label="VLAN ID">
                <Input />
              </Form.Item>
              <Form.Item name="X_VENDOR_VLAN1_IPADDR_METHOD" label="VLAN IP Address Type">
                <Select placeholder="Select IP address type" onChange={setVlan1IpAddrType}>
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {vlan1IpAddrType === "1" && (
                <>
                  <Form.Item name="X_VENDOR_VLAN1_IPADDR_STATIC" label="VLAN IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name="X_VENDOR_VLAN1_IPADDR_MASK" label="VLAN Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name="X_VENDOR_VLAN1_DEFAULT_GW" label="VLAN Default Gateway">
                    <Input />
                  </Form.Item>
                </>
              )}
            </ProCard>
            <ProCard title="Management VLAN">
              <Form.Item name="X_VENDOR_VLAN2_VID" label="VLAN ID">
                <Input />
              </Form.Item>
              <Form.Item name="X_VENDOR_VLAN2_IPADDR_METHOD" label="VLAN IP Address Type">
                <Select placeholder="Select IP address type" onChange={setVlan2IpAddrType}>
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {vlan2IpAddrType === "1" && (
                <>
                  <Form.Item name="X_VENDOR_VLAN2_IPADDR_STATIC" label="VLAN IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name="X_VENDOR_VLAN2_IPADDR_MASK" label="VLAN Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name="X_VENDOR_VLAN2_DEFAULT_GW" label="VLAN Default Gateway">
                    <Input />
                  </Form.Item>
                </>
              )}
            </ProCard>
          </>
        )}
        <ProCard title="WAN MTU/Mode">
          <Form.Item name="X_VENDOR_WAN_MTU_CONFIG" label="MTU">
            <InputNumber />
          </Form.Item>
          <Form.Item name="X_VENDOR_WAN_MODE_CONFIG" label="Mode">
            <Select placeholder="Select WAN mode">
              <Select.Option key="0" value="0">Auto</Select.Option>
            </Select>
          </Form.Item>
        </ProCard>

        <ProCard title="Extra WAN IP1">
          <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_CONFIG"]} label="Extra WAN IP1">
            <Select onChange={setExtraWanIp1Enable}>
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          {extraWanIp1Enable === "1" && (
            <>
              <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG"]} label="VLAN Tag Flag">
                <Input />
              </Form.Item>
              <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_IPADDR_METHOD"]} label="IP Address Tye">
                <Select placeholder="Select IP address type" onChange={setExtraWanIp1AddrType} >
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {extraWanIp1AddrType === "1" && (
                <>
                  <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_IPADDR_STATIC"]} label="IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_IPADDR_MASK"]} label="Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_DEFAULT_GW"]} label="Default Gateway">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","1","X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP"]} label="Destination Subnet">
                    <Input />
                  </Form.Item>
                </>
              )}
            </>
          )}
        </ProCard>

        <ProCard title="Extra WAN IP2">
          <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_CONFIG"]} label="Extra WAN IP2">
            <Select onChange={setExtraWanIp2Enable}>
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          {extraWanIp2Enable === "1" && (
            <>
              <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG"]} label="VLAN Tag Flag">
                <Input />
              </Form.Item>
              <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_IPADDR_METHOD"]} label="IP Address Tye">
                <Select placeholder="Select IP address type" onChange={setExtraWanIp2AddrType}>
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {extraWanIp2AddrType === "1" && (
                <>
                  <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_IPADDR_STATIC"]} label="IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_IPADDR_MASK"]} label="Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_DEFAULT_GW"]} label="Default Gateway">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","2","X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP"]} label="Destination Subnet">
                    <Input />
                  </Form.Item>
                </>
              )}
            </>
          )}
        </ProCard>
        <ProCard title="Extra WAN IP3">
          <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_CONFIG"]} label="Extra WAN IP3">
            <Select onChange={setExtraWanIp3Enable}>
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          {extraWanIp3Enable === "1" && (
            <>
              <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG"]} label="VLAN Tag Flag">
                <Input />
              </Form.Item>
              <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_IPADDR_METHOD"]} label="IP Address Tye">
                <Select placeholder="Select IP address type" onChange={setExtraWanIp3AddrType}>
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {extraWanIp3AddrType === "1" && (
                <>
                  <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_IPADDR_STATIC"]} label="IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_IPADDR_MASK"]} label="Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_DEFAULT_GW"]} label="Default Gateway">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","3","X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP"]} label="Destination Subnet">
                    <Input />
                  </Form.Item>
                </>
              )}
            </>
          )}
        </ProCard>

        <ProCard title="Extra WAN IP4">
          <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_CONFIG"]} label="Extra WAN IP4">
            <Select onChange={setExtraWanIp4Enable}>
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          {extraWanIp4Enable === "1" && (
            <>
              <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_VLAN1_TAG_FLAG"]} label="VLAN Tag Flag">
                <Input />
              </Form.Item>
              <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_IPADDR_METHOD"]} label="IP Address Tye">
                <Select placeholder="Select IP address type" onChange={setExtraWanIp4AddrType}>
                  <Select.Option key="0" value="0">DHCP</Select.Option>
                  <Select.Option key="1" value="1">Static</Select.Option>
                </Select>
              </Form.Item>
              {extraWanIp4AddrType === "1" && (
                <>
                  <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_IPADDR_STATIC"]} label="IP Address">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_IPADDR_MASK"]} label="Subnet Mask">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_DEFAULT_GW"]} label="Default Gateway">
                    <Input />
                  </Form.Item>
                  <Form.Item name={["EXTRA_IP_LIST","4","X_VENDOR_EXTRA_IP_ROUTE_TARGET_IP"]} label="Destination Subnet">
                    <Input />
                  </Form.Item>
                </>
              )}
            </>
          )}
        </ProCard>
      </ProCard>
    </ProForm>
  );
};

export default WANSettingForm;
