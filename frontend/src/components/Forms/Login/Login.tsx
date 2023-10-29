import { Button, TextField, Typography } from '@mui/material';
import { FormEventHandler, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { authService } from '../../../service/authService';
import { ButtonWrapper, FieldWrapper, FormCard } from '../styles';

export default function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [formError, setFormError] = useState(false);
  const navigate = useNavigate();

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (event) => {
    event.preventDefault();
    try {
      await authService.login({ email, password });
      navigate('/todos');
    } catch (e) {
      setFormError(true);
      return;
    }
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
