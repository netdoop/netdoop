import React, { useState } from 'react';
import { Modal } from 'antd';
import { deleteDevice } from '@/models/device';

interface DeleteDeviceModalProps {
  visible: boolean;
  device?: API.Device;
  onCancel: () => void;
  onSuccess: () => void;
}

const DeleteDeviceModal: React.FC<DeleteDeviceModalProps> = ({
  visible,
  device,
  onCancel,
  onSuccess,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleDelete = async () => {
    try {
      setSubmitting(true);
      if (device) {
        await deleteDevice(device);
        onSuccess();
      }
    } finally{
      onCancel();
      setSubmitting(false);
    }
  };

  const handleCancel = () => {
    onCancel();
  };
  return (
    <Modal
      title="Delete Device"
      open={visible}
      onCancel={handleCancel}
      onOk={handleDelete}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <p>Are you sure you want to delete the device {device?.Oui} - {device?.ProductClass}?</p>
    </Modal>
  );
};

export default DeleteDeviceModal;
