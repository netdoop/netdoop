import React, { useState } from 'react';
import { Modal, Form, Input, Button, Upload, message, Select } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import { createFirmware } from '@/models/firmware';
import { useProducts } from '@/models/product';

interface CreateFirmwareModalProps {
  productType: string;
  visible: boolean;
  onCancel: () => void;
  onSuccess: () => void;
}

const CreateFirmwareModal: React.FC<CreateFirmwareModalProps> = ({
  productType,
  visible,
  onCancel,
  onSuccess,
}) => {
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);
  const [uploadFile, setUploadFile] = useState<File | undefined>(undefined);
  const { products, loading: productsLoading } = useProducts(productType);
  const options = products.map(product => ({
    label: product.ModelName,
    value: product.ModelName,
  }))

  const handleSubmit = async () => {
    if (uploadFile) {
      try {
        const values = await form.validateFields();
        setSubmitting(true);
        await createFirmware({
          Version: values.Version,
          Products: values.Products,
          ProductType: productType,
        }, uploadFile);
        onSuccess();
      } finally {
        setSubmitting(false);
        onCancel();
        form.resetFields();
      }
    }

  };

  const handleCancel = () => {
    form.resetFields();
    onCancel();
  };

  const beforeUpload = (file: File) => {
    const allowedExtensions = ['bin'];
    const fileExtension = file.name.split('.').pop();
    const isAllowed = allowedExtensions.includes(fileExtension || '');

    if (!isAllowed) {
      message.error(`File type not allowed. Only ${allowedExtensions.join(', ')} files are allowed.`);
      return false
    }
    const isLt10MB = file.size / 1024 / 1024 < 10;
    if (!isLt10MB) {
      message.error('Firmware file must be smaller than 10MB!');
      return false
    }
    setUploadFile(file)
    return true;
  };

  return (
    <Modal
      open={visible}
      title="Create Firmware"
      okText="Create"
      onOk={handleSubmit}
      onCancel={handleCancel}
      confirmLoading={submitting}
      forceRender
    >
      <Form form={form} layout="vertical">
        <Form.Item name="Version" label="Version" rules={[{ required: true }]}>
          <Input placeholder="Enter Version" />
        </Form.Item>
        <Form.Item name="Products" label="Products" rules={[{ required: true }]}>
          <Select
            mode="multiple"
            loading={productsLoading}
            allowClear
            placeholder="Please select"
            options={options}
          />
        </Form.Item>
        <Form.Item name="file" label="Firmware File" rules={[{ required: true }]}>
          <Upload beforeUpload={beforeUpload} >
            <Button icon={<UploadOutlined />}>Select File</Button>
          </Upload>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateFirmwareModal;
