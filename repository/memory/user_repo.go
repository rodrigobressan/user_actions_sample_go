package memory

import (
	"errors"
	"user_actions_sample_go/models"
)

// InMemoryUserRepo is an in-memory implementation of the UserRepository interface.
type InMemoryUserRepo struct {
	users []models.User
}

// NewUserRepo creates a new in-memory user repository with the provided users.
func NewUserRepo(users []models.User) *InMemoryUserRepo {
	return &InMemoryUserRepo{users}
}

// GetByID retrieves a user by their ID from the in-memory repository.
func (r *InMemoryUserRepo) GetByID(id int) (*models.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAll retrieves all users from the in-memory repository.
func (r *InMemoryUserRepo) GetAll() ([]models.User, error) {
	return r.users, nil
}
