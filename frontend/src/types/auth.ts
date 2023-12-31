import { Dispatch, SetStateAction } from 'react';

export interface AuthData {
  user: AuthDataUser;
  accessToken: string;
  refreshToken: string;
  isLoggedIn: boolean;
}
export interface AuthDataUser {
  id?: number;
  email?: string;
  username?: string;
}

export interface LoginProps {
  email: string;
  password: string;
}

export interface SignupProps extends LoginProps {
  username: string;
  confirm_password: string;
}

export interface AuthContextProps {
  authData: AuthData;
  setAuthData: Dispatch<SetStateAction<AuthData>>;
  clearAuthData: () => void;
}

export interface StoreAuthDataProps {
  user: {
    id: string;
    email: string;
  };
  access_token: string;
  refresh_token: string;
}
