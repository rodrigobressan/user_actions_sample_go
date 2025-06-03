package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"surfe_assignment/handlers"
	"surfe_assignment/models"
	"surfe_assignment/repository/memory"
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

	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}/actions/count", handler.GetUserActionCount).Methods("GET")
	r.HandleFunc("/actions/next/{type}", handler.GetNextActionProbabilities).Methods("GET")

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
