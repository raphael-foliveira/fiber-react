import { Button, Typography } from '@mui/material';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import { HomeButtonsContainer, HomeContainer } from './styles';
import { Link } from 'react-router-dom';

export function Home() {
  useDocumentTitle('Home');
  return (
    <HomeContainer>
      <Typography variant='h3' sx={{ fontWeight: 'bold', marginBottom: 4 }}>
        Bem vindo
      </Typography>
      <HomeButtonsContainer>
        <Link to='/login'>
          <Button variant='contained'>Login</Button>
        </Link>
        <Link to='/signup'>
          <Button variant='contained'>Cadastre-se</Button>
        </Link>
      </HomeButtonsContainer>
    </HomeContainer>
  );
}
