export interface AuthData {
  user: {
    id: string;
    email: string;
  };
  accessToken: string;
  refreshToken: string;
}

export interface LoginProps {
  email: string;
  password: string;
}

export interface SignupProps extends LoginProps {
  username: string;
  confirm_password: string;
}