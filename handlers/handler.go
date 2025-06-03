package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"surfe_assignment/repository"
	"surfe_assignment/utils"

	"github.com/gorilla/mux"
)

// Handler is the main struct that holds the repositories needed for handling requests.
type Handler struct {
	UserRepository    repository.UserRepository
	ActionsRepository repository.ActionsRepository
}

// GetUser handles the request to get a user by ID.
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepository.GetByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetUserActionCount handles the request to get the count of all actions for a specific user.
func (h *Handler) GetUserActionCount(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	actions, err := h.ActionsRepository.GetActionsForUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Possible improvement: Use a struct for the response instead of a map
	response := make(map[string]int)
	response["count"] = len(actions)
	json.NewEncoder(w).Encode(response)
}

// GetNextActionProbabilities handles the request to get the next action probabilities based on the action type.
func (h *Handler) GetNextActionProbabilities(w http.ResponseWriter, r *http.Request) {
	actionType := mux.Vars(r)["type"]

	allActions, err := h.ActionsRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	probs := utils.ComputeNextActionProbs(allActions, actionType)
	json.NewEncoder(w).Encode(probs)
}

// GetReferralIndex handles the request to compute the referral index for all users based on their actions.
func (h *Handler) GetReferralIndex(w http.ResponseWriter, r *http.Request) {
	allActions, err := h.ActionsRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index := utils.ComputeReferralIndex(allActions)
	json.NewEncoder(w).Encode(index)
}
