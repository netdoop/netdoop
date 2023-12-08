import { useEffect, useState } from 'react';
import { Form, Modal, Select } from 'antd';
import { uploadDeviceFile } from '@/models/device';

interface Props {
  visible: boolean;
  device?: API.Device;
  onCancel: () => void;
  onSuccess: () => void;
}

const UploadFileModal: React.FC<Props> = ({
  visible,
  device,
  onCancel,
  onSuccess
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);
  const [options, setOptions] = useState<{ label: string; value: number; }[]>([]);

  useEffect(() => {
    const items = [
      { label: 'Configuration File', value: 1 },
      { label: 'Log File', value: 2 },
    ];
    if (device?.ProductType === 'enb'){
      items.push({label: "NRM File", value: 4})
    }
    setOptions(items)
  }, [device]);


  const handleSubmit = async () => {
    try {
      setSubmitting(true)
      if (device) {
        const values = await form.validateFields();
        const { FileType } = values
        if (FileType === 1) {
          await uploadDeviceFile(device,"1 Vendor Configuration File" )
        } else if (FileType === 2) {
          await uploadDeviceFile(device,"2 Vendor Log File" )
        } else if (FileType === 4) {
          await uploadDeviceFile(device,"4 Vendor Log File" )
        }
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
      title="Upload Device File"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item
          label="File Type"
          name="FileType"
          rules={[{ required: true, message: 'Please select file type' }]}
        >
          <Select options={options} placeholder="Select file type" />
        </Form.Item>
      </Form>
    </Modal>
  );
};
export default UploadFileModal;