import React, { useState, useEffect } from 'react';
import { Modal, Form } from 'antd';
import GroupTreeSelect from '../group/GroupTreeSelect';
import { setGroupForDevice } from '@/models/device';

interface SetDeviceGroupModalProps {
  visible: boolean;
  device?: API.Device;
  onCancel: () => void;
  onSuccess: () => void;
}

const SetDeviceGroupModal: React.FC<SetDeviceGroupModalProps> = ({
  visible,
  device,
  onCancel,
  onSuccess,
}) => {
  const [submitting, setSubmitting] = useState(false);
  const [form] = Form.useForm();

  useEffect(() => {
    if (visible) {
      form.setFieldsValue(device);
    }
  }, [visible, form, device]);


  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (device) {
        const values = await form.validateFields();
        await setGroupForDevice(device, values);
        onSuccess();
      }
    } finally {
      onCancel();
      setSubmitting(false);
      form.resetFields()
    }
  };

  const handleCancel = () => {
    onCancel()
    form.resetFields()
  };

  return (
    <Modal
      title="Set Device Group"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Group"
          name="GroupId"
          rules={[{ required: true, message: 'Please select a group' }]}
        >
          <GroupTreeSelect />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default SetDeviceGroupModal;