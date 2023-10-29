import Navbar from '../../components/Navbar/Navbar';
import { Outlet } from 'react-router-dom';

export function Root() {
  return (
    <>
      <Navbar />
      <Outlet />
    </>
  );
}
