package repository

import "surfe_assignment/models"

// ActionsRepository defines the interface for interacting with user actions in the repository.
type ActionsRepository interface {
	GetActionsForUser(userID int) ([]models.Action, error)
	GetAll() ([]models.Action, error)
}
