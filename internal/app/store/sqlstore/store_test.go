package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=root password=root host=localhost dbname=db_test  sslmode=disable"
	}

	os.Exit(m.Run())
}
