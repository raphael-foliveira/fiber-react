package apperror

type HTTPError struct {
	Code    int
	Message string
}

func (httpErr HTTPError) Error() string {
	return httpErr.Message
}
