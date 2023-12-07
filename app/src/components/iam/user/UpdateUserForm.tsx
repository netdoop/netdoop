import React, { useState, useEffect } from 'react';
import { Modal, Form, Input } from 'antd';
import { updateUser } from '@/models/iam_users';

interface UpdateUserFormProps {
  visible: boolean;
  user: API.User | undefined;
  onCancel: () => void;
  onSuccess: () => void;
}

const UpdateUserForm: React.FC<UpdateUserFormProps> = ({
  visible,
  user,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    if (visible) {
      form.setFieldsValue(user);
    }
  }, [visible, form, user]);

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      if (user?.Id){
        await updateUser(user, values);
        onSuccess();
      }
    } finally {
      setSubmitting(false);
      onCancel()
      form.resetFields()
    }
  };

  const handleCancel = async () => {
    onCancel()
    form.resetFields()
  }

  return (
    <Modal
      title="Update User"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Alias"
          name="Alias"
          rules={[{ required: true, message: 'Please input user alias' }]}
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default UpdateUserForm;
