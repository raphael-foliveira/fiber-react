import { apiClient } from '../clients/apiClient';

interface LoginProps {
  username: string;
  password: string;
}

export const authService = {
  login: async (credentials: LoginProps) => {
    const response = await apiClient.post('/auth/login', credentials);
    return response.json();
  },
};
