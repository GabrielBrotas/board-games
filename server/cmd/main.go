package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("OPENAI_KEY") == "" {
		log.Fatalf("OPENAI_KEY environment variable not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/ws", handleConnections)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
