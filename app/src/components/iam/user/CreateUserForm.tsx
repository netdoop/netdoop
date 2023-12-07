import React, { useState } from 'react';
import { Modal, Form, Input } from 'antd';
import { createUser } from '@/models/iam_users';

interface CreateUserFormProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
}

const CreateUserForm: React.FC<CreateUserFormProps> = ({
  visible,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      const values = await form.validateFields();
      await createUser(values);
      onSuccess();
    } finally {
      setSubmitting(false);
      onCancel();
      form.resetFields();
    }
  };
  const handleCancel = async () => {
    onCancel();
    form.resetFields();
  };

  return (
    <Modal
      title="Create User"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Name"
          name="Name"
          rules={[{ required: true, message: 'Please input user name' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Alias"
          name="Alias"
          rules={[{ required: true, message: 'Please input user alias' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Password"
          name="Password"
          rules={[{ required: true, message: 'Please input user password' }]}
        >
          <Input.Password />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateUserForm;
