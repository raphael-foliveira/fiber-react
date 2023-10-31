import { Suspense, lazy, useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Loading from '../../components/Loading/Loading';
import { AuthContext } from '../../contexts/authContext';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import { authService } from '../../service/authService';

const LoginForm = lazy(() => import('../../components/Forms/Login/Login'));

export function Login() {
  useDocumentTitle('Login');
  const { authData, setAuthData, clearAuthData } = useContext(AuthContext);
  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    const handleAuthenticatedUser = async () => {
      const { refreshToken, user, accessToken } = authData;
      try {
        if (refreshToken && user.id) {
          const { accessToken } = await authService.refreshToken({
            refreshToken,
            userId: user.id,
          });
          setAuthData({
            ...authData,
            accessToken,
          });
          navigate('/todos');
        }
      } catch {
        authService.logout({ accessToken });
        clearAuthData();
      } finally {
        setIsLoading(false);
      }
    };
    handleAuthenticatedUser();
  }, []);

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
