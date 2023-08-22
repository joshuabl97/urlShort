package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/joshuabl97/urlShort/data"
)

// post /shortcut
func (h *HandlerHelper) CreateShortcut(w http.ResponseWriter, r *http.Request) {
	request, err := getShortcutRequest(h.l, w, r)
	if err != nil {
		h.l.Error().Err(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	if request.Endpoint == "" {
		// only a URL is required in the request body
		if request.URL == "" {
			h.l.Error().Msg("Missing URL from JSON")
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		// check to see if the endpoint already exists
		exists, _ := data.CheckEndpoint(h.l, h.db, request.Endpoint)
		if exists {
			h.l.Error().
				Str("Endpoint", request.Endpoint).
				Str("URL", request.URL).
				Msg("Endpoint or URL already exists")
			badEndpoint := fmt.Sprintf("Endpoint or URL already exists - Endpoint: %v URL: %v", request.Endpoint, request.URL)
			http.Error(w, badEndpoint, http.StatusBadRequest)
			return
		}

		// implement a skipto feature where it checks if a random string is already in the db maybe?
		request.Endpoint = generateRandomString(5)
	}

	h.l.Info().
		Str("Endpoint: ", request.Endpoint).
		Str("URL: ", request.URL).
		Msg("Processing UniqueShortcut")

	// adds the new shortcut to the database
	h.db, err = data.AddEndpoint(h.db, request.Endpoint, request.URL)
	if err != nil {
		h.l.Error().Err(err).Msg("Failed to add request body to DB")
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
		h.l.Error().Err(err).Msg("Failed to marshal JSON response")
		return
	}

	// Set the Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON data to the ResponseWriter
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		return
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
