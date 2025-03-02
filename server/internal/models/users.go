package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(user *dto.CreateUserDTO) (*User, error) {
	stmt := `INSERT INTO user (user_id, username, email, password_hash) VALUES (?, ?, ?, ?)`

	var hashedPassword string
	if user.PasswordHash == "" {
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if err != nil {
			return nil, err
		}
		hashedPassword = string(hashedPasswordBytes)
	} else {
		hashedPassword = user.PasswordHash
	}

	_, err := m.DB.Exec(stmt, user.UserID, user.UserName, user.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return m.Get(user.UserID)
}

func (m *UserModel) Authenticate(email string, password string) (*User, error) {

	u := &User{}

	stmt := "SELECT user_id, username, email, password_hash, created_at, updated_at  FROM user WHERE email = ?"
	err := m.DB.QueryRow(stmt, email).Scan(&u.UserID, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		} else {
			return nil, ErrInvalidCredentials
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	return u, nil
}

func (m *UserModel) Get(user_id string) (*User, error) {
	stmt := `SELECT user_id, username, email, password_hash, created_at, updated_at 
			FROM user
			WHERE user_id = ?
			`

	row := m.DB.QueryRow(stmt, user_id)

	u := &User{}

	err := row.Scan(&u.UserID, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord

		} else {
			return nil, err
		}
	}

	return u, nil
}

func (m *UserModel) List() ([]*User, error) { //pueden haber otras condiciones. Por ahora solo se listaran 10 por orden de id
	stmt := `SELECT user_id, username, email, password_hash, created_at, updated_at 
			FROM user
			ORDER BY user_id DESC 
			LIMIT 10
			`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		u := &User{}

		err = rows.Scan(&u.UserID, &u.Username, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) Update(ctx context.Context, user *User) error {
	query := `
	 UPDATE user
	 SET 
	 username = ?, 
	 email = ?, 
	 updated_at = ?
	 WHERE user_id = ?
	`

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Username, user.Email, time.Now(), user.UserID)
	if err != nil {
		return fmt.Errorf("error executing update statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("user not found") //O puedes usar tu ErrNoRecord si lo tienes definido
	}

	return nil
}

func (m *UserModel) Delete(ctx context.Context, userID string) error {
	// Iniciar una transacción
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback() // Rollback en caso de error

	// Eliminar registros dependientes (por ejemplo, mensajes y seguidores)
	_, err = tx.ExecContext(ctx, `
        DELETE FROM post WHERE user_id = ?
    `, userID)
	if err != nil {
		return fmt.Errorf("error deleting dependent messages: %w", err)
	}

	_, err = tx.ExecContext(ctx, `
        DELETE FROM followers WHERE user_id = ? OR followee_id = ?
    `, userID, userID)
	if err != nil {
		return fmt.Errorf("error deleting dependent followers: %w", err)
	}

	// Eliminar el usuario
	result, err := tx.ExecContext(ctx, `
        DELETE FROM user WHERE user_id = ?
    `, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	// Verificar si se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNoRecord
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT user_id, created_at, name, email, password_hash FROM user WHERE email = ?`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.UserID,
		&user.CreatedAt,
		// &user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, err
		default:
			return nil, err
		}
	}
	return &user, nil
}
