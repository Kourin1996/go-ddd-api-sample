package errors

type NotFoundError interface {
	error
	NotFoundError()
}

type notFoundError struct {
	BaseAppError
}

func (e *notFoundError) NotFoundError() {}

func NewNotFoundError(e error) NotFoundError {
	return &notFoundError{BaseAppError: BaseAppError{e}}
}
