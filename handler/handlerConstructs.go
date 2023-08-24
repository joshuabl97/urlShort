package handler

import (
	"database/sql"

	"github.com/joshuabl97/urlShort/data"
	"github.com/rs/zerolog"
)

// fields passed to handlers
type HandlerHelper struct {
	l       *zerolog.Logger
	db      *sql.DB
	Message string
}

type Shortcuts struct {
	Shortcuts []data.Endpoint `json:"shortcuts"`
}
