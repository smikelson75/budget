package liberrors

type DefaultError struct {
	err  error
	Type string
	Message string
}

func NewDefaultError(err error) *DefaultError {
	return &DefaultError{err, "Internal", "An error occured, unwrap error for more details."}
}

func (e *DefaultError) Error() string {
	return e.Message
}

func (e *DefaultError) Unwrap() error {
	return e.err
}