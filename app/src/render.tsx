import React from 'react';
import { AvatarProps } from 'antd';
import { HeaderProps } from '@ant-design/pro-components';
import AvatarMenu from './components/login/AvatarMenu';
// import { AlarmStatsCard } from './components';
// import dayjs from 'dayjs';

export const avatarRender: (props: AvatarProps, defaultDom: React.ReactNode) => React.ReactNode = (props, defaultDom) => {
    return (
        <AvatarMenu props={props} defaultDom={defaultDom} />
    );
};

export const actionsRender: (props: HeaderProps) => React.ReactNode[] = () => {
    return [
        // <AlarmStatsCard key="alarms_stats" from={dayjs.unix(0).unix()} to={dayjs(Date.now()).unix()}/>
    ];
};

