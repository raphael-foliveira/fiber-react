import { Suspense, lazy, useEffect } from 'react';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import { useNavigate } from 'react-router-dom';
import Loading from '../../components/Loading/Loading';

const LoginForm = lazy(() => import('../../components/Forms/Login/Login'));

export function Login() {
  useDocumentTitle('Login');
  const navigate = useNavigate();
  useEffect(() => {
    if (localStorage.getItem('accessToken')) {
      navigate('/todos');
    }
  });

  return (
    <Suspense fallback={<Loading />}>
      <LoginForm />
    </Suspense>
  );
}
