package repository

import "user_actions_sample_go/models"

// UserRepository defines the interface for interacting with users in the repository.
type UserRepository interface {
	GetByID(id int) (*models.User, error)
	GetAll() ([]models.User, error)
}
