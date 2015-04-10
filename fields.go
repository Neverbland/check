package check

import (
	"reflect"
)

// Map allows validation of maps
type Map map[string]Validator

// Validate execute validation using the validators.
func (m Map) Validate(v interface{}) Error {

	val, ok := v.(map[string]interface{})

	if !ok {
		return ValidationErr("map.invalid", "not a map", v)
	}

	e := ErrorMap{}

	for fieldname, validator := range m {

		fieldvalue, exists := val[fieldname]

		if !exists {
			if _, ok := validator.(Child); ok {
				e[fieldname] = ValidationErr("child.empty", "no such key", fieldname)
			}
			continue
		}

		if err := validator.Validate(fieldvalue); err != nil {
			e[fieldname] = err
		}
	}

	if len(e) > 0 {
		return e
	}

	return nil
}

type Child []Validator

func (validator Child) Validate(v interface{}) Error {
	return And(validator).Validate(v)
}

// Struct allows validation of structs
type Struct map[string]Validator

// Validate execute validation using the validators.
func (s Struct) Validate(v interface{}) Error {

	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return ValidationErr("struct.type", "not a struct")
	}

	e := ErrorMap{}

	for fieldname, validator := range s {
		field := val.FieldByName(fieldname)

		if field.Kind() == reflect.Invalid {
			e[fieldname] = ValidationErr("struct.field", "missing field", fieldname)
			continue
		}

		if err := validator.Validate(field.Interface()); err != nil {
			e[fieldname] = err
		}
	}

	if len(e) > 0 {
		return e
	}

	return e
}
