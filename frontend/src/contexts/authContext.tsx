import {
  createContext,
  Dispatch,
  SetStateAction,
  useEffect,
  useState,
} from 'react';
import { AuthData } from '../types/authData';

interface AuthContextProps {
  authData: AuthData | undefined;
  setAuthData: Dispatch<SetStateAction<AuthData | undefined>>;
}

export const AuthContext = createContext<AuthContextProps>({
  authData: undefined,
  setAuthData: () => {},
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [authData, setAuthData] = useState<AuthData | undefined>(undefined);
  useEffect(() => {
    const userString = localStorage.getItem('user');
    const accessToken = localStorage.getItem('accessToken');
    const refreshToken = localStorage.getItem('refreshToken');
    if (userString && accessToken && refreshToken) {
      const user = JSON.parse(userString);
      setAuthData({ user, accessToken, refreshToken });
      return;
    }
    setAuthData(undefined);
  }, []);

  return (
    <AuthContext.Provider value={{ authData, setAuthData }}>
      {children}
    </AuthContext.Provider>
  );
}
