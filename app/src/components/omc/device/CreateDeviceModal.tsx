import React, { useState } from 'react';
import { Modal, Form, Input } from 'antd';
import { createDevice } from '@/models/device';

interface CreateDeviceModalProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
}

const CreateDeviceModal: React.FC<CreateDeviceModalProps> = ({
  visible,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true)
      const values = await form.validateFields();
      await createDevice(values);
      onSuccess();
    } finally{
      onCancel();
      setSubmitting(false);
      form.resetFields();  
    }
  };

  const handleCancel = () => {
    onCancel();
    form.resetFields();
  };

  return (
    <Modal
      title="Create Device"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Serial Number"
          name="SerialNumber"
          rules={[{ required: true, message: 'Please input serial number' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Model"
          name="Model"
          rules={[{ required: true, message: 'Please input model' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Firmware Version"
          name="FirmwareVersion"
          rules={[{ required: true, message: 'Please input firmware version' }]}
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

export default CreateDeviceModal;