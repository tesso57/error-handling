package application

import (
	"github.com/tesso57/error-handling-sample/internal/domain"
	"github.com/tesso57/error-handling-sample/internal/errors"
)

// UserService provides user use cases.
type UserService struct {
	repo domain.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{repo: r}
}

// RegisterUser registers a new user.
func (s *UserService) RegisterUser(id domain.UserID, name string, email domain.Email) (*domain.User, error) {
	user, err := domain.NewUser(id, name, email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	if existing, err := s.repo.FindByEmail(email); err != nil {
		return nil, errors.Wrap(err, "failed to check existing email")
	} else if existing != nil {
		return nil, ErrEmailAlreadyRegistered
	}
	if err := s.repo.Save(user); err != nil {
		return nil, errors.Wrap(err, "failed to save user")
	}
	return user, nil
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(id domain.UserID) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user by ID")
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
