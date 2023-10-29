import { createBrowserRouter } from 'react-router-dom';
import { Home } from '../pages/home/Home';
import { Login } from '../pages/login/Login';
import { Root } from '../pages/root/Root';
import { Signup } from '../pages/signup/Signup';
import { Todos } from '../pages/todos/Todos';
import { CreateTodo } from '../pages/todos/create/CreateTodo';

export const router = createBrowserRouter([
  {
    path: '',
    element: <Root />,
    errorElement: <h1>404</h1>,
    children: [
      {
        path: '',
        element: <Home />,
      },
      {
        path: 'login',
        element: <Login />,
      },
      {
        path: 'signup',
        element: <Signup />,
      },
      {
        path: 'todos',
        element: <Todos />,
      },
      {
        path: 'todos/create',
        element: <CreateTodo />,
      },
    ],
  },
]);
