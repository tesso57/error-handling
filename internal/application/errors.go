package application

import "github.com/tesso57/error-handling-sample/internal/errors"

// Application-specific sentinel errors
var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrUserNotFound          = errors.New("user not found")
)