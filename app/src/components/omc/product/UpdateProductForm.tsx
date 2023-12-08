import React, { useState, useEffect } from 'react';
import { Modal, Form, Input, message } from 'antd';
import services from '@/services/netdoop';
const { updateProductInfo } = services.omcProducts;

interface UpdateProductFormProps {
  visible: boolean;
  product: API.Product | null;
  onCancel: () => void;
  onSuccess: () => void;
}

const UpdateProductForm: React.FC<UpdateProductFormProps> = ({
  visible,
  product,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    if (visible) {
      form.setFieldsValue(product);
    }
  }, [visible, form, product]);

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      if (product?.Id){
        await updateProductInfo({ id: product.Id }, values);
        message.success('Device type updated successfully');
      }
      form.resetFields()
      onSuccess();
    } catch (error) {
      message.error('Failed to update product');
    } finally {
      setSubmitting(false);
    }
  };

  const handleCancel = async () => {
    form.resetFields()
    onCancel()
  }

  return (
    <Modal
      title="Update Product"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Name"
          name="Name"
          rules={[{ required: true, message: 'Please input product alias' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Manufacturer"
          name="Manufacturer"
          rules={[{ required: true, message: 'Please input product manufacturer' }]}
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default UpdateProductForm;
