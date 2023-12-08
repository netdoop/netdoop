import React, { useEffect, useRef, useState } from 'react';
import { Tree } from 'antd';
import { useKPIMeasures } from '@/models/kpi_meas';
import { treeLayout } from '@/constants/style';

interface TreeNode {
  measTypeSet: string,
  title?: JSX.Element;
  key: string | number;
  children?: TreeNode[];
}

interface Props {
  onSelect?: (selectedKeys: React.Key[]) => void;
}

const KpiMeasSetTree: React.FC<Props> = ({
  onSelect,
}) => {
  const { measuresSets } = useKPIMeasures("enb");
  const [nodes, setNodes] = useState<TreeNode[]>();
  const ref = useRef<any>();

  const handleSelect = (selectedKeys: React.Key[]) => {
    if (onSelect) {
      onSelect(selectedKeys)
    }
  };

  useEffect(() => {
    const children: TreeNode[] = Object.keys(measuresSets).map((set) => ({
      measTypeSet: set,
      title: (<span>{set}</span>),
      key: set,
    }));
    const update: TreeNode[] = [
      {
        measTypeSet: "",
        title: (<span>All</span>),
        key: "All",
        children: children,
      },
    ]
    setNodes(update)
    const keys = ['Customize'];
    ref.current?.setState(
      {
        selectedKeys: keys,
      }
    )
    handleSelect(keys);
  }, [measuresSets]);


  return (
    <>
      <Tree
       {...treeLayout}
        ref={ref}
        showLine
        treeData={nodes}
        expandedKeys={["All"]}
        onSelect={handleSelect}
        blockNode={true}
      />
    </>
  );
};

export default KpiMeasSetTree;
