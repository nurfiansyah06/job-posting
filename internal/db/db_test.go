package db_test

import (
	"job-posting/internal/db"
	"testing"
)

func TestConnectDB(t *testing.T) {
	t.Run("ConnectDB", func(t *testing.T) {
		db, err := db.ConnectDB()
		if err != nil {
			t.Errorf("Error connecting to the database: %v", err)
		}
		defer db.Close()
	})
}
