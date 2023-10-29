import { LoginForm } from '../../components/Forms/Login/Login';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

export function Login() {
  useDocumentTitle('Login');
  return <LoginForm />;
}
