import { Suspense, lazy, useEffect } from 'react';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import { useNavigate } from 'react-router-dom';
import Loading from '../../components/Loading/Loading';
import { useSession } from '../../hooks/useSession';

const LoginForm = lazy(() => import('../../components/Forms/Login/Login'));

export function Login() {
  useDocumentTitle('Login');
  const { isLoggedIn } = useSession();
  const navigate = useNavigate();
  useEffect(() => {
    if (isLoggedIn) {
      navigate('/todos');
    }
  }, [isLoggedIn, navigate]);

  return (
    <Suspense fallback={<Loading />}>
      <LoginForm />
    </Suspense>
  );
}
