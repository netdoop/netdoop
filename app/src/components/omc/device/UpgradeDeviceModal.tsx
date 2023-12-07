import { useState } from 'react';
import { Form, Modal, Select } from 'antd';
import { useProductFirmwares } from '@/models/product';
import { upgradeDevice } from '@/models/device';

interface Props {
  visible: boolean;
  device?: API.Device;
  onCancel: () => void;
  onSuccess: () => void;
}

const UpgradeDeviceModel: React.FC<Props> = ({
  visible,
  device,
  onCancel,
  onSuccess
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);
  const { firmwares } = useProductFirmwares(device?.ProductId)

  const options = firmwares.map(v => ({
    label: v.Version,
    value: v.Id,
  }))

  const handleSubmit = async () => {
    try {
      setSubmitting(true)
      if (device) {
        const values = await form.validateFields();
        const { Firmware } = values
        await upgradeDevice(device, Firmware)
        onSuccess();
      }
    } finally {
      onCancel();
      setSubmitting(false);
    }
  };

  const handleCancel = () => {
    onCancel()
  };

  return (
    <Modal
      title="Upgrade Device"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="Firmware"
          name="Firmware"
          rules={[{ required: true, message: 'Please select firmware' }]}
        >
          <Select options={options} placeholder="Select firmware" />
        </Form.Item>
      </Form>
    </Modal>
  );
};
export default UpgradeDeviceModel;