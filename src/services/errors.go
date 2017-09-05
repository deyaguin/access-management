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

type EntityCreateError struct {
	Message string
}

func (err *EntityCreateError) Error() string {
	return err.Message
}

func NewEntityCreateError(message string) *EntityCreateError {
	return &EntityCreateError{
		message,
	}
}

type EntityRemoveError struct {
	Message string
}

func (err *EntityRemoveError) Error() string {
	return err.Message
}

func NewEntityRemoveError(message string) *EntityRemoveError {
	return &EntityRemoveError{
		message,
	}
}

type EntityUpdateError struct {
	Message string
}

func (err *EntityUpdateError) Error() string {
	return err.Message
}

func NewEntityUpdateError(message string) *EntityUpdateError {
	return &EntityUpdateError{
		message,
	}
}

type GetEntitiesError struct {
	Message string
}

func (err *GetEntitiesError) Error() string {
	return err.Message
}

func NewGetEntitiesError(message string) *GetEntitiesError {
	return &GetEntitiesError{
		message,
	}
}
