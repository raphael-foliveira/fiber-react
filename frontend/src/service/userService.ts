import { apiClient } from '../clients/apiClient';

export const todoService = {
  getUserInfo: async () => {
    apiClient.get('/users');
  },
};
