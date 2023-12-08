import React, { useState } from 'react';
import { Modal, Form, Input, message } from 'antd';
import services from '@/services/netdoop';
const { createProduct } = services.omcProducts;

interface CreateProductFormProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
}

const CreateProductForm: React.FC<CreateProductFormProps> = ({
  visible,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      await createProduct(values);
      message.success('Device type created successfully');
      form.resetFields()
      onSuccess();
    } catch (error) {
      message.error('Failed to create product');
    } finally {
      setSubmitting(false);
    }
  };
  const handleCancel = async () => {
    form.resetFields()
    onCancel()
  };

  return (
    <Modal
      title="Create Product"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="OUI"
          name="Oui"
          rules={[{ required: true, message: 'Please input OUI' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Product Class"
          name="ProductClass"
          rules={[{ required: true, message: 'Please input product class' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Manufacturer"
          name="Manufacturer"
          rules={[{ required: true, message: 'Please input manufacturer' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Name"
          name="Name"
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateProductForm;
