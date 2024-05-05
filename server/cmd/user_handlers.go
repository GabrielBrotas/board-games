package main

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrotas/board-games/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// handleCreateUserOrLogin processes user login or registration.
func handleCreateUserOrLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := models.NewUser(body.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if existingUser := usersRepository.GetUserByName(user.Name); existingUser != nil {
		respondWithJSON(w, http.StatusCreated, existingUser.ToOut())
		return
	}

	if err := usersRepository.AddUser(user); err != nil {
		http.Error(w, "Failed to add user", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusCreated, user.ToOut())
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, user.ToOut())
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := usersRepository.GetUsers()
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"users": users})
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.UpdateName(body.Name)

	if err := usersRepository.UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusOK, user.ToOut())
}