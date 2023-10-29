import { createContext, useEffect, useState } from 'react';
import { AuthContextProps, AuthData } from '../types/auth';

export const AuthContext = createContext<AuthContextProps>({
  authData: { isLoggedIn: false, accessToken: '', refreshToken: '' },
  setAuthData: () => {},
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [authData, setAuthData] = useState<AuthData>({
    isLoggedIn: false,
    accessToken: '',
    refreshToken: '',
  });

  useEffect(() => {
    const userString = localStorage.getItem('user');
    const accessToken = localStorage.getItem('accessToken');
    const refreshToken = localStorage.getItem('refreshToken');
    try {
      if (userString && accessToken && refreshToken) {
        const user = JSON.parse(userString);
        setAuthData({ user, accessToken, refreshToken, isLoggedIn: true });
        return;
      }
      setAuthData({ isLoggedIn: false, accessToken: '', refreshToken: '' });
    } catch {
      setAuthData({ isLoggedIn: false, accessToken: '', refreshToken: '' });
    }
  }, []);

  return (
    <AuthContext.Provider value={{ authData, setAuthData }}>
      {children}
    </AuthContext.Provider>
  );
}
