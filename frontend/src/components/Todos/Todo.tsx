import { Box, Button, Card, Checkbox, Typography } from '@mui/material';
import { green, red } from '@mui/material/colors';
import { useState } from 'react';
import { todosService } from '../../service/todosService';
import { Todo } from '../../types/todos';
import UpdateTodo from '../Forms/Todos/UpdateTodo';
import { useMutation, useQueryClient } from '@tanstack/react-query';

interface SingleTodoProps {
  todo: Todo;
  accessToken: string;
}

export function SingleTodo({ todo, accessToken }: SingleTodoProps) {
  const queryClient = useQueryClient();
  const [isEditing, setIsEditing] = useState(false);

  const completeMutation = useMutation({
    mutationFn: async () => {
      await todosService.updateTodo(accessToken, {
        ...todo,
        completed: !todo.completed,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
    },
  });

  const deleteMutation = useMutation({
    mutationFn: async () => {
      await todosService.deleteTodo(accessToken, todo.id);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
    },
  });

  return (
    <Card
      sx={{
        padding: '20px',
        width: 1,
        maxWidth: '400px',
        height: '400px',
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: todo.completed ? green[100] : red[100],
      }}
    >
      {isEditing ? (
        <UpdateTodo todo={todo} setIsEditing={setIsEditing} />
      ) : (
        <>
          <Typography variant='h4' textAlign={'center'}>
            {todo.title}
          </Typography>
          <Box padding={4}>
            <Typography>{todo.description}</Typography>
          </Box>
          <Typography textAlign={'right'} marginTop={'auto'}>
            Feito:{' '}
            <Checkbox
              checked={todo.completed}
              onClick={() => {
                completeMutation.mutate();
              }}
            />
          </Typography>
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'space-between',
            }}
          >
            <Button onClick={() => setIsEditing(true)}>Editar</Button>
            <Button
              color='error'
              onClick={() => {
                deleteMutation.mutate();
              }}
            >
              Apagar
            </Button>
          </Box>
        </>
      )}
    </Card>
  );
}
