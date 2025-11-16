package store

import (
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

var SQLiteStore = &SQLiteStorage{}

func (s *SQLiteStorage) Init() error {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	// Create table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user_data (
			user_id TEXT PRIMARY KEY,
			data TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *SQLiteStorage) Write(userId string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
		INSERT OR REPLACE INTO user_data (user_id, data)
		VALUES (?, ?)`, userId, string(jsonData))
	return err
}

func (s *SQLiteStorage) Read(userId string) (interface{}, error) {
	var jsonData string
	err := s.db.QueryRow("SELECT data FROM user_data WHERE user_id = ?", userId).Scan(&jsonData)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal([]byte(jsonData), &result)
	return result, err
}

func (s *SQLiteStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
