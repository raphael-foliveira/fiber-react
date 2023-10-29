import { AppBar, IconButton, Toolbar, Typography } from '@mui/material';
import { Link } from 'react-router-dom';
import { useSession } from '../../hooks/useSession';
import { useEffect } from 'react';

export default function Navbar() {
  const authData = useSession();

  useEffect(() => {});
  return (
    <AppBar position='static'>
      <Toolbar>
        <IconButton
          size='large'
          edge='start'
          color='inherit'
          aria-label='menu'
          sx={{ mr: 2 }}
        ></IconButton>
        <Link to='/'>
          <Typography variant='h6' sx={{ cursor: 'pointer' }}>
            Home
          </Typography>
          {authData.isLoggedIn && <Typography variant='h6'>Logout</Typography>}
        </Link>
      </Toolbar>
    </AppBar>
  );
}
