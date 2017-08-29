package api

import "fmt"

type InvalidQueryError struct {
	ParamName  string
	ParamValue string
}

func (err *InvalidQueryError) Error() string {
	return fmt.Sprintf("Value %s of param %s is not valid", err.ParamValue, err.ParamName)
}

func NewInvalidQueryError(paramName, paramValue string) *InvalidQueryError {
	return &InvalidQueryError{
		paramName,
		paramValue,
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
