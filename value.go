package check

import (
	"reflect"
)

type Enum []interface{}

func (validator Enum) Validate(v interface{}) error {
	if len(validator) == 0 {
		return nil
	}

	for _, val := range validator {
		if err := (Equal{val}).Validate(v); err == nil {
			return nil
		}
	}

	return ValidationErr("enum", "must be equal to one of %v. Given %v", validator, v)
}

type Equal struct {
	Value interface{}
}

func (validator Equal) Validate(v interface{}) error {

	if !reflect.DeepEqual(v, validator.Value) {
		return ValidationErr("equal", "must be equal to %v. Given %v", validator.Value, v)
	}

	return nil
}
