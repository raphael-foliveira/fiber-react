import { Button, Card, TextField, Typography } from '@mui/material';
import { ButtonWrapper, FieldWrapper } from '../styles';

export function LoginForm() {
  return (
    <Card
      sx={{
        display: 'flex',
        maxWidth: '400px',
        margin: '100px auto',
        padding: '20px',
        flexWrap: 'wrap',
      }}
    >
      <Typography
        variant='h4'
        sx={{
          textAlign: 'center',
          width: '100%',
          marginBottom: '20px',
        }}
      >
        Login
      </Typography>
      <FieldWrapper>
        <TextField label='Email' />
      </FieldWrapper>
      <FieldWrapper>
        <TextField label='Senha' />
      </FieldWrapper>
      <ButtonWrapper>
        <Button variant='contained'>Login</Button>
        <Button variant='contained'>Cadastre-se</Button>
      </ButtonWrapper>
    </Card>
  );
}
