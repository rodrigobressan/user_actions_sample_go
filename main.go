package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"user_actions_sample_go/handlers"
	"user_actions_sample_go/middleware"
	"user_actions_sample_go/models"
	"user_actions_sample_go/repository/memory"
)

func loadUsers() ([]models.User, error) {
	var users []models.User
	userFile, err := os.ReadFile("data/users.json")
	if err != nil {
		return nil, errors.New("error reading users.json: " + err.Error())
	}

	if err := json.Unmarshal(userFile, &users); err != nil {
		return nil, errors.New("error unmarshalling users.json: " + err.Error())
	}

	return users, nil
}

func loadActions() ([]models.Action, error) {
	var actions []models.Action
	actionFile, err := os.ReadFile("data/actions.json")
	if err != nil {
		return nil, errors.New("error reading actions.json: " + err.Error())
	}

	if err := json.Unmarshal(actionFile, &actions); err != nil {
		return nil, errors.New("error unmarshalling actions.json: " + err.Error())
	}

	return actions, nil
}

func main() {
	users, err := loadUsers()
	if err != nil {
		log.Fatalf("Failed to load users: %v", err)
	}

	actions, err := loadActions()
	if err != nil {
		log.Fatalf("Failed to load actions: %v", err)
	}

	userRepo := memory.NewUserRepo(users)
	actionsRepo := memory.NewActionsRepo(actions)

	handler := &handlers.Handler{
		UserRepository:    userRepo,
		ActionsRepository: actionsRepo,
	}

	r := mux.NewRouter()
	r.Use(middleware.JSONMiddleware)
	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/", handler.Index).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}/actions/count", handler.GetUserActionCount).Methods("GET")
	r.HandleFunc("/actions/next/{type}", handler.GetNextActionProbabilities).Methods("GET")
	r.HandleFunc("/referral_index", handler.GetReferralIndex).Methods("GET")

	log.Println("Server running on :8080")

	// Alternatively, we could also load this port (and any other configuration) from an environment variable or a config file
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
