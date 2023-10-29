import { Button, TextField, Typography } from '@mui/material';
import { FormEventHandler, useState } from 'react';
import { useSession } from '../../../hooks/useSession';
import { todosService } from '../../../service/todosService';
import { ButtonWrapper, FieldWrapper, FormCard } from '../styles';

export default function CreateTodoForm() {
  const { accessToken } = useSession();
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [formError, setFormError] = useState(false);
  const [formErrorMessage, setFormErrorMessage] = useState('');

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setFormError(false);
    setFormErrorMessage('');
    try {
      await todosService.createTodo({ title, description }, accessToken);
    } catch (err) {
      setFormError(true);
      if (err instanceof Error) {
        setFormErrorMessage(err.message);
        return;
      }
      setFormErrorMessage('Erro inesperado. Tente novamente.');
    }
  };

  return (
    <form action='' onSubmit={handleSubmit}>
      <FormCard>
        <Typography variant='h4'>Cadastro</Typography>
        <FieldWrapper>
          <TextField
            label='Título'
            name='title'
            id='title'
            type='text'
            value={title}
            error={formError}
            onChange={(e) => setTitle(e.target.value)}
          />
        </FieldWrapper>
        <FieldWrapper>
          <TextField
            label='Descrição'
            name='description'
            id='description'
            type='text'
            value={description}
            error={formError}
            onChange={(e) => setDescription(e.target.value)}
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
