import CreateTodoForm from '../../../components/Forms/Todos/CreateTodo';
import { useSession } from '../../../hooks/useSession';

export function CreateTodo() {
  const { accessToken } = useSession();
  return <CreateTodoForm accessToken={accessToken} />;
}
