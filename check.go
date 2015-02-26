// Package check allows data validation of values in different types
package check

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	// ErrorMessages contains default error messages
	ErrorMessages = map[string]string{}
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate(v interface{}) Error
}

// Error is the default validation error. The Params() method returns the params
// to be used in error messages
type Error interface {
	Error() string
	ErrorMap() map[string][]Error
	Compound() bool
	Value() Error
}

type CompoundValidationError map[string][]Error

func (em CompoundValidationError) AddError(field string, err Error) {
	if _, ok := em[field]; !ok {
		em[field] = []Error{}
	}
	em[field] = append(em[field], err)
}

func (em CompoundValidationError) IsEmpty() bool {
	return len(em) == 0
}

func (em CompoundValidationError) Error() string {
	errs := make([]string, 0, len(em))

	for field, fielderrs := range em {
		ferrs := make([]string, 0, len(fielderrs))

		for _, err := range fielderrs {
			ferrs = append(ferrs, err.Error())
		}

		errs = append(errs, fmt.Sprintf("'%s':\n\t %s", field, strings.Join(ferrs, ";")))
	}

	return strings.Join(errs, "\n")
}

// ErrorMap returns the error map
func (em CompoundValidationError) ErrorMap() map[string][]Error { return map[string][]Error(em) }

func (em CompoundValidationError) Compound() bool { return true }

func (em CompoundValidationError) Value() Error {
	if em.IsEmpty() {
		return nil
	}

	return em
}

func ValidationErr(err, message string, params ...interface{}) *ValidationError {
	return &ValidationError{err, message, params}
}

// ValidationError implements Error
type ValidationError struct {
	Name    string
	Message string
	Params  []interface{}
}

func (e ValidationError) Error() string {

	if e.Message == "" {
		return fmt.Sprintf(e.Message, e.Params)
	} else {
		return fmt.Sprintf(e.Message, e.Params)
	}

}

// ErrorMap returns the error map
func (e ValidationError) ErrorMap() map[string][]Error {
	return map[string][]Error{"error": []Error{e}}
}

func (e ValidationError) Compound() bool { return false }

func (e ValidationError) Value() Error { return e }

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

type Callback func(v interface{}) Error

func (c Callback) Validate(v interface{}) Error { return c(v) }
