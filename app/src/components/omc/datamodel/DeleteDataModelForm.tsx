import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteDataModel } from '@/models/datamodel';

interface Props {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  record: API.KPITemplate | undefined;
}

const DeleteDataModelForm: React.FC<Props> = ({
  visible,
  onCancel,
  onSuccess,
  record,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (record?.Id) {
        await deleteDataModel(record);
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
      title="Delete Data Model"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
    >
      <p>Are you sure you want to delete this data model?</p>
    </Modal>
  );
};

export default DeleteDataModelForm;
