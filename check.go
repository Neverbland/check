// Package check allows data validation of values in different types
package check

import (
	"reflect"
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate(v interface{}) error
}

func Validate(v Validator, val interface{}) Reader {
	return Reader{v.Validate(val)}
}

// NonEmpty check that the value is not a zeroed value depending on its type
type NonEmpty struct{}

// Validate value to not be a zeroed value
func (validator NonEmpty) Validate(v interface{}) error {

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

//Callback validator
type Callback func(interface{}) error

func (validator Callback) Validate(v interface{}) error {
	return validator(v)
}

type Normalizer func(interface{}) (interface{}, error)

//Pass value through Normalizer callbacks chain and pass it to inner validator
type Normalize struct {
	Normalizers []Normalizer
	Validator   Validator
}

func (validator Normalize) Validate(v interface{}) error {

	for _, fn := range validator.Normalizers {
		if val, err := fn(v); err != nil {
			return ValidationErr("normalize", err.Error())
		} else {
			v = val
		}
	}

	return validator.Validator.Validate(v)
}

//Overrides error if present
type CustomError struct {
	Validator
	Error error
}

func (validator CustomError) Validate(v interface{}) error {
	if err := validator.Validator.Validate(v); err != nil {
		return validator.Error
	}
	return nil
}
