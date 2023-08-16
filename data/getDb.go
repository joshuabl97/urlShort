package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
)

// returns a *sql.DB and an error
// you still need to close the db!!!
// i.e db, _  := getDb()
// defer db.Close()
func MakeDb(l *zerolog.Logger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "mydatabase.db")
	if err != nil {
		l.Fatal().Err(err).Msg("Error opening database")
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS endpoints (
        endpoint TEXT,
        url TEXT
    )
	`)

	if err != nil {
		l.Fatal().Err(err).Msg("Failed to create endpoints sqlite table")
		return nil, err
	}

	return db, nil
}
