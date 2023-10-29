import { useEffect, useState } from 'react';
import { useSession } from '../../hooks/useSession';
import { todosService } from '../../service/todosService';
import { Todo } from '../../types/todos';
import { SingleTodo } from './Todo';
import { TodoBox } from './styles';

export default function TodosList() {
  const { accessToken } = useSession();
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    const fetchTodos = async () => {
      const todos = await todosService.getTodos(accessToken);
      setTodos(todos);
    };
    fetchTodos();
  }, [accessToken]);

  return todos.map((todo) => (
    <TodoBox key={todo.id}>
      <SingleTodo todo={todo} />
    </TodoBox>
  ));
}
