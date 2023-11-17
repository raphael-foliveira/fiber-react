import { Box, Button, Container, Typography } from '@mui/material';
import { useQuery } from '@tanstack/react-query';
import { Link } from 'react-router-dom';
import { useSession } from '../../hooks/useSession';
import { todosService } from '../../service/todosService';
import { SingleTodo } from './Todo';
import Loading from '../Loading/Loading';

export default function TodosList() {
  const authData = useSession();
  const fetchTodos = async () => {
    return todosService.getUserTodos(authData.accessToken, authData.user.id);
  };

  const query = useQuery({
    queryKey: ['todos'],
    queryFn: fetchTodos,
  });
  if (query.error) {
    return <h1>Something went wrong</h1>;
  }

  if (query.isLoading) {
    return <Loading />;
  }

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
        {query.data?.map((todo) => (
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
