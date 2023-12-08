import React, { useState, useEffect } from 'react';
import { Modal, Form, Input } from 'antd';
import { updateDeviceInfo } from '@/models/device';

interface UpdateDeviceInfoModalProps {
  visible: boolean;
  device?: API.Device;
  onCancel: () => void;
  onSuccess: () => void;
}

const UpdateDeviceInfoModal: React.FC<UpdateDeviceInfoModalProps> = ({
  visible,
  device,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    if (visible) {
      form.setFieldsValue(device);
    }
  }, [visible, form, device]);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (device) {
        const values = await form.validateFields();
        await updateDeviceInfo(device, values);

        onSuccess();

      }
    } finally {
      onCancel();
      setSubmitting(false);
      form.resetFields();
    }
  };

  const handleCancel = async () => {
    form.resetFields()
    onCancel()
  }

  return (
    <Modal
      title="Update Device Information"
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
          rules={[{ required: true, message: 'Please input device alias' }]}
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default UpdateDeviceInfoModal;
