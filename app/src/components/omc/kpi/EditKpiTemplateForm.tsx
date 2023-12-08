import React, { useEffect, useState } from 'react';
import { Col, Form, Input, Radio, Row, Select, Space } from 'antd';
import { createKPITemplate, updateKPITemplate, useKPITemplate } from '@/models/kpi_temp';
import { ProCard, ProForm } from '@ant-design/pro-components';
import { SelectKpiMeasInput } from '.';
import { SelectDeviceInput } from '../device';
import { SelectGroupInput } from '../group';

const layout = {
  labelCol: { span: 4 },
  wrapperCol: { span: 24 },
};

interface Props {
  id?: number,
}

const EditKpiTemplateForm: React.FC<Props> = ({ id }) => {
  const [form] = Form.useForm();
  const [selectType, setSelectType] = useState<string>('Device');

  const { data } = useKPITemplate(id);

  useEffect(() => {
    if (data) {
      form.setFieldsValue(data);
    } else {
      const data: API.KPITemplate = {
        PeriodicInterval: 900,
        SelectType: 'Device',
        SelectIds: [],
      }
      form.setFieldsValue(data);
    }
  }, [data]);

  const onFinish = async (data2: API.KPITemplate) => {
    try {
      if (data) {
        await updateKPITemplate(data, {
          periodicInterval: data2.PeriodicInterval,
          selectType: data2.SelectType,
          selectIds: data2.SelectIds,
          measTypeIds: data2.MeasTypeIds
        })
      } else {
        await createKPITemplate({
          name: data2.Name,
          periodicInterval: data2.PeriodicInterval,
          selectType: data2.SelectType,
          selectIds: data2.SelectIds,
          measTypeIds: data2.MeasTypeIds
        });
      }
    } finally {
      const data: API.KPITemplate = {
        SelectType: 'Device',
        SelectIds: [],
      }
      form.setFieldsValue(data);
      history.back()
    }
  };

  const onReset = () => {
    if (data) {
      form.setFieldsValue(data);
    } else {
      const data: API.KPITemplate = {
        SelectType: 'Device',
        SelectIds: [],
      }
      form.setFieldsValue(data);
    }
  };

  return (
    <ProCard
      title={data ? "Edit KPI Template: " + data.Name : "Create KPI Template"}
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
          <Form.Item
            label="Name"
            name="Name"
            rules={[{ required: true, message: 'Please input template name' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            label="Periodic Interval"
            name="PeriodicInterval"
          >
            <Select placeholder="Select peridoci interval">
              <Select.Option key={900} value={900}>15m</Select.Option>
              <Select.Option key={1800} value={1800}>30m</Select.Option>
              <Select.Option key={3600} value={3600}>60m</Select.Option>
            </Select>
          </Form.Item>

          <Form.Item name="SelectType" label="Select Type">
            <Radio.Group onChange={({ target }) => { setSelectType(target.value); }}>
              <Radio value="Device">Device</Radio>
              <Radio value="DeviceGroup">DeviceGroup</Radio>
            </Radio.Group>
          </Form.Item>

          {selectType === 'Device' ? (
            <Form.Item name="SelectIds" label="Devices">
              <SelectDeviceInput productType="enb" />
            </Form.Item>
          ) : (
            <Form.Item name="SelectIds" label="Groups">
              <SelectGroupInput />
            </Form.Item>
          )}

          <Form.Item name="MeasTypeIds" label="KPI Measures">
            <SelectKpiMeasInput />
          </Form.Item>
        </ProForm>
      </ProCard>
    </ProCard>
  );
};

export default EditKpiTemplateForm;
