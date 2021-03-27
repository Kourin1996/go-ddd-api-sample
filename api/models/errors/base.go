package errors

type BaseAppError struct {
	err error
}

func (e *BaseAppError) Error() string {
	return e.err.Error()
}

func (e *BaseAppError) Cause() error {
	return e.err
}
