import React, { useEffect, useState } from 'react';
import { Tree } from 'antd';
import { fetchGroups } from '@/models/groups';
import { treeLayout } from '@/constants/style';

interface GroupTreeNode {
  group: API.Group,
  title?: JSX.Element;
  key: string | number;
  children?: GroupTreeNode[];
}
interface Props {
  value?: React.Key[];
  onChange?: (value: React.Key[]) => void;
}

const SelectGroupInput: React.FC<Props> = ({
  value,
  onChange,
}) => {
  const [groupTree, setGroupTree] = useState<GroupTreeNode[]>([]);

  const renderTitle = (group: API.Group) => (
      <span>{group.Name}</span>
  );

  const buildGroupTree = (groups: API.Group[]): GroupTreeNode[] => {
    const treeData = groups.map((group) => ({
      group: group,
      title: renderTitle(group),
      key: group.Id || 0,
      children: group.Children ? buildGroupTree(group.Children) : undefined,
    }));
    return treeData;
  };

  const handleFetchGroups = async () => {
    const groups = await fetchGroups();
    const treeData = buildGroupTree(groups);
    setGroupTree(treeData);
  }

  useEffect(() => {
    handleFetchGroups();
  }, []);

  const handleSelect = (selectedKeys: React.Key[]) => {
    if (onChange) {
      onChange(selectedKeys)
    }
  };

  return (
    <Tree
    {...treeLayout}
    checkable
      showLine
      treeData={groupTree}
      expandedKeys={[0]}
      defaultExpandAll
      onSelect={handleSelect}
      selectedKeys={value}
      blockNode={true}
    />
  );
};

export default SelectGroupInput;
