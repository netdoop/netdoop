import { Button, Form, Input } from "antd";
import { ProCard, ProForm, ProTable } from '@ant-design/pro-components';
import { useEffect, useState } from "react";

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};

interface SIMLockFormProps {
  device: API.Device,
  data?: Record<string, any>,
};

const SIMLockForm: React.FC<SIMLockFormProps> = ({ data }) => {
  const [form] = Form.useForm();
  const [dataSource, setDataSource] = useState<{
    index: number;
    plmn: string;
  }[]>([])

  const onFinish = async (values: Record<string, any>) => {
    console.log('Form values:', values);
    // Submit the form data to the server
  };
  const onReset = () => {
    const simLock:Record<string, any> = data?.SIMLock || {};
    form.setFieldsValue(simLock);
  };
  useEffect(() => {
    const simLock:Record<string, any> = data?.SIMLock || {};
    form.setFieldsValue(simLock);
  }, [data]);

  useEffect(() => {
    const plmnText = data?.PLMN || ''
    if (plmnText === ''){
      setDataSource([])
    }else{
      const plmnArray:string[] = plmnText.split(";");
      const plmnListDataSource = plmnArray.map((plmn, index) => {
        let numericValue = plmn;
        if (plmn.startsWith("plmn_")) {
          numericValue = plmn.substring(5);
        }
        return {
          index: index + 1,
          plmn: numericValue,
        };
      });
      setDataSource(plmnListDataSource)
    }
  }, [data])

  const plmnListColumns = [
    { title: "Index", dataIndex: "index", key: "index" },
    { title: "PLMN", dataIndex: "plmn", key: "plmn" },
  ]

  return (
    <ProCard split="horizontal">
      <ProCard title="Settings">
        <ProForm
          {...layout}
          form={form}
          layout="horizontal"
          onFinish={onFinish}
          onReset={onReset}
          submitter={false}
          labelWrap
        >
          <Form.Item name="PLMN" label="PLMN">
            <Input />
          </Form.Item>
          <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
            <Button htmlType="submit">Add</Button>
          </Form.Item>
        </ProForm>
      </ProCard>
      <ProCard title="PLMN List">
        <ProTable dataSource={dataSource} columns={plmnListColumns} pagination={false} search={false} options={false} />;
      </ProCard>
    </ProCard>
  );
};

export default SIMLockForm;
