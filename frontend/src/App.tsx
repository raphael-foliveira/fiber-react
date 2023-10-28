import { RouterProvider } from 'react-router-dom';
import { router } from './routes/routes';
import Navbar from './components/Navbar/Navbar';

function App() {
  return (
    <>
      <Navbar />
      <RouterProvider router={router} />
    </>
  );
}

export default App;
