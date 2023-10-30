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

type ConflictError struct {
	Message string
	Field   string
}

func (conflictErr ConflictError) Error() string { return conflictErr.Message }
