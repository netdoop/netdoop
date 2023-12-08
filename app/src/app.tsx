// 运行时配置
import { AntdConfig, history } from '@umijs/max';
import { message } from 'antd';
import { theme } from 'antd';
import { actionsRender, avatarRender } from './render';
import { RuntimeConfig } from '@umijs/max';

import services from '@/services/netdoop';

import logo from '@/assets/img/logo.png';

export const request: RuntimeConfig['request'] = {
  baseURL: "/api/v1",
  timeout: 15000,
  // other axios options you want
  errorConfig: {
    // 错误抛出
    errorThrower: () => {
    },
    // 错误接收及处理
    errorHandler: (error: any, opts: any) => {
      if (opts?.skipErrorHandler) throw error;
      // 我们的 errorThrower 抛出的错误。
      if (error.response.status) {
        if (error.response.data.message) {
          message.error(`${error.response.status} ${error.response.statusText}: ${error.response.data?.message}`);
        } else {
          message.error(`${error.response.status} ${error.response.statusText}`);
        }
        if (error.response.status === 401) {
          history.push('/login')
        }
      } else if (error.request) {
        // 请求已经成功发起，但没有收到响应
        // \`error.request\` 在浏览器中是 XMLHttpRequest 的实例，
        // 而在node.js中是 http.ClientRequest 的实例
        message.error('None response! Please retry.');
      } else {
        // 发送请求时出了点问题
        message.error('Request error, please retry.');
      }
    },

  },
  requestInterceptors: [
    (url, options) => {
      const token = localStorage.getItem('auth_token');
      if (typeof token === 'string') {
        options.headers = {
          ...options.headers,
          Authorization: `Bearer ${token}`,
        };
      }
      return { url, options };
    },
  ],
  responseInterceptors: [
    (response) => {
      if (response.status === 401 || response.status === 403) {
        message.error('You are not authorized to access this page.');
      }
      return response;
    },
    (error) => {
      return error
    }
  ]
};

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
export async function getInitialState(): Promise<{
  systemInfo: API.systemInfoData| undefined,
  name: string,
  avatar: string,
  loginUser: API.User | undefined,
}> {
  let loginUser: API.User | undefined = undefined;
  let systemInfo: API.systemInfoData | undefined = undefined;

  try {
    systemInfo = await services.system.getSystemInfo();
  } catch (error) {
  } finally {
  }

  try {
    const token = localStorage.getItem('auth_token');
    if (typeof token === 'string') {
      loginUser = await services.current.getCurrent({});
    }
  } catch (error) {
    loginUser = undefined
    // console.error('Failed to get current user', error);
  } finally {
    // console.log("getInitialState final", loginUser)
  }

  if (loginUser) {
    return {
      systemInfo,
      name: loginUser?.Name || '',
      avatar: 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png',
      loginUser,
    };
  }
  history.push('/login')
  return {
    systemInfo,
    name: '',
    avatar: '',
    loginUser,
  };
}

export const antd: RuntimeConfig['antd'] = (memo: AntdConfig) => {
  memo.theme ||= {};
  memo.theme.algorithm = (theme.compactAlgorithm);
  return memo;
};

export const layout: RuntimeConfig['layout'] = (initData) => {
  // const intl = getIntl();
  const { initialState } = initData;
  console.log("1000", initialState?.systemInfo)
  return {
    // 常用属性
    title: initialState?.systemInfo?.Name || 'NetDoop',
    logo: logo,
    contentWidth: 'Fluid',
    layout: 'mix',
    fixedHeader: true,
    actionsRender: actionsRender,
    siderWidth: 220,
    token: {
      header: {
        heightLayoutHeader: 36,
      }
    },
    defaultCollapsed: true,
    breakpoint: false,
    splitMenus: true,
    menu: {
      locale: true,
    },
    avatarProps: {
      title: initialState?.name,
      icon: <img src="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png" alt="avatar"></img>,
      render: avatarRender,
    },
    logout: () => {
      const { refresh } = initData;
      localStorage.removeItem('auth_token')
      refresh()
    },
    onPageChange: () => {
      // const { initialState, refresh } = initData;
      // const { location } = history;
      console.log("onPageChange", location.pathname, initialState?.loginUser)
      // if (!initialState?.loginUser) {
      //   refresh()
      // }
      // if (!initialState?.loginUser && location.pathname !== '/login') {
      //   history.push('/login');
      // }
      // if (initialState?.loginUser && location.pathname === '/login') {
      //   history.push('/')
      // }
    },
    // 其他属性见：https://procomponents.ant.design/components/layout#prolayout
  };
};


