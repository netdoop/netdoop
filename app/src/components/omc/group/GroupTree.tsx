import React, { useEffect, useState } from 'react';
import { Tree, Dropdown, Button } from 'antd';
import { CreateGroupModal, DeleteGroupModal, MoveGroupModal } from '.';
import { MenuProps } from 'antd';
import { useIntl } from '@umijs/max';
import { DeleteOutlined, EditOutlined, MoreOutlined } from '@ant-design/icons';
import { fetchGroups } from '@/models/groups';
import { treeLayout } from '@/constants/style';
import { useAccess } from '@umijs/max';
interface GroupTreeNode {
  group: API.Group,
  title?: JSX.Element;
  key: string | number;
  children?: GroupTreeNode[];
}
interface GroupTreeProps {
  onSelect?: (selectedKeys: React.Key[]) => void;
}

const GroupTree: React.FC<GroupTreeProps> = ({
  onSelect,
}) => {
  const access = useAccess()
  const intl = useIntl()

  const [groupTree, setGroupTree] = useState<GroupTreeNode[]>([]);
  const [isCreateGroupModalVisible, setIsCreateGroupModalVisible] = useState(false);
  const [isDeleteGroupModalVisible, setIsDeleteGroupModalVisible] = useState(false);
  const [isMoveGroupModalVisible, setIsMoveGroupModalVisible] = useState(false);
  const [selectedGroup, setSelectedGroup] = useState<API.Group | null>(null);

  const handleAddGroup = (group: API.Group) => {
    setSelectedGroup(group);
    setIsCreateGroupModalVisible(true);
  };
  const handleMoveGroup = (group: API.Group) => {
    setSelectedGroup(group);
    setIsMoveGroupModalVisible(true);
  };
  const handleDeleteGroup = (group: API.Group) => {
    setSelectedGroup(group);
    setIsDeleteGroupModalVisible(true);
  };

  const moreItems = (record: API.Group): MenuProps => {
    const items: MenuProps['items'] = [
      {
        key: 'information',
        icon: (<EditOutlined />),
        disabled: !access.canCreateOMCGroup,
        label: (
          <a onClick={() => handleAddGroup(record)}>
            {intl.formatMessage({ id: 'common.add' })}
          </a>
        ),
      },
      {
        key: 'move',
        icon: (<EditOutlined />),
        disabled: !access.canUpdateOMCGroupParent,
        label: (
          <a onClick={() => handleMoveGroup(record)}>
            {intl.formatMessage({ id: 'common.move' })}
          </a>
        ),
      },
      {
        key: 'delete',
        icon: (<DeleteOutlined />),
        disabled: !access.canDeleteOMCGroup,
        label: (
          <a onClick={() => handleDeleteGroup(record)}>
            {intl.formatMessage({ id: 'common.delete' })}
          </a>
        ),
      },
    ]
    return { items }
  }

  const renderTitle = (group: API.Group) => (
    <>
      <span>{group.Name}</span>
      <Dropdown menu={moreItems(group)}>
        <Button type="link" onClick={(e) => e.preventDefault()} style={{ float: 'right' }}>
          <MoreOutlined />
        </Button>
      </Dropdown>
    </>

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

  const handleCreateGroupSuccess = async () => {
    await handleFetchGroups()
    setIsCreateGroupModalVisible(false);
  };

  const handleMoveGroupSuccess = async () => {
    await handleFetchGroups()
    setIsMoveGroupModalVisible(false);
  };
  const handleDeleteGroupSuccess = async () => {
    setSelectedGroup(null);
    await handleFetchGroups()
    setIsDeleteGroupModalVisible(false);
  };

  const handleSelect = (selectedKeys: React.Key[]) => {
    if (onSelect) {
      onSelect(selectedKeys)
    }
  };

  return (
    <>
      {groupTree.length > 0 &&
        <Tree
          {...treeLayout}
          showLine
          treeData={groupTree}
          expandedKeys={[0]}
          defaultExpandAll
          onSelect={handleSelect}
          blockNode={true}
        />
      }
      <CreateGroupModal
        visible={isCreateGroupModalVisible}
        onCancel={() => setIsCreateGroupModalVisible(false)}
        onSuccess={handleCreateGroupSuccess}
        parent={selectedGroup}
      />
      <MoveGroupModal
        visible={isMoveGroupModalVisible}
        onCancel={() => setIsMoveGroupModalVisible(false)}
        onSuccess={handleMoveGroupSuccess}
        group={selectedGroup}
      />
      <DeleteGroupModal
        visible={isDeleteGroupModalVisible}
        onCancel={() => setIsDeleteGroupModalVisible(false)}
        onSuccess={handleDeleteGroupSuccess}
        group={selectedGroup}
      />

    </>
  );
};

export default GroupTree;
