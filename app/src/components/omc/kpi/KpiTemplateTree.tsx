import React, { useEffect, useState, useRef } from 'react';
import { Tree } from 'antd';
import { useKPITemplates } from '@/models/kpi_temp';
import { treeLayout } from '@/constants/style';

interface TreeNode {
  name?: string,
  title?: JSX.Element;
  key: string | number;
  children?: TreeNode[];
}

interface Props {
  onSelect?: (template: API.KPITemplate | undefined) => void;
}

const KpiTemplateTree: React.FC<Props> = ({
  onSelect,
}) => {
  const { kpiTemplates, kpiTemplateById } = useKPITemplates();
  const [nodes, setNodes] = useState<TreeNode[]>([]);
  const ref = useRef<any>();

  const handleSelect = (selectedKeys: React.Key[]) => {
    if (onSelect && selectedKeys[0]) {
      const id = selectedKeys[0] as number;
      onSelect(kpiTemplateById(id))
    }
  };

  useEffect(() => {
    const update: TreeNode[] = kpiTemplates.map((v) => ({
      name: v.Name,
      title: (<span>{v.Name}</span>),
      key: v.Id || 0,
      children: [],
    }));

    setNodes(update)
    ref.current?.setState(
      {
        selectedKeys: [update[0]?.key],
      }
    )
    handleSelect([update[0]?.key])
  }, [kpiTemplates]);



  return (
    <>
      <Tree
        {...treeLayout}
        ref={ref}
        showLine
        treeData={nodes}
        onSelect={handleSelect}
        blockNode={true}
        defaultSelectedKeys={[nodes[0]?.key]}
      />
    </>
  );
};

export default KpiTemplateTree;
