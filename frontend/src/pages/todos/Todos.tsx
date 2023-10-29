import TodosList from '../../components/Todos/TodosList';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

export function Todos() {
  useDocumentTitle('Todos');
  return <TodosList />;
}
