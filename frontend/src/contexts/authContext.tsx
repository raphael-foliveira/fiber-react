import { createContext, useEffect, useState } from 'react';
import { AuthContextProps, AuthData } from '../types/auth';

export const AuthContext = createContext<AuthContextProps>({
  authData: {
    isLoggedIn: false,
    accessToken: '',
    refreshToken: '',
    user: {},
  },
  setAuthData: () => {},
  clearAuthData: () => {},
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [authData, setAuthData] = useState<AuthData>({
    isLoggedIn: false,
    accessToken: '',
    refreshToken: '',
    user: {},
  });

  const clearAuthData = () => {
    setAuthData({
      isLoggedIn: false,
      accessToken: '',
      refreshToken: '',
      user: {},
    });
    localStorage.removeItem('user');
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
  };

  useEffect(() => {
    const userString = localStorage.getItem('user');
    const accessToken = localStorage.getItem('accessToken');
    const refreshToken = localStorage.getItem('refreshToken');
    if (userString && accessToken && refreshToken) {
      try {
        const user = JSON.parse(userString);
        setAuthData({ user, accessToken, refreshToken, isLoggedIn: true });
        return;
      } finally {
        clearAuthData();
      }
    }
  }, []);

  return (
    <AuthContext.Provider value={{ authData, setAuthData, clearAuthData }}>
      {children}
    </AuthContext.Provider>
  );
}
