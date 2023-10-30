import { Box } from '@mui/material';
import styled from 'styled-components';

export const FieldWrapper = styled(Box)`
  width: 100%;
  display: flex;
  justify-content: space-around;
  margin-bottom: 30px;

  .MuiFormControl-root {
    width: 100%;
  }
`;

export const ButtonWrapper = styled(Box)`
  width: 100%;
  display: flex;
  justify-content: space-around;

  .MuiButtonBase-root {
    width: 150px;
  }
`;
