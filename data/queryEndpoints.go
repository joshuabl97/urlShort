package data

import (
	"database/sql"

	"github.com/rs/zerolog"
)

// returns a map of all the endpoints stored in the database
func GetEndpoints(l *zerolog.Logger, db *sql.DB) map[int][]string {
	// m[id][endpoint, url]
	m := make(map[int][]string)
	rows, err := db.Query("SELECT id, endpoint, url FROM endpoints")
	if err != nil {
		l.Fatal().Err(err).Msg("Error querying data")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var endpoint, url string
		err := rows.Scan(&id, &endpoint, &url)
		if err != nil {
			l.Fatal().Err(err).Msg("Error reading queried data")
			return nil
		}
		m[id] = []string{endpoint, url}
		l.Info().Msg(endpoint + " " + url)
	}

	return m
}
