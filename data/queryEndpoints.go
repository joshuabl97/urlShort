package data

import (
	"database/sql"

	"github.com/rs/zerolog"
)

// returns a map of all the endpoints stored in the database
func GetEndpoints(l *zerolog.Logger, db *sql.DB) map[string]string {
	// m[id][endpoint, url]
	m := make(map[string]string)
	rows, err := db.Query("SELECT endpoint, url FROM endpoints")
	if err != nil {
		l.Fatal().Err(err).Msg("Error querying data")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var endpoint, url string
		err := rows.Scan(&endpoint, &url)
		if err != nil {
			l.Fatal().Err(err).Msg("Error reading queried data")
			return nil
		}
		m[endpoint] = url
		l.Info().Msg(endpoint + " " + url)
	}

	return m
}
