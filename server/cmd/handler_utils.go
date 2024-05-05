package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// respondWithJSON sends a JSON response.
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// parseUUID parses a UUID from the provided map using the specified key.
func parseUUID(data map[string]interface{}, key string) (uuid.UUID, error) {
	idStr, ok := data[key].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid %s value", key)
	}
	return uuid.Parse(idStr)
}
