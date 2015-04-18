package check

import (
	"github.com/neverbland/fail"
)

type Or []Validator

func (validators Or) Validate(v interface{}) error {

	e := fail.List{}

	for _, validator := range validators {
		if err := fail.OrNil(validator.Validate(v)); err == nil {
			return nil
		} else {
			e.Append(err)
		}
	}

	return e
}

type And []Validator

// Validate implements Validator
func (validators And) Validate(v interface{}) error {

	e := fail.List{}

	for _, validator := range validators {
		if err := fail.OrNil(validator.Validate(v)); err != nil {
			e.Append(err)
		}
	}

	return fail.OrNil(e)
}
