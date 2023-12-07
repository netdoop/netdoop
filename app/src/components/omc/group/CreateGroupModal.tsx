import React, { useState } from 'react';
import { Modal, Form, Input } from 'antd';
import GroupTreeSelect from './GroupTreeSelect';
import { createGroup } from '@/models/groups';

interface CreateGroupModalProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  parent: API.Group | null;
}

const CreateGroupModal: React.FC<CreateGroupModalProps> = ({
  visible,
  onCancel,
  onSuccess,
  parent,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);
  if (parent){
    form.setFieldValue('parentId', parent.Id)
  }
  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      const values = await form.validateFields();
      await createGroup(values);
      form.resetFields();
      onSuccess();
    } finally {
      onCancel();
      setSubmitting(false);
      form.resetFields();
    }
  };

  const handleCancel = async () => {
    onCancel();
    form.resetFields();
  };

  return (
    <Modal
      title="Create Group"
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
          rules={[{ required: true, message: 'Please input group name' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Parent"
          name="ParentId"
          rules={[{ required: true, message: 'Please select parent group' }]}
        >
          <GroupTreeSelect />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateGroupModal;
