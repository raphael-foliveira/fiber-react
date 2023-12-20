import { Button, TextField, Typography } from '@mui/material';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { Dispatch, SetStateAction, useState } from 'react';
import { useSession } from '../../../hooks/useSession';
import { todosService } from '../../../service/todosService';
import { Todo } from '../../../types/todos';
import { ButtonWrapper, FieldWrapper } from '../styles';

interface UpdateTodoProps {
  todo: Todo;
  setIsEditing: Dispatch<SetStateAction<boolean>>;
}

export default function UpdateTodo({ todo, setIsEditing }: UpdateTodoProps) {
  const queryClient = useQueryClient();
  const { accessToken } = useSession();
  const [title, setTitle] = useState(todo.title);
  const [description, setDescription] = useState(todo.description);
  const [formError, setFormError] = useState(false);
  const [formErrorMessage, setFormErrorMessage] = useState('');

  const handleSubmit = async () => {
    try {
      await todosService.updateTodo(accessToken, {
        ...todo,
        title,
        description,
      });
      setIsEditing(false);
    } catch (err) {
      setFormError(true);
      if (err instanceof Error) {
        setFormErrorMessage(err.message);
        return;
      }
      setFormErrorMessage('Erro desconhecido');
    }
  };

  const mutation = useMutation({
    mutationFn: handleSubmit,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
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
      <Typography variant='h4' sx={{ marginBottom: 4 }}>
        Atualizar tarefa
      </Typography>
      <FieldWrapper>
        <TextField
          variant='standard'
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
          variant='standard'
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
        <Button type='submit'>Atualizar</Button>
        <Button type='button' onClick={() => setIsEditing(false)}>
          Cancelar
        </Button>
      </ButtonWrapper>
    </form>
  );
}
