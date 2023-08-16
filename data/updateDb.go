package data

import (
	"database/sql"

	"github.com/rs/zerolog"
)

func AddToDB(l *zerolog.Logger, db *sql.DB, endpoint string, url string) (*sql.DB, error) {
	_, err := db.Exec("INSERT INTO endpoints (endpoint, url) VALUES (?, ?)", endpoint, url)
	if err != nil {
		l.Error().Err(err).Msg("Error inserting data")
		return db, err
	}

	return db, nil
}
