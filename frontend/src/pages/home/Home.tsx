import { Button } from '@mui/material';
import { HomeButtonsContainer, HomeContainer } from './styles';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

export function Home() {
  useDocumentTitle('Home');
  return (
    <HomeContainer>
      <h1>Bem vindo</h1>
      <HomeButtonsContainer>
        <a href='/login'>
          <Button variant='contained'>Login</Button>
        </a>
        <a href=''>
          <Button variant='contained'>Sobre</Button>
        </a>
      </HomeButtonsContainer>
    </HomeContainer>
  );
}
