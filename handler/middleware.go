package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joshuabl97/urlShort/data"
	"github.com/rs/zerolog"
)

// checks if endpoint and url are present in a request body
// returns true if valid and returns the data as a struct (Type JsonRequest)
func validateRequestJson(l *zerolog.Logger, w http.ResponseWriter, r *http.Request) (bool, data.Endpoint) {
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

func getAndMarshalEndpoints(l *zerolog.Logger, db *sql.DB, w http.ResponseWriter) ([]data.Endpoint, http.ResponseWriter) {
	m, err := data.GetEndpoints(l, db)
	if err != nil {
		http.Error(w, "Cannot query database for endpoints table", http.StatusInternalServerError)
		return nil, w
	}

	var rows []data.Endpoint
	for k, v := range m {
		rows = append(rows, data.Endpoint{Endpoint: k, URL: v})
	}

	return rows, w
}
