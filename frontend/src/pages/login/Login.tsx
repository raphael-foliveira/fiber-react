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
      try {
        const { refreshToken, user } = authData;
        if (refreshToken && user) {
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
        setIsLoading(false);
      } catch {
        authService.logout();
      } finally {
        clearAuthData();
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
