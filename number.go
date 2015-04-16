package check

import (
	"strconv"
)

type Number []Validator

func (validators Number) Validate(v interface{}) Error {

	switch v.(type) {
	case int, float64:
		return And(validators).Validate(v)

	default:
		return ValidationErr("number.type", "%T not a number", v)
	}

	return nil
}

type Integer []Validator

func (validators Integer) Validate(v interface{}) Error {

	val, ok := v.(int)

	if !ok {
		return ValidationErr("integer.type", " %T not an integer", v)
	}

	return And(validators).Validate(val)

}

type Float []Validator

func (validators Float) Validate(v interface{}) Error {

	val, ok := v.(float64)

	if !ok {
		return ValidationErr("float.type", "%T not a float", v)
	}

	return And(validators).Validate(val)
}

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
}

func (validator LowerThan) Validate(v interface{}) Error {

	switch val := v.(type) {
	case float64:
		if validator.Constraint <= val {
			return ValidationErr("number.lower", "%v is not lower than %v", strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
		}
	case int:
		if validator.Constraint <= float64(val) {
			return ValidationErr("number.lower", "%v is not lower than %v", strconv.Itoa(val), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
		}
	}

	return nil
}

// GreaterThan validates that a number must be greater than its value
type GreaterThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator GreaterThan) Validate(v interface{}) Error {

	switch val := v.(type) {
	case float64:
		if validator.Constraint >= val {
			return ValidationErr("number.greater", "%v is not greater than %v", strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
		}
	case int:
		if validator.Constraint >= float64(val) {
			return ValidationErr("number.greater", "%v is not greater than %v", strconv.Itoa(val), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
		}
	}

	return nil
}
