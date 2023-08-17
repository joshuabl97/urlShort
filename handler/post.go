package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joshuabl97/urlShort/data"
)

type JsonRequest struct {
	Endpoint string `json:"endpoint"`
	URL      string `json:"url"`
}

// Put /shortcut creates a new shortcut in the endpoints table
func (h *HandlerHelper) CreateShortcut(w http.ResponseWriter, r *http.Request) {
	// parse the JSON body
	var request JsonRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("%v", request)
	if err != nil {
		h.l.Error().Err(err).Str("PUT /shortcut", "Invalid JSON format")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// validate the JSON structure
	if request.Endpoint == "" || request.URL == "" {
		h.l.Error().Err(err).Str("PUT /shortcut", "Missing required fields")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	h.l.Info().
		Str("Endpoint: ", request.Endpoint).
		Str("URL: ", request.URL).
		Msg("Processing AddEndpoint")

	h.db, err = data.AddEndpoint(h.l, h.db, request.Endpoint, request.URL)
	if err != nil {
		h.l.Error().Err(err).Msg("Failed to add request body to DB")
		http.Error(w, "Failed to add request body to DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request body received and validated"))
}
