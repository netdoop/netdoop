import React from 'react';
import { PageContainer } from '@ant-design/pro-components';
import KpiMgmtPage from './kpi-mgmt';
import PerfMgmtPage from './perf-mgmt';
import KpiTempPage from './kpi-temp';
import { useParams } from '@umijs/max';
import { Tabs } from 'antd';
import { history } from '@umijs/max';
import KpiValuePage from './kpi-view';
import KpiValueChartsPage from './kpi-chart';

const PerfPage: React.FC = () => {
  const params = useParams();

  const handleTabChange = (key: string) => {
    history.push(`/omc/perf/${key}`);
  };
  const tabItems = [
    {
      label: 'KPI View',
      key: 'KpiView',
      children: <KpiValuePage />,
    },
    {
      label: 'KPI Charts',
      key: 'KpiCharts',
      children: <KpiValueChartsPage />,
    },
    {
      label: 'KPI Template',
      key: 'KpiTemp',
      children: <KpiTempPage />,
    },
    {
      label: 'KPI Mgmt',
      key: 'KpiMgmt',
      children: <KpiMgmtPage />,
    },
    {
      label: 'Perf Mgmt',
      key: 'PerfMgmt',
      children: <PerfMgmtPage />,
    },
  ];

  return (
    <PageContainer
      header={{
        title: '',
      }}
    >
      <Tabs
        tabPosition="top"
        items={tabItems}
        defaultActiveKey={params.key || "KpiView"}
        onChange={(v) => handleTabChange(v)}
      />
    </PageContainer>
  );
};

export default PerfPage;
