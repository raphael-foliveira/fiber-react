import { AppBar, Button, Toolbar, Typography } from '@mui/material';
import { Link } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../../contexts/authContext';

export default function Navbar() {
  const { authData, setAuthData } = useContext(AuthContext);

  const handleLogout = () => {
    setAuthData({ isLoggedIn: false });
    localStorage.removeItem('user');
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
  };

  return (
    <AppBar position='static'>
      <Toolbar>
        <Link to='/'>
          <Typography variant='h6' sx={{ cursor: 'pointer' }}>
            Home
          </Typography>
        </Link>
        {authData.isLoggedIn && (
          <Button
            onClick={handleLogout}
            sx={{ color: 'inherit', marginLeft: 'auto' }}
          >
            Logout
          </Button>
        )}
      </Toolbar>
    </AppBar>
  );
}
