import { Box, Button, Card, Checkbox, Typography } from '@mui/material';
import { useState } from 'react';
import { todosService } from '../../service/todosService';
import { Todo } from '../../types/todos';
import UpdateTodo from '../Forms/Todos/UpdateTodo';

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
  const [todoState, setTodoState] = useState(todo);
  const [isEditing, setIsEditing] = useState(false);

  const handleCompleteTodo = async () => {
    const response = await todosService.updateTodo(accessToken, {
      ...todoState,
      completed: !todoState.completed,
    });
    setTodoState({ ...todoState, completed: response.completed });
  };

  const handleDeleteTodo = async () => {
    await todosService.deleteTodo(accessToken, todoState.id);
    updateTodos(accessToken);
  };

  return (
    <Card
      sx={{
        padding: '20px',
        width: 1,
        maxWidth: '400px',
        height: '400px',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      {isEditing ? (
        <UpdateTodo
          todo={todo}
          setTodoState={setTodoState}
          setIsEditing={setIsEditing}
        />
      ) : (
        <>
          <Typography variant='h4' textAlign={'center'}>
            {todoState.title}
          </Typography>
          <Box padding={4}>
            <Typography>{todoState.description}</Typography>
          </Box>
          <Typography textAlign={'right'} marginTop={'auto'}>
            Feito:{' '}
            <Checkbox
              checked={todoState.completed}
              onClick={handleCompleteTodo}
            />
          </Typography>
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'space-between',
            }}
          >
            <Button onClick={() => setIsEditing(true)}>Editar</Button>
            <Button color='error' onClick={handleDeleteTodo}>
              Apagar
            </Button>
          </Box>
        </>
      )}
    </Card>
  );
}
