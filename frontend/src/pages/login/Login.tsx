import { Suspense, lazy } from 'react';
import Loading from '../../components/Loading/Loading';
import { useAuthenticatedUser } from '../../hooks/useAuthenticatedUser';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

const LoginForm = lazy(() => import('../../components/Forms/Login/Login'));

export function Login() {
  useDocumentTitle('Login');
  const isLoading = useAuthenticatedUser();

  return (
    <>
      {isLoading ? (
        <Loading />
      ) : (
        <Suspense fallback={<Loading />}>
          <LoginForm />
        </Suspense>
      )}
    </>
  );
}
