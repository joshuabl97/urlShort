package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joshuabl97/urlShort/data"
	"github.com/rs/zerolog"
)

// checks if endpoint and url are present in a request body
// returns true if valid and returns the data as a struct (Type JsonRequest)
func ValidateRequestJson(l *zerolog.Logger, w http.ResponseWriter, r *http.Request) (bool, data.Endpoint) {
	request, err := getShortcutRequest(l, w, r)
	if err != nil {
		l.Error().Err(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
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

func getShortcutRequest(l *zerolog.Logger, w http.ResponseWriter, r *http.Request) (data.Endpoint, error) {
	var request data.Endpoint
	// parse the json body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, fmt.Errorf("Invalid JSON format: %v", err)
	}

	return request, nil
}
