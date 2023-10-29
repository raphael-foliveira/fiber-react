import { Suspense, lazy } from 'react';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import Loading from '../../components/Loading/Loading';

const SignupForm = lazy(() => import('../../components/Forms/Signup/Signup'));

export function Signup() {
  useDocumentTitle('Signup');
  return (
    <Suspense fallback={<Loading />}>
      <SignupForm />
    </Suspense>
  );
}
