import { apiClient } from '../clients/apiClient';
import { ValidationError } from '../errors/ValidationError';
import { AuthData } from '../types/authData';

interface LoginProps {
  email: string;
  password: string;
}

interface SignupProps extends LoginProps {
  username: string;
  confirm_password: string;
}

export const authService = {
  login: async (credentials: LoginProps): Promise<AuthData> => {
    return apiClient.post('/auth/login', credentials);
  },

  signup: async (credentials: SignupProps): Promise<AuthData> => {
    validateSignupCredentials(credentials);
    console.log('trying to hit api');
    return apiClient.post('/auth/signup', credentials);
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
