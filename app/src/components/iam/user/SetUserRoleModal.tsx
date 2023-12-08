import React, { useState, useEffect } from 'react';
import { Modal, Form } from 'antd';
import { setRolesForUser } from '@/models/iam_users';
import { RoleSelect } from '..';

interface SetUserRoleModalProps {
  visible: boolean;
  user: API.User | undefined;
  onCancel: () => void;
  onSuccess: () => void;
}

const SetUserRoleModal: React.FC<SetUserRoleModalProps> = ({
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
        await setRolesForUser(user, values);
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
          label="Roles"
          name="RoleNames"
        >
          <RoleSelect />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default SetUserRoleModal;
