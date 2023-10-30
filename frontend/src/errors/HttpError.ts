export class HttpError {
  constructor(
    public message: string,
    public status: number,
    public json: unknown
  ) {}
}
