package services

import (
	"fmt"
)

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

type EntityNotFoundError struct {
	Entity string
	ID     int
}

func (err *EntityNotFoundError) Error() string {
	return fmt.Sprintf("%s with identifier %d not found", err.Entity, err.ID)
}

func NewEntityNotFoundError(entity string, id int) *EntityNotFoundError {
	return &EntityNotFoundError{
		entity,
		id,
	}
}

type InvalidQueryError struct {
	Message string
}

func (err *InvalidQueryError) Error() string {
	return err.Message
}

func NewInvalidQueryError(message string) *InvalidQueryError {
	return &InvalidQueryError{
		message,
	}
}

type UnprocessableBodyError struct {
	Message string
}

func (err *UnprocessableBodyError) Error() string {
	return err.Message
}

func NewUnprocessableBodyError(message string) *UnprocessableBodyError {
	return &UnprocessableBodyError{
		message,
	}
}
