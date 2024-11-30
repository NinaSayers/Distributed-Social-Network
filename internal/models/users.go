package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(username, email, password string) (int, error) {
	stmt := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec(stmt, username, email, hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {

	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM users WHERE email = ?"
	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, ErrInvalidCredentials
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return int(id), nil
}

func (m *UserModel) Get(user_id int) (*User, error) {
	stmt := `SELECT user_id, username, email, password_hash, created_at, updated_at 
			FROM users
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
			FROM users
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
	 UPDATE users
	 SET 
	 username = ?, 
	 email = ?, 
	 password_hash = ?, 
	 updated_at = ?
	 WHERE user_id = ?
	`

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Username, user.Email, user.PasswordHash, time.Now(), user.UserID)
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

func (m *UserModel) Delete(ctx context.Context, userID int) error {
	query := `
 			DELETE FROM users
 			WHERE user_id = ?
			`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error preparing delete statement: %w", err)
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, userID)
	if err != nil {
		return fmt.Errorf("error executing delete statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("user not found") // or use your custom ErrNoRecord if defined
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT user_id, created_at, name, email, password_hash FROM users WHERE email = ?`
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
