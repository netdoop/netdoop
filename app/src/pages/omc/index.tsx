import { history } from '@umijs/max';
import { PageContainer  } from '@ant-design/pro-components';


const HomePage: React.FC = () => {
  history.push('/omc/dashboard')


  return (
    <PageContainer>
    </PageContainer >
  );
};

export default HomePage;