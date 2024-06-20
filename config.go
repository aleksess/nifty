package nifty

import (
	"encoding/json"
	"io"
)

type Config struct {
	Port           uint   `json:"port"`
	DatabaseUrl    string `json:"databaseUrl"`
	DatabaseEngine string `json:"databaseEngine"`
	AuthSecret  string `json:"sessionSecret"`
}

var DefaultConfig = Config{Port: 3000, DatabaseUrl: ":memory:", DatabaseEngine: "sqlite", AuthSecret: "secret"}

func LoadConfig(r io.Reader) Config {
	var config Config
	json.NewDecoder(r).Decode(&config)

	return config
}
