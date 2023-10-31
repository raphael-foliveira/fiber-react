import { Box, Button, Container, Typography } from '@mui/material';
import { Link } from 'react-router-dom';
import { useTodos } from '../../hooks/useTodos';
import { SingleTodo } from './Todo';
import { useSession } from '../../hooks/useSession';

export default function TodosList() {
  const authData = useSession();
  const { todos, fetchTodos } = useTodos({ authData });
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
              accessToken={authData.accessToken}
            />
          </Container>
        ))}
      </Box>
    </>
  );
}
