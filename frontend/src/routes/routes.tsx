import { createBrowserRouter } from 'react-router-dom';
import { Home } from '../pages/home/Home';
import { Login } from '../pages/login/Login';
import { Signup } from '../pages/signup/Signup';

export const router = createBrowserRouter([
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
]);
