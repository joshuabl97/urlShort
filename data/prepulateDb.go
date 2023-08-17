package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
)

type Endpoint struct {
	Endpoint string `yaml:"endpoint"`
	URL      string `yaml:"url"`
}

// adds the yamlfile to the DB
func AddYamlToDB(yamlPath string, l *zerolog.Logger, db *sql.DB) *sql.DB {
	// open the YAML file
	file, err := os.ReadFile(yamlPath)
	if err != nil {
		l.Error().Err(err).Msg("Error reading YAML file")
	}

	// create a map to store the data
	data := []Endpoint{}

	// unmarshal the YAML into the map
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		l.Error().Err(err).Msg("Error unmarshaling YAML")
	}

	for _, shortcut := range data {
		fmt.Println(shortcut.Endpoint, shortcut.URL)
		db, err = AddEndpoint(l, db, shortcut.Endpoint, shortcut.URL)
		if err != nil {
			l.Error().Err(err).Msg("Error adding shortcut to db")
		}
	}

	return db
}
