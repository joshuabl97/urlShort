package handler

import (
	"database/sql"

	"github.com/rs/zerolog"
)

// creates a handler helper
func NewHandlerHelper(l *zerolog.Logger, db *sql.DB) *HandlerHelper {
	return &HandlerHelper{l, db}
}
