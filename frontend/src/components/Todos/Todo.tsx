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
  updateTodos: (accessToken: string) => void;
}

export function SingleTodo({
  todo,
  updateTodos,
  accessToken,
}: SingleTodoProps) {
  const queryClient = useQueryClient();
  const [isEditing, setIsEditing] = useState(false);

  const handleCompleteTodo = async () => {
    await todosService.updateTodo(accessToken, {
      ...todo,
      completed: !todo.completed,
    });
  };

  const completeMutation = useMutation({
    mutationFn: handleCompleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
    },
  });

  const handleDeleteTodo = async () => {
    await todosService.deleteTodo(accessToken, todo.id);
    updateTodos(accessToken);
  };

  const deleteMutation = useMutation({
    mutationFn: handleDeleteTodo,
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
