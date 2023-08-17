package data

import (
	"database/sql"

	"github.com/rs/zerolog"
)

// returns a map of all the endpoints stored in the database
func GetEndpoints(l *zerolog.Logger, db *sql.DB) (map[string]string, error) {
	// m[id][endpoint, url]
	m := make(map[string]string)
	rows, err := db.Query("SELECT endpoint, url FROM endpoints")
	if err != nil {
		l.Fatal().Err(err).Msg("Error querying data")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var endpoint, url string
		err := rows.Scan(&endpoint, &url)
		if err != nil {
			l.Fatal().Err(err).Msg("Error reading queried data")
			return nil, err
		}
		m[endpoint] = url
	}

	return m, nil
}

// returns false if the endpoint is not found
func CheckEndpoint(l *zerolog.Logger, db *sql.DB, endpoint string) (bool, string) {
	var url string
	err := db.QueryRow("SELECT url FROM endpoints WHERE endpoint = ?", endpoint).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			// endpoint not found
			return false, ""
		}
		l.Error().Err(err).Msg("Error checking endpoint existence")
		return false, ""
	}

	// endpoint found
	return true, url
}
