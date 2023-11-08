import { useEffect, useState } from 'react';
import { todosService } from '../service/todosService';
import { AuthData } from '../types/auth';
import { Todo } from '../types/todos';

export function useTodos({ authData }: { authData: AuthData }) {
  const { user, accessToken } = authData;
  const [todos, setTodos] = useState<Todo[]>([]);

  const fetchTodos = async (accessToken: string) => {
    const todos = await todosService.getUserTodos(accessToken, user?.id);
    setTodos(
      todos.sort(
        (a, b) =>
          new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      )
    );
  };

  useEffect(() => {
    fetchTodos(accessToken!);
  }, []);

  return { todos, fetchTodos };
}
