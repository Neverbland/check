// Package check allows data validation of values in different types
package check

import (
	"reflect"
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate(v interface{}) Error
}

func Validate(v Validator, val interface{}) ErrorReader {
	return ErrorReader{v.Validate(val)}
}

// NonEmpty check that the value is not a zeroed value depending on its type
type NonEmpty struct{}

// Validate value to not be a zeroed value, return error and empty slice of strings
func (validator NonEmpty) Validate(v interface{}) Error {

	err := ValidationErr("empty", "value cannot be empty")

	if v == nil {
		return err
	}

	t := reflect.TypeOf(v)

	switch t.Kind() {
	default:
		if reflect.DeepEqual(reflect.Zero(t).Interface(), v) {
			return err
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if reflect.ValueOf(v).Len() == 0 {
			return err
		}
	}

	return nil
}

// Validate value to not be a zeroed value, return error and empty slice of strings
func (validator NonEmpty) ValidateString(v string) Error {

	if len(v) == 0 {
		return ValidationErr("empty", "value cannot be empty")
	}

	return nil
}

//Callback validator
type Callback func(interface{}) Error

func (validator Callback) Validate(v interface{}) Error {
	return validator(v)
}
