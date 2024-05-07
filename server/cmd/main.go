package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/GabrielBrotas/board-games/internal/repository"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var usersRepository = repository.NewUserRepository()

func main() {
	ensureEnvironmentVariables()

	r := setupRouter()

	log.Printf("Server running on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), setupCORS(r)))
}

func ensureEnvironmentVariables() {
	requiredEnvVars := map[string]string{
		"OPENAI_KEY": "OPENAI_KEY environment variable not set",
		"PORT":       "PORT environment variable must be set",
	}

	for envVar, errMsg := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf(errMsg)
		}
	}
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	// API version endpoint
	r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"version": "0.1.0"}`))
	}).Methods(http.MethodGet)

	// User management endpoints
	r.HandleFunc("/login", handleCreateUserOrLogin).Methods(http.MethodPost)
	r.HandleFunc("/users", handleGetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handleGetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handleUpdateUser).Methods(http.MethodPut)

	// Game-specific routes for the "Impostor" game
	r.HandleFunc("/games/impostor/ws", handleImpostorConnections)
	r.HandleFunc("/games/impostor/player-list", handleGetImpostorPlayerList)
	r.HandleFunc("/games/impostor/status", handleImpostorGetGameStatus)

	// Spyfall
	r.HandleFunc("/games/spyfall/ws", handleSpyfallConnections)
	r.HandleFunc("/games/spyfall/player-list", handleGetSpyfallPlayerList)
	r.HandleFunc("/games/spyfall/status", handleSpyfallGetPlayerGameStatus)

	return r
}

// setupCORS configures Cross-Origin Resource Sharing (CORS) settings.
func setupCORS(r *mux.Router) http.Handler {
	corsOrigins := strings.Split(os.Getenv("CORS_ORIGINS"), ",")
	log.Printf("CORS origins: %v", corsOrigins)

	return handlers.CORS(
		handlers.AllowedOrigins(corsOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(r)
}
