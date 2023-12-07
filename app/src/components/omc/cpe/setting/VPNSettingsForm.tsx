import { Col, Form, Input, InputNumber, Row, Select, Space } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const vpnProtocolOptions = [
  { label: 'PPTP', value: 'pptp' },
  { label: 'L2TP', value: 'l2tp' },
  { label: 'L2TPv3', value: 'l2tpv3' },
  { label: 'GRE', value: 'gre' },
];

const bearDeviceOptions = [
  { label: 'APN1', value: 'APN1' },
  { label: 'APN2', value: 'APN2' },
  { label: 'APN3', value: 'APN3' },
  { label: 'APN4', value: 'APN4' },
];

const vpnLayerOptions = [
  { label: 'Layer2', value: "2" },
  { label: 'Layer3', value: "3" },
];

interface VPNSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

interface vpnStatusItem {
  key: string,
  Index?: number,
  LocalAddress?: string,
  RemoteAddress?: string,
  OnlineTime?: string,
  Username?: string,
}

const VPNSettingsForm: React.FC<VPNSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [vpnStatusItems, setVpnStatusItems] = useState<vpnStatusItem[]>([]);
  const [selectedProtocol, setSelectedProtocol] = useState<string>("pptp")
  const prefixName = 'Device.WEB_GUI.VPN.'

  const onFinish = async (data2: Record<string, any>) => {
    let values: Record<string, any> = {}
    values[prefixName + "Enable"] = data2.Enable
    values[prefixName + "Protocol"] = data2.Protocol

    if (data2.Protocol === "pptp" || data2.Protocol === "l2tp") {
      values[prefixName + "Server"] = data2.Server
      values[prefixName + "Username"] = data2.Username
      values[prefixName + "Password"] = data2.Password
    }
    if (data2.Protocol !== "pptp") {
      values[prefixName + "IPsecEnable"] = data2.IPsecEnable
      values[prefixName + "IPsecPassword"] = data2.IPsecPassword
    }
    if (data2.Protocol === "l2tpv3") {
      values[prefixName + ".L2TPv3.LocalCookie"] = data2["L2TPv3.LocalCookie"];
      values[prefixName + ".L2TPv3.LocalSessionID"] = data2["L2TPv3.LocalSessionID"];
      values[prefixName + ".L2TPv3.LocalTunnelID"] = data2["L2TPv3.LocalTunnelID"];
      values[prefixName + ".L2TPv3.LocalTunnelIPAddress"] = data2["L2TPv3.LocalTunnelIPAddress"];
      values[prefixName + ".L2TPv3.RemoteCookie"] = data2["L2TPv3.RemoteCookie"];
      values[prefixName + ".L2TPv3.RemoteIPAddress"] = data2["L2TPv3.RemoteIPAddress"];
      values[prefixName + ".L2TPv3.RemoteSessionID"] = data2["L2TPv3.RemoteSessionID"];
      values[prefixName + ".L2TPv3.RemoteTunnelID"] = data2["L2TPv3.RemoteTunnelID"];
      values[prefixName + ".L2TPv3.RemoteTunnelIPAddress"] = data2["L2TPv3.RemoteTunnelIPAddress"];
      values[prefixName + ".L2TPv3.UDPDestinationPort"] = data2["L2TPv3.UDPDestinationPort"];
      values[prefixName + ".L2TPv3.UDPSourcePort"] = data2["L2TPv3.UDPSourcePort"];
      values[prefixName + ".L2TPv3.VLANID"] = data2["L2TPv3.VLANID"];
      values[prefixName + ".L2TPv3.VLANIDEnable"] = data2["L2TPv3.VLANIDEnable"];
    } else if (data2.Protocol === "gre") {
      values[prefixName + ".GRE.GREDestinationAddress"] = data2["GRE.GREDestinationAddress"];
      values[prefixName + ".GRE.HostIPAddress"] = data2["GRE.HostIPAddress"];
      values[prefixName + ".GRE.RemoteIPAddress"] = data2["GRE.RemoteIPAddress"];
      values[prefixName + ".GRE.RemotePrivateIPAddress"] = data2["GRE.RemotePrivateIPAddress"];
      values[prefixName + ".GRE.RemotePrivateIPAddressPrefix"] = data2["GRE.RemotePrivateIPAddressPrefix"];
      values[prefixName + ".GRE.VPNLayer"] = data2["GRE.VPNLayer"];
    }
    setDeviceParameterValues(device, values)
  };
  const onReset = () => {
    const vpn:Record<string, any> = data?.WEB_GUI?.VPN || {};
    form.setFieldsValue(vpn);
  };
  useEffect(() => {
    const vpn:Record<string, any> = data?.WEB_GUI?.VPN || {};
    const statusList:Record<string, Record<string, any>> = data?.WEB_GUI?.VPN?.Status?.List || {};
    setSelectedProtocol(vpn?.Protocol || "pptp")

    let items: vpnStatusItem[] = []
      Object.entries(statusList).forEach(([k, v]) => {
        items.push({
          key: k,
          Index: v.Index,
          LocalAddress: v.LocalAddress || "-",
          RemoteAddress: v.RemoteAddress || "-",
          OnlineTime: v.OnlineTime || "-",
          Username: v.Username || "-",
        })
      });
    setVpnStatusItems(items)
    form.setFieldsValue(vpn);
  }, [data]);

  const vpnStatusColumns = [
    { title: "#", dataIndex: "key", key: "key" },
    { title: "Username", dataIndex: "Username", key: "Username" },
    { title: "LocalAddress", dataIndex: "LocalAddress", key: "LocalAddress" },
    { title: "RemoteAddress", dataIndex: "RemoteAddress", key: "RemoteAddress" },
    { title: "OnlineTime", dataIndex: "OnlineTime", key: "OnlineTime" },
  ]

  return (
    <ProCard split="horizontal">
      <ProCard title="VPN Settings">
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
          <Form.Item key="Enable" name="Enable" label="VPN">
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="Protocol" label="Protocol">
            <Select placeholder="Select Protocol" onChange={(v) => { setSelectedProtocol(v) }}>
              {vpnProtocolOptions.map((option) => (
                <Select.Option key={option.value} value={option.value}>
                  {option.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          {(selectedProtocol === "l2tp" || selectedProtocol === "pptp") && (
            <>
              <Form.Item key="Server" name="Server" label="VPN Server">
                <Input placeholder="Enter Server" />
              </Form.Item>
              <Form.Item key="Username" name="Username" label="Username">
                <Input placeholder="Enter Username" />
              </Form.Item>
              <Form.Item key="Password" name="Password" label="Password">
                <Input.Password placeholder="Enter Password" />
              </Form.Item>
            </>
          )}
          {selectedProtocol === "pptp" && (
            <Form.Item key="BearDevice" name="BearDevice" label="Bear Device">
              <Select placeholder="Select Bear Device">
                {bearDeviceOptions.map((option) => (
                  <Select.Option key={option.value} value={option.value}>
                    {option.label}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          )}
          {selectedProtocol === "l2tpv3" && (
            <>
              <Form.Item key="EncapsulationType" name={["L2APv3.EncapsulationType"]} label="Encapsulation Type">
                <Input placeholder="Enter Encapsulation Type" />
              </Form.Item>
              <Form.Item key="VLANIDEnable" name={["L2APv3.VLANIDEnable"]} label="VLAN ID Enable">
                <Select >
                  <Select.Option key="0" value="0">Disable</Select.Option>
                  <Select.Option key="1" value="1">Enable</Select.Option>
                </Select>
              </Form.Item>
              <Form.Item key="VLANID" name={["L2APv3.VLANID"]} label="VLAN ID">
                <Input placeholder="Enter VLAN ID" />
              </Form.Item>
              <Form.Item key="LocalCookie" name={["L2APv3.LocalCookie"]} label="Local Cookie">
                <Input placeholder="Enter Local Cookie" />
              </Form.Item>
              <Form.Item key="RemoteCookie" name={["L2APv3.RemoteCookie"]} label="Remote Cookie">
                <Input placeholder="Enter Remote Cookie" />
              </Form.Item>
              <Form.Item key="LocalTunnelID" name={["L2APv3.LocalTunnelID"]} label="Local Tunnel ID">
                <Input placeholder="Enter Local Tunnel ID" />
              </Form.Item>
              <Form.Item key="RemoteTunnelID" name={["L2APv3.RemoteTunnelID"]} label="Remote Tunnel ID">
                <Input placeholder="Enter Remote Tunnel ID" />
              </Form.Item>
              <Form.Item key="RemoteIPAddress" name={["L2APv3.RemoteIPAddress"]} label="Remote IP Address">
                <Input placeholder="Enter Remote IP Address" />
              </Form.Item>
              <Form.Item key="UDPSourcePort" name={["L2APv3.UDPSourcePort"]} label="UDP Source Port">
                <InputNumber placeholder="Enter UDP Source Port" min={0} />
              </Form.Item>
              <Form.Item key="UDPDestinationPort" name={["L2APv3.UDPDestinationPort"]} label="UDP Destination Port">
                <InputNumber placeholder="Enter UDP Destination Port" min={0} />
              </Form.Item>

              <Form.Item key="LocalSessionID" name={["L2APv3.LocalSessionID"]} label="Local Session ID">
                <Input placeholder="Enter Local Session ID" />
              </Form.Item>
              <Form.Item key="RemoteSessionID" name={["L2APv3.RemoteSessionID"]} label="Remote Session ID">
                <Input placeholder="Enter Remote Session ID" />
              </Form.Item>
              <Form.Item key="LocalTunnelIPAddress" name={["L2APv3.LocalTunnelIPAddress"]} label="Local Tunnel IP Address">
                <Input placeholder="Enter Local Tunnel IP Address" />
              </Form.Item>
              <Form.Item key="RemoteTunnelIPAddress" name={["L2APv3.RemoteTunnelIPAddress"]} label="Remote Tunnel IP Address">
                <Input placeholder="Enter Remote Tunnel IP Address" />
              </Form.Item>
            </>
          )}
          {selectedProtocol === "gre" && (
            <>
              <Form.Item key="GREDestinationAddress" name={["GRE.GREDestinationAddress"]} label="GRE Destination Address">
                <Input placeholder="Enter GRE Destination Address" />
              </Form.Item>
              <Form.Item key="HostIPAddress" name={["GRE.HostIPAddress"]} label="Host IP Address">
                <Input placeholder="Enter Host IP Address" />
              </Form.Item>
              <Form.Item key="RemoteIPAddress" name={["GRE.RemoteIPAddress"]} label="Remote IP Address">
                <Input placeholder="Enter Remote IP Address" />
              </Form.Item>
              <Form.Item key="RemotePrivateIPAddress" name={["GRE.RemotePrivateIPAddress"]} label="Remote Private IP Address">
                <Input placeholder="Enter Remote Private IP Address" />
              </Form.Item>
              <Form.Item key="RemotePrivateIPAddressPrefix" name={["GRE.RemotePrivateIPAddressPrefix"]} label="Remote Private IP Address Prefix">
                <Input placeholder="Enter Remote Private IP Address Prefix" />
              </Form.Item>
              <Form.Item key="VPNLayer" name={["GRE.VPNLayer"]} label="VPN Layer">
                <Select placeholder="Select VPN Layer">
                  {vpnLayerOptions.map((option) => (
                    <Select.Option key={option.value} value={option.value}>
                      {option.label}
                    </Select.Option>
                  ))}
                </Select>
              </Form.Item>
            </>
          )}
          {selectedProtocol !== "pptp" && (
            <>
              <Form.Item key="IPsecEnable" name="IPsecEnable" label="IPsec Enable">
                <Select >
                  <Select.Option key="0" value="0">Disable</Select.Option>
                  <Select.Option key="1" value="1">Enable</Select.Option>
                </Select>
              </Form.Item>
              <Form.Item key="IPsecPassword" name="IPsecPassword" label="IPsec Password">
                <Input.Password placeholder="Enter IPsec Password" />
              </Form.Item>
            </>
          )}
          <Form.Item key="DefaultRouteEnable" name="DefaultRouteEnable" label="Default Route">
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
        </ProForm>
      </ProCard>
      <ProCard title="VPN Status List">
        <ProTable dataSource={vpnStatusItems} columns={vpnStatusColumns} pagination={false} search={false} options={false} />;
      </ProCard>
    </ProCard>
  );
};

export default VPNSettingsForm;
