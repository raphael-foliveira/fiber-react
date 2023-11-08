import { useContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../contexts/authContext';
import { authService } from '../service/authService';
import { AuthData } from '../types/auth';

export function useAuthenticatedUser() {
  const { authData, setAuthData, clearAuthData } = useContext(AuthContext);
  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

  const handleAuthenticatedUser = async ({
    refreshToken,
    user,
    accessToken,
  }: AuthData) => {
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

  useEffect(() => {
    handleAuthenticatedUser(authData);
  }, []);

  return isLoading;
}
