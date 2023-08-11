package exception

type ErrorNotFound struct {
	Message string
}

func NewErrorNotFound(message string) *ErrorNotFound {
	return &ErrorNotFound{
		Message: message,
	}
}

func (errornotfound *ErrorNotFound) Error() string {
	return errornotfound.Message
}
