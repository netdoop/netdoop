import React from 'react';
import { history } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
import { EditKpiTemplateForm } from '@/components/omc/kpi';
import { useParams } from '@umijs/max';

const CreateKpiTemplatePage: React.FC = () => {
  const params = useParams();

  const handleClose = (event: any) => {
    event.preventDefault();
    history.back()
  };
  return (
    <PageContainer
    title="Edit KPI Template"
      onBack={handleClose}
    >
      <EditKpiTemplateForm id={Number(params.id)} />
    </PageContainer >
  );
};

export default CreateKpiTemplatePage;
