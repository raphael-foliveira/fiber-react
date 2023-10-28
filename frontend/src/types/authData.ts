export interface AuthData {
  user: {
    id: string;
    email: string;
  };
  accessToken: string;
  refreshToken: string;
}
