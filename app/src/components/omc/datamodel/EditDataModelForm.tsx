import React, { useEffect } from 'react';
import { Col, Form, Input, Row, Select, Space } from 'antd';
import { createDataModel, useDataModel } from '@/models/datamodel';
import { ProCard, ProForm } from '@ant-design/pro-components';

const layout = {
  labelCol: { span: 4 },
  wrapperCol: { span: 24 },
};

interface Props {
  id?: number,
}

const EditDataModelForm: React.FC<Props> = ({ id }) => {
  const [form] = Form.useForm();
  const { data } = useDataModel(id);

  useEffect(() => {
    if (data) {
      form.setFieldsValue(data);
    } else {
      const data: API.DataModel = {
        ProductType: 'enb',
        ParameterPath: 'Device.',
      }
      form.setFieldsValue(data);
    }
  }, [data]);

  const onFinish = async (data2: API.DataModel) => {
    try {
      if (data) {

      } else {
        await createDataModel({
          name: data2.Name,
          productType: data2.ProductType,
          parameterPath: data2.ParameterPath,
        });
      }
    } finally {
      const data: API.DataModel = {
        ProductType: 'enb',
        ParameterPath: 'Device.',
      }
      form.setFieldsValue(data);
      history.back()
    }
  };

  const onReset = () => {
    if (data) {
      form.setFieldsValue(data);
    } else {
      const data: API.DataModel = {
        ProductType: 'enb',
        ParameterPath: 'Device.',
      }
      form.setFieldsValue(data);
    }
  };

  return (
    <ProCard
      title={data ? "Edit Data Model: " + data.Name : "Create Data Model"}
      gutter={8}
      style={{ marginBlockStart: 8 }}>
      <ProCard colSpan='90%' direction="column">
        <ProForm
          {...layout}
          form={form}
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
          <Form.Item label="Name" name="Name"
            rules={[{ required: true, message: 'Please input template name' }]}>
            <Input />
          </Form.Item>
          <Form.Item label="Product Type" name="ProductType"
            rules={[{ required: true, message: 'Please select product type' }]}>
            <Select>
              <Select.Option key={'enb'} value={'enb'}>ENB</Select.Option>
              <Select.Option key={'cpe'} value={'cpe'}>CPE</Select.Option>
            </Select>
          </Form.Item>

          <Form.Item name="ParameterPath" label="Parameter Path"
            rules={[{ required: true, message: 'Please select parameter path' }]}>
            <Select>
              <Select.Option key={'Device.'} value={'Device.'}>Device.</Select.Option>
              <Select.Option key={'InternatGatewayDevice.'} value={'InternatGatewayDevice.'}>InternatGatewayDevice.</Select.Option>
            </Select>
          </Form.Item>

        </ProForm>
      </ProCard>
    </ProCard>
  );
};

export default EditDataModelForm;
