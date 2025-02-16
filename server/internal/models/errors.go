package models

import (
	"errors"
	"fmt"
)

var (
	ErrNoRecord           	   = errors.New("models: no matching record found")
	ErrInvalidCredentials 	   = errors.New("models: invalid credentials")
	ErrDuplicateEmail     	   = errors.New("models: duplicate email")
	ErrUserCheck          	   = errors.New("models: error checking user existence")
	ErrMessageCheck            = errors.New("models: error checking message existence")
	ErrPostCreation            = errors.New("models: error creating post")
	ErrPostIDRetrieval         = errors.New("models: error retrieving post ID")
	ErrDatabaseOperationFailed = errors.New("models: database operation failed")
	ErrRelationshipExists      = errors.New("models: relationship already exists")
)

// Función para facilitar la creación de errores con contexto
func NewErrUserCheck(err error) error {
	return fmt.Errorf("%w: %v", ErrUserCheck, err)
}

func NewErrDatabaseOperationFailed(err error) error {
    return fmt.Errorf("%w: %v", ErrDatabaseOperationFailed, err)
}
