import { Button, TextField, Typography } from '@mui/material';
import { FormEventHandler, useContext, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../../../contexts/authContext';
import { HttpError } from '../../../errors/HttpError';
import { ValidationError } from '../../../errors/ValidationError';
import { authService } from '../../../service/authService';
import { FormCard } from '../FormCard';
import { ButtonWrapper, FieldWrapper } from '../styles';

const errMessages: Record<string, string> = {
  email: 'Esse email já está cadastrado',
  username: 'Esse nome de usuário já está cadastrado',
};

export default function SignupForm() {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [formError, setFormError] = useState(false);
  const [formErrorMessage, setFormErrorMessage] = useState('');
  const { setAuthData } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setFormError(false);
    setFormErrorMessage('');
    try {
      const authResponse = await authService.signup({
        email,
        username,
        password,
        confirm_password: confirmPassword,
      });
      setAuthData(authResponse);
      navigate('/login');
    } catch (err) {
      console.error({ err });
      setFormError(true);
      if (err instanceof ValidationError) {
        setFormErrorMessage(err.message);
        return;
      }
      if (err instanceof HttpError && err.status === 409) {
        const errJson = err.json as { field?: string };
        setFormErrorMessage(errMessages[errJson.field || '']);
        return;
      }
      setFormErrorMessage('Erro ao cadastrar. Tente novamente.');
    }
  };

  return (
    <form action='' onSubmit={handleSubmit}>
      <FormCard>
        <Typography variant='h4' sx={{ textAlign: 'center', marginBottom: 4 }}>
          Cadastro
        </Typography>
        <FieldWrapper>
          <TextField
            label='Nome de usuário'
            name='username'
            id='username'
            type='text'
            value={username}
            error={formError}
            onChange={(e) => setUsername(e.target.value)}
          />
        </FieldWrapper>
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
        <FieldWrapper>
          <TextField
            label='Confirmar Senha'
            name='confirmPassword'
            id='confirmPassword'
            type='password'
            value={confirmPassword}
            error={formError}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
        </FieldWrapper>
        {formError && (
          <Typography
            variant='subtitle1'
            color='error'
            sx={{
              textAlign: 'center',
              width: '100%',
            }}
          >
            {formErrorMessage}
          </Typography>
        )}
        <ButtonWrapper>
          <Button variant='contained' type='submit'>
            Cadastrar
          </Button>
        </ButtonWrapper>
      </FormCard>
    </form>
  );
}
