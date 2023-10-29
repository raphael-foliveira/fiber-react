import { Box, CircularProgress } from '@mui/material';

export default function Loading() {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', marginTop: 10 }}>
      <CircularProgress />
    </Box>
  );
}
