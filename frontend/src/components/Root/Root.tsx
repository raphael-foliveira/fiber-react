import { Box } from '@mui/material';
import Navbar from '../Navbar/Navbar';
import { Outlet } from 'react-router-dom';

export function Root() {
  return (
    <Box>
      <Navbar />
      <Outlet />
    </Box>
  );
}
