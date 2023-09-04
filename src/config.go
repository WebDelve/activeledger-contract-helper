package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	PrivatePem string       `json:"privatePem"`
	Identity   string       `json:"identity"`
	Namespace  string       `json:"namespace"`
	Ledger     LedgerConfig `json:"ledger"`
}

type LedgerConfig struct {
	Url  string `json:"url"`
	Port int    `json:"port"`
}

func loadConfig() *Config {
	var c Config

	data := readFile("./config.json")

	if err := json.Unmarshal(data, &c); err != nil {
		handleError(err, "Error unmarshalling contect of config file")
	}

	return &c
}

func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		handleError(err, fmt.Sprintf("Error reading file with path %s", path))
	}

	return data
}
