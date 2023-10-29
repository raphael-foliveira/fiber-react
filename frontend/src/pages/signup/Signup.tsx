import { Suspense, lazy, useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Loading from '../../components/Loading/Loading';
import { AuthContext } from '../../contexts/authContext';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

const SignupForm = lazy(() => import('../../components/Forms/Signup/Signup'));

export function Signup() {
  useDocumentTitle('Signup');
  const { authData } = useContext(AuthContext);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    if (authData.isLoggedIn) {
      navigate('/todos');
    }
    setLoading(false);
  }, []);
  return loading ? (
    <Loading />
  ) : (
    <Suspense fallback={<Loading />}>
      <SignupForm />
    </Suspense>
  );
}
