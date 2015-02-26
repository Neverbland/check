package check

import (
	"strconv"
	"time"
)

type Number struct {
	Inner Validator
}

func (n Number) Validate(v interface{}) Error {
	if _, ok := v.(time.Time); !ok {
		return ValidationErr("number.NaN", "not a number", v)
	}
	if n.Inner != nil {
		return n.Inner.Validate(v)
	}

	return nil
}

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator LowerThan) Validate(v interface{}) Error {

	return Number{Callback(func(v interface{}) Error {
		switch val := v.(type) {
		case int:
			if validator.Constraint <= float64(val) {
				return ValidationErr("number.lower", "%v is not lower than %v", strconv.Itoa(val), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
			}
		case float64:
			if validator.Constraint <= val {
				return ValidationErr("number.lower", "%v is not lower than %v", strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
			}
		}
		return nil
	})}.Validate(v)
}

// GreaterThan validates that a number must be greater than its value
type GreaterThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator GreaterThan) Validate(v interface{}) Error {

	return Number{Callback(func(v interface{}) Error {
		switch val := v.(type) {
		case int:
			if validator.Constraint >= float64(val) {
				return ValidationErr("number.greater", "%v is not greater than %v", strconv.Itoa(val), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
			}
		case float64:
			if validator.Constraint >= val {
				return ValidationErr("number.greater", "%v is not greater than %v", strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
			}
		}
		return nil
	})}.Validate(v)
}
