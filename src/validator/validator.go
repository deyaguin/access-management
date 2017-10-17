package validator

import (
	"reflect"

	"gopkg.in/validator.v2"
)

func InitValidator() {
	validator.SetValidationFunc("isBool", isBool)
}

func isBool(value interface{}, param string) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Bool {
		return validator.ErrUnsupported
	}
	return nil
}
