package nifty_test

import (
	"strings"
	"testing"

	"github.com/aleksess/nifty"
)

func TestConfig(t *testing.T) {
	configString := "{\"port\": 3000, \"databaseUrl\": \":memory:\"}"
	configReader := strings.NewReader(configString)

	config := nifty.LoadConfig(configReader)

	if config.Port != 3000 || config.DatabaseUrl != ":memory:" {
		t.Errorf("Wrong config parse")
	}
}
