package services

import (
	"fmt"
)

type EntityNotFoundError struct {
	Entity string
	ID     int
}

func (err *EntityNotFoundError) Error() string {
	return fmt.Sprintf(
		"%s with identifier %d not found",
		err.Entity,
		err.ID,
	)
}

func NewEntityNotFoundError(entity string, id int) *EntityNotFoundError {
	return &EntityNotFoundError{
		entity,
		id,
	}
}

type ValidationError struct {
	Message string
}

func (err *ValidationError) Error() string {
	return err.Message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message,
	}
}
