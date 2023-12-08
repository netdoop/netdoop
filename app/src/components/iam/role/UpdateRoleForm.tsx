import React, { useState, useEffect } from 'react';
import { Modal, Form, Input } from 'antd';
import { updateRole } from '@/models/iam_roles';

interface UpdateRoleFormProps {
  visible: boolean;
  role: API.Role| undefined;
  onCancel: () => void;
  onSuccess: () => void;
}

const UpdateRoleForm: React.FC<UpdateRoleFormProps> = ({
  visible,
  role,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    if (visible) {
      form.setFieldsValue(role);
    }
  }, [visible, form, role]);

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      if (role?.Id){
        await updateRole(role, values);
        onSuccess();
      }
    } finally {
      setSubmitting(false);
      onCancel();
      form.resetFields();
    }
  };

  const handleCancel = async () => {
    onCancel();
    form.resetFields();
  }

  return (
    <Modal
      title="Update Role"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <Form
        form={form}
        layout="vertical"
      >
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

export default UpdateRoleForm;