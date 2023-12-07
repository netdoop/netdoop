import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteFirmware } from '@/models/firmware';

interface DeleteFirmwareModalProps {
  visible: boolean;
  firmware?: API.Firmware;
  onCancel: () => void;
  onSuccess: () => void;
}

const DeleteFirmwareModal: React.FC<DeleteFirmwareModalProps> = ({
  visible,
  firmware,
  onCancel,
  onSuccess,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleDelete = async () => {
    if (firmware) {
      setSubmitting(true);
      try {
        await deleteFirmware(firmware);
        onSuccess();
      } finally {
        onCancel();
        setSubmitting(false);
      }
    }
  };

  const handleCancel = () => {
    onCancel();
  };

  return (
    <Modal
      title="Delete Firmware"
      open={visible}
      onCancel={handleCancel}
      onOk={handleDelete}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <p>Are you sure you want to delete the firmware: {firmware?.Version}?</p>
    </Modal>
  );
};

export default DeleteFirmwareModal;
