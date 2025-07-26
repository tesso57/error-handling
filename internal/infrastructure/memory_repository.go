package infrastructure

import (
	"sync"
	"github.com/tesso57/error-handling-sample/internal/domain"
)

// InMemoryUserRepository is an in-memory implementation of UserRepository.
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[domain.UserID]*domain.User
	byEmail map[domain.Email]*domain.User
}

// NewInMemoryUserRepository creates a new in-memory repository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[domain.UserID]*domain.User),
		byEmail: make(map[domain.Email]*domain.User),
	}
}

// Save stores a user in memory.
func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	r.byEmail[user.Email] = user
	return nil
}

// FindByID retrieves a user by ID.
func (r *InMemoryUserRepository) FindByID(id domain.UserID) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.users[id], nil
}

// FindByEmail retrieves a user by email.
func (r *InMemoryUserRepository) FindByEmail(email domain.Email) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.byEmail[email], nil
}