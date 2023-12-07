import React, { useEffect, useState } from 'react';
import { useIntl } from '@umijs/max';
import { Checkbox, Modal } from 'antd';
import { useRules } from '@/models/iam_rules';
import { uniq } from 'lodash';
import { ProCard } from '@ant-design/pro-components';
import { setRoleRules } from '@/models/iam_roles';

interface Props {
  visible: boolean;
  role: API.Role | undefined;
  onCancel: () => void;
  onSuccess: () => void;
}


const EditRoleRulesModal: React.FC<Props> = ({ 
  visible,
  role,
  onCancel,
  onSuccess,
 }) => {
  const intl = useIntl();
  const { rules } = useRules();
  const [submitting, setSubmitting] = useState(false);
  const [selectedRules, setSelectedRules] = useState<string[]>([]);

  const filterByKeys = (rules: string[], keys: string[]) =>
    uniq(
      rules
        .filter((v) => keys.includes(v.split('.')[1]))
        .map((v) => v.split('.')[1] + '.' + v.split('.')[2])
    );
  const options = filterByKeys(rules, ['iam', 'omc']);

  const handleSubmit = async () => {
    try {
      setSubmitting(true);
      if (role?.Id){
        await setRoleRules(role, {RoleIds: selectedRules});
        onSuccess();
      }
    } finally {
      setSubmitting(false);
      onCancel()
    }
  };

  const handleCancel = async () => {
    onCancel()
  }

  const handleCheckboxChange = (option: string, values: any[]) => {
    const updatedValues = values.map((v) => `${v}`);
    const otherValues = selectedRules.filter((v) => v.startsWith(`api.${option}.`) === false);
    const allValues = uniq([...otherValues, ...updatedValues])
    setSelectedRules(allValues);
  };

  useEffect(() => {
    if (role) {
      setSelectedRules(role.Rules || []);
    }
  }, [role]);

  return (
    <Modal
      title="Set Role Rules"
      open={visible}
      onCancel={handleCancel}
      onOk={handleSubmit}
      confirmLoading={submitting}
      maskClosable={false}
      forceRender
      width={800}
    >
      <ProCard split="horizontal" ghost gutter={16}>
        {options.map((option) => (
          <ProCard key={option} ghost gutter={8}>
            <Checkbox.Group
              value={selectedRules.filter((value: string) =>
                value.startsWith(`api.${option}.`)
              )}
              onChange={(values) => handleCheckboxChange(option, values)}
              style={{ width: '100%' }}
            >
              <ProCard key={option} ghost gutter={8} title={option.toUpperCase()} wrap>
                {rules
                  .filter((rule) => rule.startsWith(`api.${option}.`))
                  .map((rule) => (
                    <ProCard key={rule} colSpan={8} ghost>
                      <Checkbox key={rule} value={rule}>
                        {intl.formatMessage({ id: rule })}
                      </Checkbox>
                    </ProCard>
                  ))}
              </ProCard>
            </Checkbox.Group>
          </ProCard>
        ))}
      </ProCard>
    </Modal>
  );
};

export default EditRoleRulesModal;
