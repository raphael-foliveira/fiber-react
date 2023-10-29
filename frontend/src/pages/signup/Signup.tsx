import { SignupForm } from '../../components/Forms/Signup/Signup';
import { useDocumentTitle } from '../../hooks/useDocumentTitle';

export function Signup() {
  useDocumentTitle('Signup');
  return <SignupForm />;
}
