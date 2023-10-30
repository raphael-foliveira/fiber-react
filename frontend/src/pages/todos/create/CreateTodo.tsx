import CreateTodoForm from '../../../components/Forms/Todos/CreateTodo';
import { useDocumentTitle } from '../../../hooks/useDocumentTitle';
import { useSession } from '../../../hooks/useSession';

export function CreateTodo() {
  useDocumentTitle('Create Todo');
  const { accessToken } = useSession();
  return <CreateTodoForm accessToken={accessToken} />;
}
