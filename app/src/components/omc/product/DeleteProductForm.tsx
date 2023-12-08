import React, { useState } from 'react';
import { Modal, message } from 'antd';
import services from '@/services/netdoop';
const { deleteProduct } = services.omcProducts;

interface DeleteProductFormProps {
  visible: boolean;
  product: API.Product | null;
  onCancel: () => void;
  onSuccess: () => void;
}

const DeleteProductForm: React.FC<DeleteProductFormProps> = ({
  visible,
  product,
  onCancel,
  onSuccess,
}) => {
  const [submitting, setSubmitting] = useState(false);

  const handleDelete = async () => {
    if (product?.Id) {
      setSubmitting(true);
      try {
        await deleteProduct({ id: product.Id });
        message.success('Device type deleted successfully');
        onSuccess();
      } catch (error) {
        message.error('Failed to delete product');
      } finally {
        setSubmitting(false);
        onCancel();
      }
    }
  };

  const handleCancel = () => {
    onCancel();
  };

  return (
    <Modal
      title="Delete Product"
      open={visible}
      onCancel={handleCancel}
      onOk={handleDelete}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
    >
      <p>Are you sure you want to delete the product {product?.Oui} - {product?.ProductClass}?</p>
    </Modal>
  );
};

export default DeleteProductForm;
