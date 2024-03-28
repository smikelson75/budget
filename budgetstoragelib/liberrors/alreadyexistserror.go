package liberrors

type AlreadyExistsError struct {
	err  error
	Type string
	Message string
}

func NewAlreadyExistsError(err error, domain string) *AlreadyExistsError {
	return &AlreadyExistsError{err, domain, "The requested resource already exists."}
}

func (e *AlreadyExistsError) Error() string {
	return e.Message
}

func (e *AlreadyExistsError) Unwrap() error {
	return e.err
}