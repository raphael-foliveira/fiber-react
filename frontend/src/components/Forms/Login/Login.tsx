import { Button, TextField, Typography } from '@mui/material';
import { ButtonWrapper, FieldWrapper, FormCard } from '../styles';
import { FormEventHandler, useState } from 'react';
import { Link } from 'react-router-dom';

export function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [formError, setFormError] = useState(false);
  const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();
    setFormError(true);
  };

  return (
    <form action='' onSubmit={handleSubmit}>
      <FormCard>
        <Typography variant='h4'>Login</Typography>
        <FieldWrapper>
          <TextField
            label='Email'
            name='email'
            id='email'
            type='email'
            value={email}
            error={formError}
            onChange={(e) => setEmail(e.target.value)}
          />
        </FieldWrapper>
        <FieldWrapper>
          <TextField
            label='Senha'
            name='password'
            id='password'
            type='password'
            value={password}
            error={formError}
            onChange={(e) => setPassword(e.target.value)}
          />
        </FieldWrapper>
        <ButtonWrapper>
          <Button variant='contained' type='submit'>
            Login
          </Button>
          <Link to='/signup'>
            <Button variant='contained' type='button'>
              Cadastre-se
            </Button>
          </Link>
        </ButtonWrapper>
        {formError && (
          <Typography
            variant='subtitle1'
            sx={{
              color: 'red',
              textAlign: 'center',
              width: '100%',
            }}
          >
            Credenciais inválidas. Tente novamente.
          </Typography>
        )}
      </FormCard>
    </form>
  );
}
