package liberrors

type IStorageError interface {
	Error() string
	Unwrap() error
}