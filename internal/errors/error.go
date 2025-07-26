package errors

import (
	"fmt"

	"connectrpc.com/connect"

	"github.com/cockroachdb/errors"
)

type ErrorResponse struct {
	Code    connect.Code
	Message string
	cause   error
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) Unwrap() error {
	return e.cause
}

func (e *ErrorResponse) Format(s fmt.State, verb rune) {
	errors.FormatError(e, s, verb)
}

func (e *ErrorResponse) ToConnectError() *connect.Error {
	return connect.NewError(e.Code, errors.New(e.Message))
}

func NewErrorResponse(code connect.Code, message string, cause error) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		cause:   cause,
	}
}

func ToConnectError(err error) *connect.Error {
	if err == nil {
		return nil
	}

	var customErr *ErrorResponse
	if errors.As(err, &customErr) {
		return customErr.ToConnectError()
	}

	return connect.NewError(connect.CodeInternal, errors.New("unknown error"))
}
