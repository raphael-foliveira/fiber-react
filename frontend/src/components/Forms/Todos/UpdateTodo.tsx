import { Typography, TextField, Button } from '@mui/material';
import { useState, FormEventHandler, Dispatch, SetStateAction } from 'react';
import { useSession } from '../../../hooks/useSession';
import { todosService } from '../../../service/todosService';
import { Todo } from '../../../types/todos';
import { FieldWrapper, ButtonWrapper } from '../styles';

interface UpdateTodoProps {
  todo: Todo;
  setTodoState: Dispatch<SetStateAction<Todo>>;
  setIsEditing: Dispatch<SetStateAction<boolean>>;
}

export default function UpdateTodo({
  todo,
  setTodoState,
  setIsEditing,
}: UpdateTodoProps) {
  const { accessToken } = useSession();
  const [title, setTitle] = useState(todo.title);
  const [description, setDescription] = useState(todo.description);
  const [formError, setFormError] = useState(false);
  const [formErrorMessage, setFormErrorMessage] = useState('');

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    try {
      const response = await todosService.updateTodo(accessToken, {
        ...todo,
        title,
        description,
      });
      setTodoState(response);
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

  return (
    <form action='' onSubmit={handleSubmit}>
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
