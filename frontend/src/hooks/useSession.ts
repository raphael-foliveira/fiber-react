import { useContext, useEffect } from 'react';
import { AuthContext } from '../contexts/authContext';
import { useNavigate } from 'react-router-dom';

export function useSession() {
  const authContext = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (!authContext?.authData) {
      navigate('/login');
    }
  }, [navigate, authContext]);

  return authContext;
}
