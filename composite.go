package check

import (
	"strconv"
)

type Or []Validator

func (or Or) Validate(v interface{}) Error {
	args, ok := v.([]interface{})

	if !ok {
		panic("slice expected")
	}

	if len(or) != len(args) {
		panic("arguments/validators count mismatch")
	}

	e := CompoundValidationError{}

	for i, validator := range or {
		if err := validator.Validate(args[i]); err == nil {
			return nil
		} else {
			e.AddError(strconv.Itoa(i), err)
		}
	}
	return e
}

// Composite allows adding multiple validators to the same value
type And []Validator

// Validate implements Validator
func (and And) Validate(v interface{}) Error {
	args, ok := v.([]interface{})

	if !ok {
		panic("slice expected")
	}

	if len(and) != len(args) {
		panic("arguments/validators count mismatch")
	}

	e := CompoundValidationError{}

	for i, validator := range and {
		if err := validator.Validate(args[i]); err != nil {
			e.AddError(strconv.Itoa(i), err)
		}
	}
	return e.Value()
}
