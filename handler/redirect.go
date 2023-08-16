package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joshuabl97/urlShort/data"
	"github.com/rs/zerolog"
)

type HandlerHelper struct {
	l  *zerolog.Logger
	db *sql.DB
}

// NewProducts creates a products handler with the given logger
func NewHandlerHelper(l *zerolog.Logger, db *sql.DB) *HandlerHelper {
	return &HandlerHelper{l, db}
}

func (h *HandlerHelper) Redirect(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")
	h.l.Info().Msg("Endpoint: " + endpoint + " being processed")

	ok, url := data.CheckEndpoint(h.l, h.db, endpoint)

	if !ok {
		h.l.Error().Msg("Endpoint not found")
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
