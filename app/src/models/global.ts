// 全局共享数据示例
import { useState } from 'react';

const useUser = () => {
  const [name, setName] = useState<string>("");
  const [schema, setSchema] = useState<string>(""); 
  return {
    name,
    setName,
    schema,
    setSchema,
  };
};

export default useUser;
