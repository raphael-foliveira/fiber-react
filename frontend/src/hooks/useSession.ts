import { useContext, useEffect } from 'react';
import { AuthContext } from '../contexts/authContext';
import { useNavigate } from 'react-router-dom';

export function useSession() {
  const { authData } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (!authData.isLoggedIn) {
      navigate('/login');
    }
  }, [navigate, authData]);

  return authData;
}
