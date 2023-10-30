import { useEffect, useState } from 'react';
import { useSession } from '../../hooks/useSession';
import { todosService } from '../../service/todosService';
import { Todo } from '../../types/todos';
import { SingleTodo } from './Todo';
import { Box, Button, Container, Typography } from '@mui/material';
import { Link } from 'react-router-dom';

export default function TodosList() {
  const { accessToken } = useSession();
  const [todos, setTodos] = useState<Todo[]>([]);
  const fetchTodos = async (accessToken: string) => {
    const todos = await todosService.getTodos(accessToken);
    setTodos(
      todos.sort(
        (a, b) =>
          new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
      )
    );
  };

  useEffect(() => {
    fetchTodos(accessToken!);
  }, []);

  return (
    <>
      <Typography
        textAlign='center'
        variant='h3'
        fontWeight='500'
        margin='2rem 0'
      >
        Tarefas
      </Typography>
      <Box sx={{ textAlign: 'center', marginBottom: 4 }}>
        <Link to='/todos/create'>
          <Button>Criar nova tarefa</Button>
        </Link>
      </Box>
      <Box sx={{ display: 'flex', flexWrap: 'wrap', padding: '0 5rem' }}>
        {todos.map((todo) => (
          <Container
            key={todo.id}
            sx={{
              marginBottom: 5,
              width: {
                xs: 1,
                lg: 1 / 3,
              },
              display: 'flex',
              flexWrap: 'wrap',
              justifyContent: 'center',
            }}
          >
            <SingleTodo
              todo={todo}
              updateTodos={fetchTodos}
              accessToken={accessToken}
            />
          </Container>
        ))}
      </Box>
    </>
  );
}
