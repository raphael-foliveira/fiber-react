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
        headers: { Authorization: 'Bearer ' + accessToken },
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

  getUserTodos: async (
    accessToken: string = '',
    userId: number = 0
  ): Promise<Todo[]> => {
    try {
      const response: Todo[] = await apiClient.get(`/users/${userId}/todos`, {
        headers: { Authorization: 'Bearer ' + accessToken },
      });
      return response.sort(
        (a, b) =>
          new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      );
    } catch (err) {
      console.error(err);
      if (err instanceof HttpError) {
        if (err.status === 401) {
          authService.logout({ accessToken });
          throw new Error('Não autorizado');
        }
      }
      throw new Error('Erro inesperado, tente novamente.');
    }
  },

  updateTodo: async (accessToken: string = '', todo: Todo): Promise<Todo> => {
    try {
      const { title, description, completed } = todo;
      const response = await apiClient.put(
        `/todos/${todo.id}`,
        { title, description, completed },
        {
          headers: { Authorization: 'Bearer ' + accessToken },
        }
      );
      return response;
    } catch (err) {
      if (err instanceof HttpError) {
        if (err.status === 401 || err.status === 403) {
          authService.logout({ accessToken });
          throw new Error('Não autorizado');
        }
        if (err.status === 400 || err.status === 422) {
          throw new Error('Dados inválidos');
        }
      }
      throw new Error('Erro inesperado, tente novamente.');
    }
  },

  deleteTodo: async (
    accessToken: string = '',
    todoId: number
  ): Promise<void> => {
    try {
      await apiClient.delete(`/todos/${todoId}`, {
        headers: { Authorization: 'Bearer ' + accessToken },
      });
    } catch (err) {
      if (err instanceof HttpError) {
        if (err.status === 401 || err.status === 403) {
          authService.logout({ accessToken });
          throw new Error('Não autorizado');
        }
      }
      throw new Error('Erro inesperado, tente novamente.');
    }
  },
};
