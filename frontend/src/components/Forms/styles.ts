import { Card } from '@mui/material';
import styled from 'styled-components';

export const FieldWrapper = styled.div`
  width: 100%;
  display: flex;
  justify-content: space-around;
  margin-bottom: 30px;

  .MuiFormControl-root {
    width: 100%;
  }
`;

export const ButtonWrapper = styled.div`
  width: 100%;
  display: flex;
  justify-content: space-around;
  margin-bottom: 30px;

  .MuiButtonBase-root {
    width: 150px;
  }
`;

export const FormCard = styled(Card)`
  display: flex;
  max-width: 400px;
  margin: 100px auto;
  padding: 20px;
  flex-wrap: wrap;

  h4 {
    text-align: center;
    width: 100%;
    margin-bottom: 20px;
  }
`;
