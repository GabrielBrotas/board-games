package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GabrielBrotas/who-is-the-imposter/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func handleCreateUserOrLogin(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		Name string `json:"name"`
	}

	var body reqBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.NewUser(body.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userExists := usersRepository.GetUserByName(user.Name)

	if userExists != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s"}`, userExists.ID, userExists.Name)))
		return
	}

	err = usersRepository.AddUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s"}`, user.ID, user.Name)))
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s"}`, user.ID, user.Name)))
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := usersRepository.GetUsers()

	usersList := make([]models.User, 0, len(users))

	for _, user := range users {
		// print the userid
		log.Printf("User ID: %v", user.ID)
		log.Printf("User ID: %v", &user.ID)
		usersList = append(usersList, *user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"users": %v}`, usersList)))
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type reqBody struct {
		Name string `json:"name"`
	}

	var body reqBody

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.UpdateName(body.Name)

	err = usersRepository.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s"}`, user.ID, user.Name)))
}
