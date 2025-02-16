package persistence

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Implements IPersistance from godemlia library
type SqliteDb struct {
	db *sql.DB
}

// NewSqliteDb initializes a new SQLite database connection.
func NewSqliteDb(dbPath string) *SqliteDb {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS kv_store (
		key TEXT PRIMARY KEY,
		value BLOB
	);`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return &SqliteDb{db: db}
}
func (s *SqliteDb) Handle(action string, path string, data *[]byte) (*[]byte, error) {
	return nil, nil
}

// inserts a key-value pair into the SQLite database.
func (s *SqliteDb) Create(key []byte, data *[]byte) error {
	_, err := s.db.Exec("INSERT INTO kv_store (key, data) VALUES (?, ?)", key, *data)
	return err
}

// GetValue retrieves a value from the database by key.
func (s *SqliteDb) GetValue(key string) ([]byte, error) {
	var value []byte
	err := s.db.QueryRow("SELECT value FROM kv_store WHERE key = ?", key).Scan(&value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// Close closes the database connection.
func (s *SqliteDb) Close() {
	s.db.Close()
}

func (s *SqliteDb) Delete(key []byte) error {
	// Sin implementar aun

	return nil
}

func (s *SqliteDb) GetKeys() [][]byte {
	// keys, err := s.db.(context.TODO(), "*").Result()
	// result := [][]byte{}
	// if err != nil {
	// 	return result
	// }
	// for _, key := range keys {
	// 	keyBytes := base58.Decode(key)
	// 	result = append(result, keyBytes)
	// }
	// return result
	return nil
}
