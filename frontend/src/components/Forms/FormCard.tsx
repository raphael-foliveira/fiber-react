import { Card } from '@mui/material';
import { blue } from '@mui/material/colors';
import { ReactNode } from 'react';

export function FormCard({ children }: { children: ReactNode }) {
  return (
    <Card
      sx={{
        maxWidth: 400,
        padding: 8,
        margin: 'auto',
        marginTop: 10,
        backgroundColor: blue[50],
      }}
    >
      {children}
    </Card>
  );
}
