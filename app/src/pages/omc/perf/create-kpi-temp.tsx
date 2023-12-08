import React from 'react';
import { history } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
import { EditKpiTemplateForm } from '@/components/omc/kpi';

const CreateKpiTemplatePage: React.FC = () => {
  const handleClose = (event: any) => {
    event.preventDefault();
    history.back()
  };
  return (
    <PageContainer
      title="Create KPI Template"
      onBack={handleClose}
    >
      <EditKpiTemplateForm />
    </PageContainer >
  );
};

export default CreateKpiTemplatePage;
