import { Suspense, lazy } from 'react';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import Loading from '../../components/Loading/Loading';

const TodosList = lazy(() => import('../../components/Todos/TodosList'));

export function Todos() {
  useDocumentTitle('Todos');
  return (
    <Suspense fallback={<Loading />}>
      <TodosList />;
    </Suspense>
  );
}
