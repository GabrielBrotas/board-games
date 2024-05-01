package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("Server running on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
