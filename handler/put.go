package handler

import (
	"fmt"
	"net/http"

	"github.com/joshuabl97/urlShort/data"
)

func (h *HandlerHelper) UpdateEndpoint(w http.ResponseWriter, r *http.Request) {
	// validate and parse the json in the request
	valid, request := validateRequestJson(h.l, w, r)
	if !valid {
		h.l.Error().Msg("Invalid request CreateShortcut")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// check to see if the endpoint already exists
	exists, _ := data.CheckEndpoint(h.l, h.db, request.Endpoint)
	if !exists {
		h.l.Error().Str("Endpoint", request.Endpoint).Msg("Endpoint doesn't exist")
		badEndpoint := fmt.Sprintf("Endpoint %v doesn't exist", request.Endpoint)
		http.Error(w, badEndpoint, http.StatusBadRequest)
		return
	}

	h.l.Info().
		Str("Endpoint: ", request.Endpoint).
		Str("URL: ", request.URL).
		Msg("Processing UpdateEndpoint")

	// update url where endpoint is found in db
	err := data.UpdateEndpoint(h.l, h.db, request.Endpoint, request.URL)
	if err != nil {
		h.l.Error().Err(err).Msg("Internal Server Error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Sucessfully updated database"))
}
