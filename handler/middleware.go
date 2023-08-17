package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

// checks if endpoint and url are present in a request body
// returns true if valid and returns the data as a struct (Type JsonRequest)
func ValidateRequestJson(l *zerolog.Logger, w http.ResponseWriter, r *http.Request) (bool, JsonRequest) {
	var request JsonRequest
	// parse the json body
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("%v", request)
	if err != nil {
		l.Error().Err(err).Str("PUT /shortcut", "Invalid JSON format")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return false, request
	}

	// validate the JSON structure
	if request.Endpoint == "" || request.URL == "" {
		l.Error().Msg("Missing required fields")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return false, request
	}

	return true, request
}
