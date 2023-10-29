import { RouterProvider } from 'react-router-dom';
import { router } from './routes/routes';
import { AuthProvider } from './contexts/authContext';

function App() {
  return (
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  );
}

export default App;
