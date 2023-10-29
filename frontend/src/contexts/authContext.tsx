import { createContext, useEffect, useState } from 'react';
import { AuthContextProps, AuthData } from '../types/auth';

export const AuthContext = createContext<AuthContextProps>({
  authData: { isLoggedIn: false },
  setAuthData: () => {},
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [authData, setAuthData] = useState<AuthData>({ isLoggedIn: false });
  useEffect(() => {
    const userString = localStorage.getItem('user');
    const accessToken = localStorage.getItem('accessToken');
    const refreshToken = localStorage.getItem('refreshToken');
    if (userString && accessToken && refreshToken) {
      const user = JSON.parse(userString);
      setAuthData({ user, accessToken, refreshToken, isLoggedIn: true });
      return;
    }
    setAuthData({ isLoggedIn: false });
  }, []);

  return (
    <AuthContext.Provider value={{ authData, setAuthData }}>
      {children}
    </AuthContext.Provider>
  );
}
