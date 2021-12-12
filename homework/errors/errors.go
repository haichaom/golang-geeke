package errors

import (
	"fmt"
)

type Error struct {
	Code    int32
	Reason  string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d, reason = %s, message = %s", e.Code, e.Reason, e.Message)
}

func New(code int, reason string, message string) *Error {
	return &Error{
		Code:    int32(code),
		Reason:  reason,
		Message: message,
	}
}

func BadRequest(reason string, message string) *Error {
	return New(400, reason, message)
}
func InternalServer(reason string, message string) *Error {
	return New(500, reason, message)
}
