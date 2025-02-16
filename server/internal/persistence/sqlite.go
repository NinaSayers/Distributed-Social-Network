package persistence

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
	"github.com/NinaSayers/Distributed-Social-Network/server/internal/models"
	"github.com/jbenet/go-base58"
	_ "github.com/mattn/go-sqlite3"
)

// Implements IPersistance from godemlia library
type SqliteDb struct {
	db *sql.DB
	models.Models
}

// NewSqliteDb initializes a new SQLite database connection.
func NewSqliteDb(dbPath string, script string) *SqliteDb {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	// Create table if not exists
	sqlBytes, err := os.ReadFile(script)
	if err != nil {
		log.Fatalf("Failed to load sql script: %v", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return &SqliteDb{db: db, Models: models.NewModels(db)}
}

func (s *SqliteDb) MainEntity() string {
	return "user"
}

func (s *SqliteDb) Handle(action string, path string, data *[]byte) (*[]byte, error) {
	return nil, nil
}

// inserts a key-value pair into the SQLite database.
func (s *SqliteDb) Store(entity string, id []byte, data *[]byte) (*[]byte, error) {
	switch entity {
	case "user":
		var user dto.CreateUserDTO
		err := json.Unmarshal(*data, &user)
		if err != nil {
			return nil, err
		}
		user.UserID = base58.Encode(id)

		_, err = s.User.Create(&user)
		if err != nil {
			return nil, err
		}
	case "post":
		var message dto.CreatePostDTO
		err := json.Unmarshal(*data, &message)
		if err != nil {
			return nil, err
		}
		message.PostID = base58.Encode(id)

		_, err = s.Post.Create(&message)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
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

func (s *SqliteDb) Read(entity string, id []byte) (*[]byte, error) {
	switch entity {
	case "user":
		user, err := s.User.Get(base58.Encode(id))
		if err != nil {
			return nil, err
		}
		response, err := json.Marshal(user)
		return &response, err

	case "post":
		message, err := s.Post.Get(base58.Encode(id))
		if err != nil {
			return nil, err
		}
		response, err := json.Marshal(message)
		return &response, err
	}
	return nil, errors.New("entity not found")
}

func (s *SqliteDb) GetKeys() map[string][][]byte {
	keys := make(map[string][][]byte)

	// Get all table names
	rows, err := s.db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Fatalf("Failed to get table names: %v", err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		fmt.Println("Table name: ", tableName)
		if err != nil {
			fmt.Printf("Failed to scan table name: %v \n", err)
			// log.Fatalf("Failed to scan table name: %v", err)
			continue
		}

		// Get all keys from the table
		tableKeys, err := s.getTableKeys(tableName)
		if err != nil {
			fmt.Printf("Failed to get keys from table %s: %v \n", tableName, err)
			// log.Fatalf("Failed to get keys from table %s: %v", tableName, err)
			continue
		}
		keys[tableName] = tableKeys
	}

	return keys
}

func (s *SqliteDb) getTableKeys(tableName string) ([][]byte, error) {
	query := "SELECT " + tableName + "_id FROM " + tableName
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var key []byte
	var tableKeys [][]byte
	for rows.Next() {
		err := rows.Scan(&key)
		if err != nil {
			return nil, err
		}
		tableKeys = append(tableKeys, key)
	}

	return tableKeys, nil
}
