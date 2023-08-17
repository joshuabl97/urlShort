package handler

import (
	"encoding/json"
	"net/http"

	"github.com/joshuabl97/urlShort/data"
)

func (h *HandlerHelper) GetShortcuts(w http.ResponseWriter, r *http.Request) {
	m, err := data.GetEndpoints(h.l, h.db)
	if err != nil {
		http.Error(w, "Cannot query database for endpoints table", http.StatusInternalServerError)
		return
	}

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(m)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		h.l.Error().Err(err).Msg("Failed to marshal JSON")
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
