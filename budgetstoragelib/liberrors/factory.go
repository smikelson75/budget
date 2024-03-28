package liberrors

import "errors"

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) WrapError(err error, domain string) IStorageError {
	var alreadyExistsError *AlreadyExistsError

	if err != nil && errors.As(err, &alreadyExistsError) {
		return NewAlreadyExistsError(err, domain)
	}

	return NewDefaultError(err)
}
