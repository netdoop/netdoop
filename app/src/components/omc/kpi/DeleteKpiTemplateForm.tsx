import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteKPITemplate } from '@/models/kpi_temp';

interface Props {
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
  record: API.KPITemplate | undefined;
}

const DeleteKpiTemplateForm: React.FC<Props> = ({
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
        await deleteKPITemplate(record);
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
      title="Delete KPI Template"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
    >
      <p>Are you sure you want to delete this kpi template?</p>
    </Modal>
  );
};

export default DeleteKpiTemplateForm;
