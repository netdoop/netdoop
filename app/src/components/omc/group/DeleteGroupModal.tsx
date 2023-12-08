import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteGroup } from '@/models/groups';

interface DeleteGroupModalProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  group: API.Group | null;
}

const DeleteGroupModal: React.FC<DeleteGroupModalProps> = ({
  visible,
  onCancel,
  onSuccess,
  group,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (group?.Id) {
        await deleteGroup(group);
        onSuccess();
      }
    } finally {
      onCancel()
      setSubmitting(false);
    }
  };

  const handleCancel = async () => {
    onCancel();
  };

  return (
    <Modal
      title="Delete Group"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
    >
      <p>Are you sure you want to delete this group?</p>
    </Modal>
  );
};

export default DeleteGroupModal;
