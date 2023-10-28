import { apiClient } from '../clients/apiClient';
import { ValidationError } from '../errors/ValidationError';
import { AuthData, LoginProps, SignupProps } from '../types/auth';

export const authService = {
  login: async (credentials: LoginProps): Promise<AuthData> => {
    const loginResponse = await apiClient.post('/auth/login', credentials);
    storeAuthData(loginResponse);
    return loginResponse;
  },

  signup: async (credentials: SignupProps): Promise<AuthData> => {
    validateSignupCredentials(credentials);
    const signupResponse = await apiClient.post('/auth/signup', credentials);
    storeAuthData(signupResponse);
    return signupResponse;
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

function storeAuthData(authData: AuthData) {
  localStorage.setItem('user', JSON.stringify(authData.user));
  localStorage.setItem('accessToken', authData.accessToken);
  localStorage.setItem('refreshToken', authData.refreshToken);
}
