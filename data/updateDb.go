package data

import (
	"database/sql"

	"github.com/rs/zerolog"
)

func AddEndpoint(l *zerolog.Logger, db *sql.DB, endpoint string, url string) (*sql.DB, error) {
	// Check if the endpoint and URL already exist in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM endpoints WHERE endpoint = ? OR url = ?", endpoint, url).Scan(&count)
	if err != nil {
		l.Error().Err(err).Msg("Error checking existing data")
		return db, err
	}

	if count > 0 {
		l.Info().
			Str("Endpoint", endpoint).
			Str("URL", url).
			Msg("Endpoint or URL already exists in the database")
		return db, nil
	}

	// Insert data into the endpoints table
	_, err = db.Exec("INSERT INTO endpoints (endpoint, url) VALUES (?, ?)", endpoint, url)
	if err != nil {
		l.Error().Err(err).Msg("Error inserting data")
		return db, err
	}

	return db, nil
}
