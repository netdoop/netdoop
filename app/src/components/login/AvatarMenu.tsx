
import React, { useState } from 'react';
import { AvatarProps, Space, Dropdown } from 'antd';
import { getIntl, history } from '@umijs/max';
import { UserOutlined, LockOutlined, LogoutOutlined } from '@ant-design/icons';
import { ChangePasswordForm } from '@/components/login';
import { MenuProps } from 'antd';

interface AvatarMenuProps {
  props: AvatarProps;
  defaultDom: React.ReactNode;
}

const AvatarMenu: React.FC<AvatarMenuProps> = ({ defaultDom }) => {
  const intl = getIntl();
  const [changePasswordModalVisible, setChangePasswordModalVisible] = useState(false);

  const handleLogout = async () => {
    localStorage.removeItem('auth_token');
    history.push('/login');
  };
  const handleProfile = async () => {
    history.push('/profile');
  };
  const handleChangePassword = async () => {
    setChangePasswordModalVisible(true);
  };
  const items: MenuProps['items'] = [
    {
      key: 'profile',
      label: (
        <a onClick={handleProfile}>
          {intl.formatMessage({ id: 'layout.actions.profile' })}
        </a>
      ),
      icon: <UserOutlined />,
      disabled: true,
    },
    {
      key: 'change-password',
      label: (
        <a onClick={handleChangePassword}>
          {intl.formatMessage({ id: 'layout.actions.change_password' })}
        </a>
      ),
      icon: <LockOutlined />,
    },
    {
      key: 'logout',
      label: (
        <a onClick={handleLogout}>
          {intl.formatMessage({ id: 'layout.actions.logout' })}
        </a>
      ),
      icon: <LogoutOutlined />,
    },
  ];

  return (
    <>
      <Dropdown menu={{ items }}>
        <a onClick={(e) => e.preventDefault()}>
          <Space  style={{marginRight: 8}}>{defaultDom}</Space>
        </a>
      </Dropdown>
      <ChangePasswordForm
        visible={changePasswordModalVisible}
        onCancel={() => setChangePasswordModalVisible(false)}
        onSuccess={() => setChangePasswordModalVisible(false)}
      />
    </>

  );
};

export default AvatarMenu;
