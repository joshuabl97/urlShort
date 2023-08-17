package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joshuabl97/urlShort/data"
)

// deletes an endpoint from the DB
func (h *HandlerHelper) DeleteEndpoint(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")

	// check to see if the endpoint already exists
	exists, _ := data.CheckEndpoint(h.l, h.db, endpoint)
	if !exists {
		h.l.Error().Str("Endpoint", endpoint).Msg("Endpoint doesn't exist")
		http.Error(w, "Endpoint doesn't exist", http.StatusBadRequest)
		return
	}

	err := data.DeleteRow(h.l, h.db, endpoint)
	if err != nil {
		http.Error(w, "Unable to delete endpoint", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Succssfully removed endpoint from database"))
}
