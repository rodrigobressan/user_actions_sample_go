package utils

import (
	"sort"
	"user_actions_sample_go/models"
)

// ActionProbability represents a possible action and its likelihood
type ActionProbability struct {
	Type        string  `json:"type"`        // Action type, e.g., "ADD_TO_CRM"
	Probability float64 `json:"probability"` // Probability of this action occurring next
}

// ComputeNextActionProbs calculates the probability of next actions given a specific action type
func ComputeNextActionProbs(actions []models.Action, actionType string) []ActionProbability {
	// nextActionCount keeps track of how many times a specific next action appears
	nextActionCount := make(map[string]int)
	totalActionsCount := 0

	// Iterate through all actions, except the last (since we look ahead by one)
	for i := 0; i < len(actions)-1; i++ {
		current := actions[i]
		next := actions[i+1]

		// Ensure the current action matches the desired type and is from the same user
		if current.Type == actionType && next.UserID == current.UserID {
			nextActionCount[next.Type]++
			totalActionsCount++
		}
	}

	// If no transitions were found, return an empty slice
	if totalActionsCount == 0 {
		return []ActionProbability{}
	}

	// Convert the map to a slice of ActionProbability
	var result []ActionProbability
	for k, v := range nextActionCount {
		result = append(result, ActionProbability{
			Type:        k,
			Probability: float64(v) / float64(totalActionsCount),
		})
	}

	// Sort by probability in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].Probability > result[j].Probability
	})

	return result
}

func ComputeNextActionProbsNaive(actions []models.Action, actionType string) []ActionProbability {
	userActions := make(map[int][]models.Action) // key is userID, value is a slice of actions

	// Step 1: Group actions by user
	for _, a := range actions {
		userActions[a.UserID] = append(userActions[a.UserID], a)
	}

	nextActionCount := make(map[string]int)
	totalActionsCount := 0

	// Step 2: For each user's actions, count the occurrences of next actions
	for _, currentUserActions := range userActions { // iterate through actions of each user
		for i := 0; i < len(currentUserActions)-1; i++ { // iterate through actions of the user
			if currentUserActions[i].Type == actionType { // check if the current action matches the type
				nextAction := currentUserActions[i+1] // look at the nextAction action
				nextActionCount[nextAction.Type]++    // increment the count for this nextAction action type
				totalActionsCount++
			}
		}
	}

	// Step 3: Build the result
	result := make([]ActionProbability, 0, len(nextActionCount))
	if totalActionsCount == 0 {
		return []ActionProbability{}
	}

	for k, v := range nextActionCount {
		result = append(result, ActionProbability{
			Type:        k,
			Probability: float64(v) / float64(totalActionsCount),
		})
	}

	// Sort by probability in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].Probability > result[j].Probability
	})

	return result
}
