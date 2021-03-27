package errors

type InvalidDataError interface {
	error
	InvalidDataError()
}

type invalidDataError struct {
	BaseAppError
}

func (e *invalidDataError) InvalidDataError() {}

func NewInvalidDataError(e error) InvalidDataError {
	return &invalidDataError{BaseAppError: BaseAppError{e}}
}
