import React, { useState } from 'react';
import { Modal } from 'antd';

import { recoverDeletedDevice } from '@/models/device';

interface Props {
  visible: boolean;
  device?: API.Device;
  onCancel: () => void;
  onSuccess: () => void;
}

const RecoverDeletedDeviceModal: React.FC<Props> = ({
  visible,
  device,
  onCancel,
  onSuccess,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (device) {
        await recoverDeletedDevice(device);
        onSuccess();
      }
    } finally {
      onCancel();
      setSubmitting(false);
    }
  };

  const handleCancel = () => {
    onCancel();
  };

  return (
    <Modal
      title="Delete Device "
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <p>Are you sure you want to recover the device: {device?.Oui} - {device?.ProductClass}?</p>
    </Modal>
  );
};

export default RecoverDeletedDeviceModal;
