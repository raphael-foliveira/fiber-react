import { Box } from '@mui/material';
import Navbar from '../../components/Navbar/Navbar';
import { Outlet } from 'react-router-dom';

export function Root() {
  return (
    <Box sx={{ backgroundColor: 'background.paper', height: '100vh' }}>
      <Navbar />
      <Outlet />
    </Box>
  );
}
