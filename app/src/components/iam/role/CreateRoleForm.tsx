import React, { useState } from 'react';
import { Modal, Form, Input } from 'antd';
import { createRole } from '@/models/iam_roles';

interface CreateRoleFormProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
}

const CreateRoleForm: React.FC<CreateRoleFormProps> = ({
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
      await createRole(values);
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
      title="Create Role"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      forceRender
    >
      <Form form={form} layout="vertical" >
        <Form.Item
          label="Name"
          name="Name"
          rules={[{ required: true, message: 'Please input role name' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Alias"
          name="Alias"
          rules={[{ required: true, message: 'Please input role alias' }]}
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateRoleForm;
