import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteRole } from '@/models/iam_roles';

interface DeleteRoleFormProps {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  role: API.Role | undefined;
}

const DeleteRoleForm: React.FC<DeleteRoleFormProps> = ({
  visible,
  onCancel,
  onSuccess,
  role,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (role?.Id) {
        await deleteRole(role);
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
      title="Delete Role"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
    >
      <p>Are you sure you want to delete this role?</p>
    </Modal>
  );
};

export default DeleteRoleForm;
