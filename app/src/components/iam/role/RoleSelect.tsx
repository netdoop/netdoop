import { useAllRoles } from '@/models/iam_roles';
import { Select } from 'antd';

interface RoleSelectProps {
  onChange?: (value: string) => void;
}

const RoleSelect: React.FC<RoleSelectProps> = ({ onChange }) => {
  const { roles } = useAllRoles();

  if (!roles) {
    return <Select loading />;
  }
  const options = roles.map(role => ({
    label: role.Alias || role.Name,
    value: role.Name,
  }));

  return (
    <Select
      options={options}
      onChange={onChange}
      placeholder="Select roles"
      allowClear
      showSearch
      mode="multiple"
    />
  );
};

export default RoleSelect;

