package handler

import (
	"database/sql"

	"github.com/joshuabl97/urlShort/data"
	"github.com/rs/zerolog"
)

// a JsonRequest containing shortcut info
type JsonRequest struct {
	Endpoint string `json:"endpoint"`
	URL      string `json:"url"`
}

// fields passed to handlers
type HandlerHelper struct {
	l  *zerolog.Logger
	db *sql.DB
}

type Shortcuts struct {
	Shortcuts []data.Endpoint `json:"shortcuts"`
}
