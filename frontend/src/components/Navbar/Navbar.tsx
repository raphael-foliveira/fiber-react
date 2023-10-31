import { AppBar, Button, Toolbar, Typography } from '@mui/material';
import { Link } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../../contexts/authContext';
import { authService } from '../../service/authService';

export default function Navbar() {
  const { authData, clearAuthData } = useContext(AuthContext);

  const handleLogout = () => {
    authService.logout({
      accessToken: authData.accessToken,
    });
    clearAuthData();
  };

  return (
    <AppBar position='static'>
      <Toolbar>
        <Link to='/'>
          <Typography variant='h6' sx={{ cursor: 'pointer', margin: '0 2rem' }}>
            Home
          </Typography>
        </Link>
        {authData.isLoggedIn && (
          <>
            <Link to='/todos'>
              <Button sx={{ color: 'inherit', marginRight: 'auto' }}>
                Tarefas
              </Button>
            </Link>

            <Button
              onClick={handleLogout}
              sx={{ color: 'inherit', marginLeft: 'auto' }}
            >
              Logout
            </Button>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
}
