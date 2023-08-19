package data

import (
	"database/sql"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// a single shortcut
type Endpoint struct {
	Endpoint string `yaml:"endpoint"`
	URL      string `yaml:"url"`
}

func ParseYaml(yamlPath string, db *sql.DB) ([]Endpoint, error) {
	// open the YAML fil
	file, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file: %v", err)
	}

	// create a map to store the data
	data := []Endpoint{}

	// unmarshal the YAML into the []struct
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling yaml: %v", err)
	}

	return data, nil
}
