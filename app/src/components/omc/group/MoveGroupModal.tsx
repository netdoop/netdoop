import React, { useState } from 'react';
import { Modal, Form } from 'antd';
import GroupTreeSelect from './GroupTreeSelect';
import { setGroupParent } from '@/models/groups';

interface MoveGroupModalProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  group: API.Group | null;
}

const MoveGroupModal: React.FC<MoveGroupModalProps> = ({
  visible,
  onCancel,
  onSuccess,
  group,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      if (group?.Id) {
        await setGroupParent(group, { parentID: values.parentId });
        onSuccess();
      }
    } finally {
      onCancel();
      setSubmitting(false);
      form.resetFields();
    }
  };

  const handleCancel = async () => {
    form.resetFields();
    onCancel();
  };

  return (
    <Modal
      title={`Move Group: ${group?.Name}`}
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Parent"
          name="parentId"
          rules={[{ required: true, message: 'Please select a new parent group' }]}
        >
          <GroupTreeSelect hiddenGroupID={group?.Id} />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default MoveGroupModal;