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

type NotAllowedError interface {
	error
	NotAllowedError()
}

type notAllowedError struct {
	BaseAppError
}

func (e *notAllowedError) NotAllowedError() {}

func NewNotAllowedError(e error) NotAllowedError {
	return &notAllowedError{BaseAppError: BaseAppError{e}}
}
