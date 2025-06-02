package memory

import (
	"surfe_assignment/models"
)

// InMemoryActionRepo is an in-memory implementation of the ActionsRepository interface.
type InMemoryActionRepo struct {
	actions []models.Action
}

// NewActionsRepo creates a new in-memory actions repository with the provided actions.
func NewActionsRepo(actions []models.Action) *InMemoryActionRepo {
	return &InMemoryActionRepo{actions: actions}
}

// GetActionsForUser retrieves all actions for a specific user from the in-memory repository.
func (r *InMemoryActionRepo) GetActionsForUser(userID int) ([]models.Action, error) {
	var userActions []models.Action
	for _, a := range r.actions {
		if a.UserID == userID {
			userActions = append(userActions, a)
		}
	}

	return userActions, nil
}

// GetAll retrieves all actions from the in-memory repository.
func (r *InMemoryActionRepo) GetAll() ([]models.Action, error) {
	return r.actions, nil
}
