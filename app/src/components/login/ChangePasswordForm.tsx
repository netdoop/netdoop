import React, { useState } from 'react';
import { Modal, Form, Input } from 'antd';
import { changeCurrentPassword } from '@/models/iam_users';

interface ChangePasswordFormProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
}

const ChangePasswordForm: React.FC<ChangePasswordFormProps> = ({
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
      await changeCurrentPassword(values);
      onSuccess();
    } finally {
      onCancel()
      setSubmitting(false);
      form.resetFields();
    }
  };
  const handleCancel = async () => {
    onCancel()
    form.resetFields();
  }

  return (
    <Modal
      title="Change Password"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Current Password"
          name="Password"
          rules={[
            { required: true, message: 'Please input current password' },
          ]}
        >
          <Input.Password />
        </Form.Item>
        <Form.Item
          label="New Password"
          name="NewPassword"
          rules={[
            { required: true, message: 'Please input new password' },
            { min: 6, message: 'Password must be at least 6 characters' },
          ]}
        >
          <Input.Password />
        </Form.Item>
        <Form.Item
          label="Confirm New Password"
          name="ConfirmNewPassword"
          dependencies={['NewPassword']}
          rules={[
            { required: true, message: 'Please confirm new password' },
            ({ getFieldValue }) => ({
              validator(_, value) {
                if (!value || getFieldValue('NewPassword') === value) {
                  return Promise.resolve();
                }
                return Promise.reject('The two passwords do not match');
              },
            }),
          ]}
        >
          <Input.Password />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ChangePasswordForm;
