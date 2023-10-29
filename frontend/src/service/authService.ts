import { apiClient } from '../clients/apiClient';
import { ValidationError } from '../errors/ValidationError';
import {
  AuthData,
  LoginProps,
  SignupProps,
  StoreAuthDataProps,
} from '../types/auth';

export const authService = {
  login: async (credentials: LoginProps): Promise<AuthData> => {
    const { access_token, refresh_token, user } = await apiClient.post(
      '/auth/login',
      credentials
    );
    storeAuthData({ access_token, refresh_token, user });
    return {
      user: {
        id: user.id,
        email: user.email,
        username: user.username,
      },
      accessToken: access_token,
      refreshToken: refresh_token,
      isLoggedIn: true,
    };
  },

  signup: async (credentials: SignupProps): Promise<AuthData> => {
    validateSignupCredentials(credentials);
    const signupResponse = await apiClient.post('/auth/signup', credentials);
    storeAuthData(signupResponse);
    return signupResponse;
  },

  logout: () => {
    localStorage.removeItem('user');
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
  },
};

function validateSignupCredentials({
  password,
  confirm_password: confirmPassword,
  email,
  username,
}: SignupProps) {
  if (password !== confirmPassword) {
    throw new ValidationError('As senhas não coincidem');
  }
  if (password.length < 8) {
    throw new ValidationError('Senhas devem ter pelo menos 8 caracteres');
  }
  if (email.length < 6) {
    throw new ValidationError('Email inválido');
  }
  if (username.length < 4) {
    throw new ValidationError(
      'Nome de usuário deve ter pelo menos 4 caracteres'
    );
  }
}

function storeAuthData({
  user,
  access_token,
  refresh_token,
}: StoreAuthDataProps) {
  localStorage.setItem('user', JSON.stringify(user));
  if (!access_token || !refresh_token) {
    throw new Error('No auth data to store');
  }
  localStorage.setItem('accessToken', access_token);
  localStorage.setItem('refreshToken', refresh_token);
}
