import { Button, TextField, Typography } from '@mui/material';
import { useMutation } from '@tanstack/react-query';
import { useContext, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { AuthContext } from '../../../contexts/authContext';
import { authService } from '../../../service/authService';
import { FormCard } from '../FormCard';
import { ButtonWrapper, FieldWrapper } from '../styles';

export default function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [formError, setFormError] = useState(false);
  const { setAuthData } = useContext(AuthContext);
  const navigate = useNavigate();
  const mutation = useMutation({
    mutationFn: async () => {
      try {
        const loginResponse = await authService.login({ email, password });
        setAuthData(loginResponse);
        navigate('/todos');
      } catch (e) {
        setFormError(true);
        return;
      }
    },
  });

  return (
    <form
      action=''
      onSubmit={(e) => {
        e.preventDefault();
        mutation.mutate();
      }}
    >
      <FormCard>
        <Typography variant='h4' sx={{ textAlign: 'center', marginBottom: 4 }}>
          Login
        </Typography>
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
            color='error'
            sx={{
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
