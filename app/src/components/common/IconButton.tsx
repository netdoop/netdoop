import { Button } from 'antd';
import { ReactNode, MouseEventHandler } from 'react';

interface IconButtonProps {
  icon: ReactNode;
  text: string;
  onClick?: MouseEventHandler<HTMLElement>;
}

const IconButton = ({ icon, text, onClick }: IconButtonProps) => {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
      <Button type="primary" size="middle" shape="circle" icon={icon} onClick={onClick} />
      <Button type="text" size="small" style={{padding:0}} onClick={onClick}>{text}</Button>
    </div>
  );
};

export default IconButton;
