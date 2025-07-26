package domain

import (
	"strings"
	"github.com/tesso57/error-handling-sample/internal/errors"
)


// UserID is the unique identifier for a User.
type UserID string

// Email represents a user's email address.
type Email string

// User entity.
type User struct {
	ID    UserID
	Name  string
	Email Email
}

// NewUser creates a new User, validating fields.
func NewUser(id UserID, name string, email Email) (*User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	if !strings.Contains(string(email), "@") {
		return nil, errors.New("invalid email format")
	}
	return &User{ID: id, Name: name, Email: email}, nil
}

// UserRepository defines the interface for user persistence.
type UserRepository interface {
	Save(user *User) error
	FindByID(id UserID) (*User, error)
	FindByEmail(email Email) (*User, error)
}
