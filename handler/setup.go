package handler

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type JsonRequest struct {
	Endpoint string `json:"endpoint"`
	URL      string `json:"url"`
}

type HandlerHelper struct {
	l  *zerolog.Logger
	db *sql.DB
}

// creates a handler helper
func NewHandlerHelper(l *zerolog.Logger, db *sql.DB) *HandlerHelper {
	return &HandlerHelper{l, db}
}
