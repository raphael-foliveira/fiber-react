import { apiClient } from '../clients/apiClient';
import { HttpError } from '../errors/HttpError';
import { Todo, TodoProps } from '../types/todos';
import { authService } from './authService';

export const todosService = {
  createTodo: async (
    todo: TodoProps,
    accessToken: string = ''
  ): Promise<TodoProps> => {
    try {
      const response = await apiClient.post('/todos', todo, {
        headers: { Authorization: accessToken },
      });
      return response;
    } catch (err) {
      if (err instanceof HttpError) {
        if (err.status === 401) {
          throw new Error('Não autorizado');
        }
        if (err.status === 400 || err.status === 422) {
          throw new Error('Dados inválidos');
        }
      }
      throw new Error('Erro inesperado, tente novamente.');
    }
  },

  getTodos: async (accessToken: string = ''): Promise<Todo[]> => {
    try {
      const response = await apiClient.get('/todos', {
        headers: { Authorization: accessToken },
      });
      return response;
    } catch (err) {
      if (err instanceof HttpError) {
        if (err.status === 401) {
          authService.logout();
          throw new Error('Não autorizado');
        }
      }
      throw new Error('Erro inesperado, tente novamente.');
    }
  },
};
