package errs

type HTTPError struct {
	Code    int
	Message string
}

func (httpErr HTTPError) Error() string { return httpErr.Message }

type NotFoundError struct {
	Message string
}

func (notFoundErr NotFoundError) Error() string { return notFoundErr.Message }
