import { apiClient } from '../clients/apiClient';
import { TodoProps } from '../types/todos';

export const todosService = {
  createTodo: async (
    todo: TodoProps,
    accessToken: string = ''
  ): Promise<TodoProps> => {
    const response = await apiClient.post('/todos', todo, {
      headers: { Authorization: accessToken },
    });
    return response;
  },
};
