import { Dispatch, SetStateAction } from 'react';

export interface AuthData {
  user?: {
    id: string;
    email: string;
    username: string;
  };
  accessToken?: string;
  refreshToken?: string;
  isLoggedIn: boolean;
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
}

export interface StoreAuthDataProps {
  user: {
    id: string;
    email: string;
  };
  access_token: string;
  refresh_token: string;
}
