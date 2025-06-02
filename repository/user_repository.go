package repository

import "surfe_assignment/models"

// UserRepository defines the interface for interacting with users in the repository.
type UserRepository interface {
	GetByID(id int) (*models.User, error)
	GetAll() ([]models.User, error)
}
