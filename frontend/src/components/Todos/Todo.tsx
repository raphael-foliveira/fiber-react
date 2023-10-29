import { Card, Checkbox, Typography } from '@mui/material';
import { Todo } from '../../types/todos';
import { useState } from 'react';

export function SingleTodo({ todo }: { todo: Todo }) {
  const [todoState, setTodoState] = useState(todo);
  const [isEditing, setIsEditing] = useState(false);

  return (
    <Card sx={{ padding: '10px', textAlign: 'center' }}>
      <Typography variant='h4'>{todoState.title}</Typography>
      <Typography>{todoState.description}</Typography>
      <Typography>
        Feito: <Checkbox checked={todoState.completed} />
      </Typography>
    </Card>
  );
}
