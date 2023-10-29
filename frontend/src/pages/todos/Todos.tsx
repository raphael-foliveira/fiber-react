import { Suspense, lazy } from 'react';
import Loading from '../../components/Loading/Loading';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

const TodosList = lazy(() => import('../../components/Todos/TodosList'));

export function Todos() {
  useDocumentTitle('Todos');
  return (
    <Suspense fallback={<Loading />}>
      <TodosList />
    </Suspense>
  );
}
