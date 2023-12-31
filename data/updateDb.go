package data

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
)

// adds a single shortcut to the db
func AddEndpoint(db *sql.DB, endpoint string, url string) (*sql.DB, error) {
	// Check if the endpoint and URL already exist in the database
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM endpoints WHERE endpoint = ? OR url = ?", endpoint, url).Scan(&count)
	if err != nil {
		return db, fmt.Errorf("error querying db: %v", err)
	}

	if count > 0 {
		return db, fmt.Errorf("Endpoint or URL already exists - Endpoint: %v URL: %v", endpoint, url)
	}

	// Insert data into the endpoints table
	_, err = db.Exec("INSERT INTO endpoints (endpoint, url) VALUES (?, ?)", endpoint, url)
	if err != nil {
		return db, fmt.Errorf("error inserting data: %v", err)
	}

	return db, nil
}

// changes the URL to the newURL
func UpdateEndpoint(l *zerolog.Logger, db *sql.DB, endpoint string, newURL string) error {
	_, err := db.Exec("UPDATE endpoints SET url = ? WHERE endpoint = ?", newURL, endpoint)
	if err != nil {
		l.Error().Err(err).Msg("Error updating endpoint")
		return err
	}

	return nil
}

// deletes a row from the database
func DeleteRow(l *zerolog.Logger, db *sql.DB, endpoint string) error {
	_, err := db.Exec("DELETE FROM endpoints WHERE endpoint = ?", endpoint)
	if err != nil {
		l.Error().Err(err).Msg("Error deleting row")
		return err
	}

	return nil
}

// adds multiple endpoints to the DB
func AddMultipleEndpoints(data []Endpoint, l *zerolog.Logger, db *sql.DB) *sql.DB {
	for _, shortcut := range data {
		var err error
		db, err = AddEndpoint(db, shortcut.Endpoint, shortcut.URL)
		if err != nil {
			l.Error().
				Err(err).
				Str("endpoint", shortcut.Endpoint).
				Str("url", shortcut.URL).
				Msg("Error adding shortcut to db")
		}
	}

	return db
}
