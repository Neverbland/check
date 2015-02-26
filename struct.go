package check

import (
	"reflect"
)

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

	e := CompoundValidationError{}

	for fieldname, validator := range s {
		field := val.FieldByName(fieldname)

		if field.Kind() == reflect.Invalid {
			e.AddError(fieldname, ValidationErr("struct.field", "missing field", fieldname))
		}

		if err := validator.Validate(field.Interface()); err != nil {
			e.AddError(fieldname, err)
		}
	}

	return e.Value()
}
