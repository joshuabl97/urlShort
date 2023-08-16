package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
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
}
