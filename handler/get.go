package handler

import (
	"encoding/json"
	"net/http"
)

// return all the shortcuts to the user in JSON format
// i.e {"shortcuts": [
// {"endpoint":"example1","url":"https://google.com"},
// {"endpoint":"example2","url":"https://example.com/"}
// ]
func (h *HandlerHelper) GetShortcuts(w http.ResponseWriter, r *http.Request) {
	rows, w := getAndMarshalEndpoints(h.l, h.db, w)

	result := Shortcuts{Shortcuts: rows}

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(result)
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
