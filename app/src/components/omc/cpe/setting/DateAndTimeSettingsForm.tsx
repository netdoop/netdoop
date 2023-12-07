import { Col, DatePicker, Form, Input, Row, Select, Space } from "antd";
import { ProCard, ProForm } from '@ant-design/pro-components';
import { setDeviceParameterValues } from "@/models/device";
import { useEffect, useState } from "react";
import dayjs from "dayjs";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

const setTypeOptions = [
  { label: 'Set Manually', value: 'Manual' },
  { label: 'Sync from network', value: 'Sync' },
];

const timezoneOptions = [
  { label: 'GMT+12:00', value: 'GMT+12:00' },
  { label: 'GMT+11:00', value: 'GMT+11:00' },
  { label: 'GMT+10:00', value: 'GMT+10:00' },
  { label: 'GMT+09:00', value: 'GMT+09:00' },
  { label: 'GMT+08:00', value: 'GMT+08:00' },
  { label: 'GMT+07:00', value: 'GMT+07:00' },
  { label: 'GMT+06:00', value: 'GMT+06:00' },
  { label: 'GMT+05:00', value: 'GMT+05:00' },
  { label: 'GMT+04:00', value: 'GMT+04:00' },
  { label: 'GMT+03:30', value: 'GMT+03:30' },
  { label: 'GMT+03:00', value: 'GMT+03:00' },
  { label: 'GMT+02:00', value: 'GMT+02:00' },
  { label: 'GMT+01:00', value: 'GMT+01:00' },
  { label: 'GMT-12:00', value: 'GMT-12:00' },
  { label: 'GMT-11:00', value: 'GMT-11:00' },
  { label: 'GMT-10:00', value: 'GMT-10:00' },
  { label: 'GMT-09:00', value: 'GMT-09:00' },
  { label: 'GMT-08:00', value: 'GMT-08:00' },
  { label: 'GMT-07:00', value: 'GMT-07:00' },
  { label: 'GMT-06:00', value: 'GMT-06:00' },
  { label: 'GMT-05:00', value: 'GMT-05:00' },
  { label: 'GMT-04:00', value: 'GMT-04:00' },
  { label: 'GMT-03:30', value: 'GMT-03:30' },
  { label: 'GMT-03:00', value: 'GMT-03:00' },
  { label: 'GMT-02:00', value: 'GMT-02:00' },
  { label: 'GMT-01:00', value: 'GMT-01:00' },
];

const monthOptions = [
  { label: 'Jan', value: '1' },
  { label: 'Feb', value: '2' },
  { label: 'Mar', value: '3' },
  { label: 'Apr', value: '4' },
  { label: 'May', value: '5' },
  { label: 'Jun', value: '6' },
  { label: 'Jul', value: '7' },
  { label: 'Aug', value: '8' },
  { label: 'Sep', value: '9' },
  { label: 'Oct', value: '10' },
  { label: 'Nov', value: '11' },
  { label: 'Dec', value: '12' },
];

const weekOptions = [
  { label: 'First', value: '1' },
  { label: 'Second', value: '2' },
  { label: 'Third', value: '3' },
  { label: 'Fourth', value: '4' },
  { label: 'Last', value: '5' },
];

const weekdayOptions = [
  { label: 'Sun', value: '0' },
  { label: 'Mon', value: '1' },
  { label: 'Tue', value: '2' },
  { label: 'Wed', value: '3' },
  { label: 'Thu', value: '4' },
  { label: 'Fri', value: '5' },
  { label: 'Sat', value: '6' },
];

const hourOptions: { label: string, value: string }[] = [];
for (let i = 0; i < 24; i++) {
  const label = i < 10 ? `0${i}` : `${i}`; // add a leading 0 for single-digit minutes
  const value = i.toString();
  hourOptions.push({ label, value });
}

const minuteOptions: { label: string, value: string }[] = [];
for (let i = 0; i < 60; i++) {
  const label = i < 10 ? `0${i}` : `${i}`; // add a leading 0 for single-digit minutes
  const value = i.toString();
  minuteOptions.push({ label, value });
}

interface DateAndTimeSettingsFormProps {
  device: API.Device,
  data: Record<string, any>,
};

