import { defineConfig } from '@umijs/max';

export default defineConfig({
  hash: true,
  history: {
    type: 'hash',
  },
  locale: {
    default: 'en-US',
    baseSeparator: '-',
    antd: true,
    title: false,
    baseNavigator: false,
  },
  access: {},
  antd: {
    // configProvider
    configProvider: {
      componentSize: 'small',
      pagination: {
        showSizeChanger: true,
      },
    },
    // themes
    dark: false,
    compact: true,
    // babel-plugin-import
    import: false,
    // less or css, default less
    style: 'less',
    // shortcut of `configProvider.theme`
    // use to configure theme token, antd v5 only
    theme: {},
    // antd <App /> valid for version 5.1.0 or higher, default: undefined
    appConfig: {},
  },
  theme: {
    '@import': '~src/global.less',
  },
  model: {},
  initialState: {},
  request: {},
  layout: {
    title: 'NetDoop',
  },
  favicons: ['/favicon.ico'],
  routes: [
    {
      path: '/',
      redirect: '/omc/dashboard',
    },
    {
      name: 'login',
      path: '/login',
      layout: false,
      component: './login',
    },
    {
      name: 'omc',
      path: '/omc',
      icon: 'CloudOutlined',
      access: 'canViewOMCCpePages',
      routes: [
        {
          path: '/omc',
          redirect: '/omc/dashboard',
        },
        {
          name: 'dashboard',
          path: '/omc/dashboard',
          icon: 'DashboardOutlined',
          component: './omc/dashboard',
          access: 'canViewOMCDashboardPage',
        },
        {
          name: 'enb',
          path: 'enb',
          icon: 'WifiOutlined',
          access: 'canViewOMCCpePages',
          routes: [
            {
              path: '/omc/enb',
              redirect: '/omc/enb/monitor',
            },
            {
              name: 'monitor',
              path: 'monitor',
              component: './omc/enb/monitor',
              access: 'canViewOMCCpeMonitorPage',
            },
            {
              path: 'monitor/:id/information',
              component: './omc/enb/monitor/information',
              access: 'canViewOMCCpeMonitorInformationPage',
            },
            {
              path: 'monitor/:id/setting',
              access: 'canViewOMCCpeMonitorSettingPage',
              routes: [
                {
                  path: '/omc/enb/monitor/:id/setting',
                  redirect: '/omc/enb/monitor/:id/setting/Network/WANSetting',
                },
                {
                  path: ':key/:sub',
                  component: './omc/enb/monitor/setting',
                  access: 'canViewOMCCpeMonitorSettingPage',
                },
              ],
            },

            {
              name: 'log',
              path: 'log',
              access: 'canViewOMCCpeLogsPage',
              routes: [
                {
                  path: '/omc/enb/log',
                  redirect: '/omc/enb/log/events',
                },
                {
                  path: ':key',
                  component: './omc/enb/log',
                  access: 'canViewOMCCpeLogsPage',
                },
                {
                  path: 'events/:ts/detail',
                  component: './omc/enb/log/event/detail',
                  access: 'canViewOMCCpeEventDetailPage'
                },
                {
                  path: 'method-calls/:ts/detail',
                  component: './omc/enb/log/method-call/detail',
                  access: 'canViewOMCCpeMethodCallDetailPage'
                },
              ],
            },

            {
              name: 'upload',
              path: 'upload',
              access: 'canViewOMCCpeUploadFilesPage',
              routes: [
                {
                  path: '/omc/enb/upload',
                  redirect: '/omc/enb/upload/Configuraions',
                },
                {
                  path: ':key',
                  component: './omc/enb/upload',
                  access: 'canViewOMCCpeUploadFilesPage',
                },
              ],
            },
            {
              name: 'upgrade',
              path: 'upgrade',
              component: './omc/enb/upgrade',
              access: 'canViewOMCCpeUpgradePage',
            },
            {
              name: 'inventory',
              path: 'inventory',
              component: './omc/enb/inventory',
              access: 'canViewOMCCpeInventoryPage',
            },
            {
              path: 'inventory/:id/information',
              component: './omc/cpe/inventory/information',
              access: 'canViewOMCCpeInventoryInformationPage',
            },
          ],
        },
        {
          name: 'cpe',
          path: 'cpe',
          icon: 'MobileOutlined',
          access: 'canViewOMCCpePages',
          routes: [
            {
              path: '/omc/cpe',
              redirect: '/omc/cpe/monitor',
            },
            {
              name: 'monitor',
              path: 'monitor',
              component: './omc/cpe/monitor',
              access: 'canViewOMCCpeMonitorPage',
            },
            {
              path: 'monitor/:id/information',
              component: './omc/cpe/monitor/information',
              access: 'canViewOMCCpeMonitorInformationPage',
            },
            {
              path: 'monitor/:id/statistics',
              component: './omc/cpe/monitor/statistics',
              access: 'canViewOMCCpeMonitorStatisticsPage',
            },
            {
              path: 'monitor/:id/setting',
              component: './omc/cpe/monitor/setting',
              access: 'canViewOMCCpeMonitorSettingPage',
              routes: [
                {
                  path: '/omc/cpe/monitor/:id/setting',
                  redirect: '/omc/cpe/monitor/:id/setting/Network/WANSetting',
                },
                {
                  path: ':key/:sub',
                  component: './omc/cpe/monitor/setting',
                  access: 'canViewOMCCpeMonitorSettingPage',
                },
              ],
            },

            {
              name: 'log',
              path: 'log',
              access: 'canViewOMCCpeLogsPage',
              routes: [
                {
                  path: '/omc/cpe/log',
                  redirect: '/omc/cpe/log/events',
                },
                {
                  path: ':key',
                  component: './omc/cpe/log',
                  access: 'canViewOMCCpeLogsPage',
                },
                {
                  path: 'events/:ts/detail',
                  component: './omc/cpe/log/event/detail',
                  access: 'canViewOMCCpeEventDetailPage'
                },
                {
                  path: 'method-calls/:ts/detail',
                  component: './omc/cpe/log/method-call/detail',
                  access: 'canViewOMCCpeMethodCallDetailPage'
                },
              ],
            },

            {
              name: 'upload',
              path: 'upload',
              access: 'canViewOMCCpeUploadFilesPage',
              routes: [
                {
                  path: '/omc/cpe/upload',
                  redirect: '/omc/cpe/upload/Configuraions',
                },
                {
                  path: ':key',
                  component: './omc/cpe/upload',
                  access: 'canViewOMCCpeUploadFilesPage',
                },
              ],
            },
            {
              name: 'upgrade',
              path: 'upgrade',
              component: './omc/cpe/upgrade',
              access: 'canViewOMCCpeUpgradePage',
            },
            {
              name: 'inventory',
              path: 'inventory',
              component: './omc/cpe/inventory',
              access: 'canViewOMCCpeInventoryPage',
            },
            {
              path: 'inventory/:id/information',
              component: './omc/cpe/inventory/information',
              access: 'canViewOMCCpeInventoryInformationPage',
            },
          ],
        },
        {
          name: 'alarms',
          path: 'alarms',
          icon: 'WarningOutlined',
          component: './omc/alarm',
          access: 'canViewOMCAlarmsPage',
        },
        {
          path: 'alarms/:ts/detail',
          component: './omc/alarm/detail',
          access: 'canViewOMCAlarmDetailPage',
        },
        {
          name: 'perf',
          path: 'perf',
          icon: 'FundOutlined',
          access: 'canViewOMCPerfPage',
          routes: [
            {
              path: '/omc/perf',
              redirect: '/omc/perf/KpiView',
            },
            {
              path: ':key',
              component: './omc/perf',
              access: 'canViewOMCPerfPage',
            },
            {
              path: 'KpiTemp/create',
              component: './omc/perf/create-kpi-temp',
              access: 'canViewOMCPerfPage',
            },
            {
              path: 'KpiTemp/edit/:id',
              component: './omc/perf/edit-kpi-temp',
              access: 'canViewOMCPerfPage',
            },
          ],
        },
        {
          name: 'admin',
          path: '/omc/admin',
          icon: 'SettingOutlined',
          access: 'canViewOMCAdminPages',
          routes: [
            {
              name: 'products',
              path: 'products',
              component: './omc/admin/products',
              access: 'canViewOMCProductsPage',
            },
          ],
        },
      ],
    },
    {
      name: 'system',
      path: '/system',
      icon: 'SettingOutlined',
      access: 'canViewSystemPages',
      routes: [
        {
          path: '/system',
          redirect: '/system/iam',
        },
        {
          name: 'iam',
          path: '/system/iam',
          icon: 'UserOutlined',
          component: './iam',
          access: 'canViewIAMPage',
        },
        {
          name: 'audit',
          path: '/system/audit',
          icon: 'AuditOutlined',
          component: './audit',
          access: 'canViewAuditPage',
        },
      ],
    },
  ],
  npmClient: 'pnpm',
  proxy: {
    '/api/v1': {
      // 'target': 'http://127.0.0.1:9176',
      'target': 'https://dev.netdoop.com',
      'changeOrigin': true,
    }
  },
  presets: [require.resolve('umi-presets-pro')],
  openAPI:[
    {
      requestLibPath: "import { request } from '@umijs/max'",
      schemaPath: __dirname + '/api.json',
      mock: false,
    },
  ], 
});

