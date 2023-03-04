package main

import (
	"encoding/json"
	"os"
	"text/template"
)

func main() {
	// create template
	tmp := template.New("settings.json")
	tmp, err := tmp.Parse(SETTINGS)
	if err != nil {
		panic(err)
	}

	// load config
	cfg, err := loadConfig("colors.json")
	if err != nil {
		panic(err)
	}

	// open & truncate target file
	file, err := os.OpenFile(cfg.Target, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// execute template
	err = tmp.Execute(file, cfg.Colors)
	if err != nil {
		panic(err)
	}
}

// LOAD CONFIG
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
