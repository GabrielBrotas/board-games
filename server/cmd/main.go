package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("OPENAI_KEY") == "" {
		log.Fatalf("OPENAI_KEY environment variable not set")
	}

	http.HandleFunc("/ws", handleConnections)
	fmt.Println("Server running on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
