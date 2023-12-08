
import React, { useEffect, useState } from 'react';
import { Form, Tabs, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { ProFormText, LoginFormPage } from '@ant-design/pro-components';
import { history, useModel } from '@umijs/max';
import { TabsProps } from 'antd';

import services from '@/services/netdoop';
const { postAuth } = services.auth;

import bg from '@/assets/img/bg.jpeg';

type LoginType = 'iam' | 'account';

const WelcomePage = () => {
    const [form] = Form.useForm();
    const [loading, setLoading] = useState(false);
    const [loginType, setLoginType] = useState<LoginType>('iam');
    const [sysinfo, setSysinfo] = useState<API.systemInfoData>({
        Name: "Netdoop",
        Version: "",
        Build: "",
    });

    const { initialState, refresh } = useModel('@@initialState');

    useEffect(() => {
        refresh()
        setSysinfo(initialState?.systemInfo||{
            Name: "Netdoop",
            Version: "",
            Build: "",
        });
        if (initialState?.loginUser !== undefined) {
            history.push('/');
        }
      }, []);

    const handleSetDebug =async () => {
        // form.setFieldsValue({
        //     username: 'admin',
        //     password: 'ann2017',
        //     remember: true,
        // });
    }
    const handleSubmit = async (values: API.authBody) => {
        setLoading(true);
        try {
            const response = await postAuth(values);
            setLoading(false);
            if (response && response.token) {
                localStorage.setItem('auth_token', response.token);
                await refresh()
                history.push('/');
            } else {
                message.error('Invalid username or password');
            }
        } catch (error) {
            localStorage.removeItem('auth_token');
            setLoading(false);
            message.error('Failed to login');
        }
    };
    const items: TabsProps['items'] = [
        {
            key: 'iam',
            label: <span>IAM User</span>,
        },
    ]
    return (
        <div style={{
            backgroundColor: 'white',
            backgroundSize: 'auto',
            height: 'calc(100vh)',  overflow: 'hidden'
        }}>
            <LoginFormPage
                title={sysinfo.Name}
                subTitle={sysinfo.Version}
                form={form}
                onFinish={handleSubmit}
                loading={loading}
                backgroundImageUrl={bg}
            >
                <Tabs
                    centered
                    activeKey={loginType}
                    onChange={(activeKey) => setLoginType(activeKey as LoginType)}
                    items={items}
                >
                </Tabs>
                {loginType === 'iam' && (
                    <>
                        <ProFormText
                            name="username"
                            fieldProps={{
                                size: 'large',
                                prefix: <UserOutlined className={'prefixIcon'} />,
                            }}
                            placeholder={'Username'}
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input username!',
                                },
                            ]}
                        />
                        <ProFormText.Password
                            name="password"
                            fieldProps={{
                                size: 'large',
                                prefix: <LockOutlined className={'prefixIcon'} />,
                            }}
                            placeholder={'Password'}
                            rules={[
                                {
                                    required: true,
                                    message: 'Please input password!',
                                },
                            ]}
                        />
                    </>
                )}
                <div style={{ marginBlockEnd: 24 }}>
                    {/* <ProFormCheckbox noStyle name="autoLogin">AutoLogin</ProFormCheckbox> */}
                    <a style={{ float: 'right', }}  >Forget Password</a>
                    <a onClick={handleSetDebug}>. </a>
                </div>
            </LoginFormPage>
        </div>
    );
};

export default WelcomePage;