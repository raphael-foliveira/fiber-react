import { Button } from '@mui/material';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';
import { HomeButtonsContainer, HomeContainer } from './styles';
import { Link } from 'react-router-dom';

export function Home() {
  useDocumentTitle('Home');
  return (
    <HomeContainer>
      <h1>Bem vindo</h1>
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
