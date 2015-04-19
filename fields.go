package check

import (
	"github.com/neverbland/fail"
	"reflect"
)

// Map allows validation of maps
type Map map[string]Validator

// Validate execute validation using the validators.
func (m Map) Validate(v interface{}) error {

	val, ok := v.(map[string]interface{})

	if !ok {
		return ValidationErr("map.invalid", "not a map")
	}

	errs := fail.Map{}

	for fieldname, validator := range m {

		fieldvalue, exists := val[fieldname]

		if !exists {
			if _, ok := validator.(Child); ok {
				errs[fieldname] = ValidationErr("child.empty", "no such key")
			}
			continue
		}

		if err := fail.OrNil(validator.Validate(fieldvalue)); err != nil {
			errs[fieldname] = err
		}
	}

	return fail.OrNil(errs)
}

type Child []Validator

func (validator Child) Validate(v interface{}) error {
	return And(validator).Validate(v)
}

// Struct allows validation of structs
type Struct map[string]Validator

// Validate execute validation using the validators.
func (s Struct) Validate(v interface{}) error {

	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return ValidationErr("struct.type", "not a struct")
	}

	errs := fail.Map{}

	for fieldname, validator := range s {
		field := val.FieldByName(fieldname)

		if field.Kind() == reflect.Invalid {
			errs[fieldname] = ValidationErr("struct.field", "no such field", fieldname)
			continue
		}

		if err := fail.OrNil(validator.Validate(field.Interface())); err != nil {
			errs[fieldname] = err
		}
	}

	return fail.OrNil(errs)
}
