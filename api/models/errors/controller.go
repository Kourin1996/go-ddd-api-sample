package errors

type InvalidRequestError interface {
	error
	InvalidRequestError()
}

type invalidRequestError struct {
	BaseAppError
}

func (e *invalidRequestError) InvalidRequestError() {}

func NewInvalidRequestError(e error) InvalidRequestError {
	return &invalidRequestError{BaseAppError: BaseAppError{e}}
}
