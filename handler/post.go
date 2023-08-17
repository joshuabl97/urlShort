package handler

import (
	"fmt"
	"net/http"

	"github.com/joshuabl97/urlShort/data"
)

// Post /shortcut creates a new shortcut in the endpoints table
func (h *HandlerHelper) CreateShortcut(w http.ResponseWriter, r *http.Request) {
	// validate and parse the json in the request
	valid, request := ValidateRequestJson(h.l, w, r)
	if !valid {
		h.l.Error().Msg("Invalid request CreateShortcut")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// check to see if the endpoint already exists
	exists, _ := data.CheckEndpoint(h.l, h.db, request.Endpoint)
	if exists {
		h.l.Error().Str("Endpoint", request.Endpoint).Msg("Endpoint already exists")
		badEndpoint := fmt.Sprintf("Endpoint %v already exists", request.Endpoint)
		http.Error(w, badEndpoint, http.StatusBadRequest)
		return
	}

	h.l.Info().
		Str("Endpoint: ", request.Endpoint).
		Str("URL: ", request.URL).
		Msg("Processing AddEndpoint")

	var err error
	h.db, err = data.AddEndpoint(h.l, h.db, request.Endpoint, request.URL)
	if err != nil {
		h.l.Error().Err(err).Msg("Failed to add request body to DB")
		http.Error(w, "Failed to add request body to DB", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Sucessfully added to database"))
}