const DateAndTimeSettingsForm: React.FC<DateAndTimeSettingsFormProps> = ({
  device,
  data,
}) => {
  const [form] = Form.useForm();
  const [selectedType, setSelectedType] = useState<string>("Sync")

  const onFinish = async (data2: Record<string, any> & {
    PrimaryNTPServer: string,
    SecondaryNTPServer: string,
    LocalTime: string,
    DSTStartTimeParts: string[],
    DSTEndTimeParts: string[],
  }) => {
    let values: Record<string, any> = {}
    const prefixName = 'Device.WebGui.System.DateAndTime';
    if (selectedType === "Manual") {
      values[prefixName + 'CurrentTime'] = data2.LocalTime;
    }
    if (selectedType === "Sync") {
      values[prefixName + 'NTPServer'] = data2.PrimaryNTPServer + "," + data2.SecondaryNTPServer;
      values[prefixName + 'OptionalNTPServer'] = data2.OptionalNTPServer;
      values[prefixName + 'TimeZone'] = data2.TimeZone;
    }

    values[prefixName + 'DST.Enable'] = data2.DST?.Enable;
    if (data2.DST?.Enable) {
      values[prefixName + 'DST.StartTime'] = data2.DSTStartTimeParts.join("-");
      values[prefixName + 'DST.EndTime'] = data2.DSTEndTimeParts.join("-");
    }

    setDeviceParameterValues(device, values)
  };

  const updateForm = () => {
    const dateAndTime:Record<string, any> = data?.WEB_GUI?.System?.DateAndTime || {};
    const dst:Record<string, any> = data?.WEB_GUI?.System?.DateAndTime.DST || {};
    
    const NTPServer = dateAndTime?.NTPServer || "";
    const parts = NTPServer ? NTPServer.split(",") : ["", ""];
    const parts2 = dst?.StartTime ? dst?.StartTime.split("-") : ["", "", "", "", ""];
    const parts3 = dst?.EndTime ? dst?.EndTime.split("-") : ["", "", "", "", ""];

    form.setFieldsValue({
      ...dateAndTime,
      SetType: selectedType,
      PrimaryNTPServer: parts[0],
      SecondaryNTPServer: parts[1],
      DSTStartTimeParts: parts2,
      DSTEndTimeParts: parts3,
    });
  }
  const onReset = () => {
    updateForm();
    setSelectedType("Sync");
  };
  useEffect(() => {
    updateForm();
    setSelectedType("Sync");
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
          <Form.Item label="Current Time" name="CurrentTime">
            <Input readOnly />
          </Form.Item>
          <Form.Item key="SetType" name="SetType" label="Set Type">
            <Select placeholder="Select Set type" value={selectedType} onChange={(v) => { setSelectedType(v) }}>
              {setTypeOptions.map((item) => (
                <Select.Option key={item.value} value={item.value}>
                  {item.label}
                </Select.Option>
              ))}
            </Select>
          </Form.Item>
          {selectedType === "Sync" && (
            <>
              <Form.Item label="Primary NTP Server" name="PrimaryNTPServer">
                <Input />
              </Form.Item>
              <Form.Item label="Sencondary NTP Server" name="SecondaryNTPServer">
                <Input />
              </Form.Item>
              <Form.Item label="Optional NTP Server" name="OptionalNTPServer">
                <Input />
              </Form.Item>
              <Form.Item key="TimeZone" name="TimeZone" label="Time Zone">
                <Select placeholder="Select Time Zone">
                  {timezoneOptions.map((item) => (
                    <Select.Option key={item.value} value={item.value}>
                      {item.label}
                    </Select.Option>
                  ))}
                </Select>
              </Form.Item>
            </>
          )}
          {selectedType === "Manual" && (
            <>
              <Form.Item label="Local Time" name="LocalTime">
                <DatePicker
                  format="YYYY-MM-DD HH:mm:ss"
                  showTime={{ defaultValue: dayjs('00:00:00', 'HH:mm:ss') }}
                />
              </Form.Item>
            </>
          )}
        </ProCard>
        <ProCard title="DST">
          <Form.Item name={"DST.Enable"} label="DST" >
            <Select >
              <Select.Option key="0" value="0">Disable</Select.Option>
              <Select.Option key="1" value="1">Enable</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label="Start Time">
            <Form.Item name={["DSTStartTimeParts", 0]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Month">
                {monthOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTStartTimeParts", 1]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Week">
                {weekOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTStartTimeParts", 2]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Weekday">
                {weekdayOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTStartTimeParts", 3]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Hour">
                {hourOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTStartTimeParts", 4]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Minute">
                {minuteOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
          </Form.Item>
          <Form.Item label="End Time">
            <Form.Item name={["DSTEndTimeParts", 0]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Month">
                {monthOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTEndTimeParts", 1]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Week">
                {weekOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTEndTimeParts", 2]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Weekday">
                {weekdayOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTEndTimeParts", 3]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Hour">
                {hourOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
            <Form.Item name={["DSTEndTimeParts", 4]} style={{ display: 'inline-block', width: 'calc(20% - 8px)' }}>
              <Select placeholder="Select Minute">
                {minuteOptions.map((item) => (<Select.Option key={item.value} value={item.value}>{item.label}</Select.Option>))}
              </Select>
            </Form.Item>
          </Form.Item>
          <Form.Item label="Status" name={"DST.Status"}>
            <Input readOnly />
          </Form.Item>
        </ProCard>
      </ProCard>
    </ProForm >
  );
};

export default DateAndTimeSettingsForm;
