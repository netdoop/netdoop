import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteUser } from '@/models/iam_users';

interface DeleteUserFormProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  user: API.User | undefined;
}

const DeleteUserForm: React.FC<DeleteUserFormProps> = ({
  visible,
  onCancel,
  onSuccess,
  user,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (user?.Id) {
        await deleteUser(user);
        onSuccess();
      }
    } finally {
      setSubmitting(false);
      onCancel();
    }
  };

  const handleCancel = async () => {
    onCancel();
  };

  return (
    <Modal
      title="Delete User"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
    >
      <p>Are you sure you want to delete this user?</p>
    </Modal>
  );
};

export default DeleteUserForm;
