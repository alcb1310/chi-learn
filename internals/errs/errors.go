package errs

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
