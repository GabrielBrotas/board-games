package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GabrielBrotas/who-is-the-imposter/internal/repository"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	usersRepository = repository.NewUserRepository()
)

func main() {
	if os.Getenv("OPENAI_KEY") == "" {
		log.Fatalf("OPENAI_KEY environment variable not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := mux.NewRouter()

	r.HandleFunc("/login", handleCreateUserOrLogin).Methods(http.MethodPost)
	r.HandleFunc("/users", handleGetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handleGetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handleUpdateUser).Methods(http.MethodPut)

	r.HandleFunc("/ws", handleConnections)
	r.HandleFunc("/player-list", handleGetPlayerList)
	r.HandleFunc("/game-status", handleGetGameStatus)

	// Configure CORS to allow all origins
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsOrigins, corsMethods, corsHeaders)(r)))
}
