package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(testing *testing.T, databaseURL string) (*Store, func(...string)) {
	testing.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	store := New(config)
	if err := store.Open(); err != nil {
		testing.Fatal()
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := store.database.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				testing.Fatal(err)
			}
		}

		store.Close()
	}
}
