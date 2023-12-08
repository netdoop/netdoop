import { useGroups } from '@/models/groups';
import { TreeSelect } from 'antd';

interface Props {
  value?: string;
  onChange?: (value: string) => void;
  hiddenGroupID?: number|null;
}

const GroupTreeSelect: React.FC<Props> = ({ value, onChange, hiddenGroupID }) => {
  const { groups } = useGroups();

  if (!groups) {
    return <TreeSelect loading />;
  }

  const renderTreeNode = (node: API.Group): React.ReactNode => {
    const { Id, Name, Children } = node;
    return (
      <TreeSelect.TreeNode key={Id} value={Id||0} title={Name} disabled={Id===hiddenGroupID}>
        {Children && Children.map(child => renderTreeNode(child))}
      </TreeSelect.TreeNode>
    );
  };

  return (
    <TreeSelect
      value={value||'0'}
      onChange={onChange}
      placeholder="Select a group"
      allowClear
      showSearch
      treeNodeFilterProp="title"
    >
      {groups.map(group => renderTreeNode(group))}
    </TreeSelect>
  );
};

export default GroupTreeSelect;

