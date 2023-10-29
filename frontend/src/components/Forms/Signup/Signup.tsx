import { Button, TextField, Typography } from '@mui/material';
import { FormEventHandler, useContext, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../../../contexts/authContext';
import { ValidationError } from '../../../errors/ValidationError';
import { authService } from '../../../service/authService';
import { ButtonWrapper, FieldWrapper, FormCard } from '../styles';

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
      navigate('/todos');
    } catch (err) {
      console.log({ err });
      setFormError(true);
      if (err instanceof ValidationError) {
        setFormErrorMessage(err.message);
        return;
      }
      setFormErrorMessage('Erro ao cadastrar. Tente novamente.');
    }
  };

  return (
    <form action='' onSubmit={handleSubmit}>
      <FormCard>
        <Typography variant='h4'>Cadastro</Typography>
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
            sx={{
              color: 'red',
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
