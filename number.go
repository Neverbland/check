package check

import (
	"strconv"
)

type Number []Validator

func (validators Number) Validate(v interface{}) error {

	switch v.(type) {
	case int, float64:
		return And(validators).Validate(v)

	default:
		return ValidationErr("number.type", "%T not a number", v)
	}

	return nil
}

type Integer []Validator

func (validators Integer) Validate(v interface{}) error {

	val, ok := v.(int)

	if !ok {
		return ValidationErr("integer.type", " %T not an integer", v)
	}

	return And(validators).Validate(val)

}

type Float []Validator

func (validators Float) Validate(v interface{}) error {

	val, ok := v.(float64)

	if !ok {
		return ValidationErr("float.type", "%T not a float", v)
	}

	return And(validators).Validate(val)
}

func LowerThan(v float64) LowerThanNumber {
	return LowerThanNumber{v, false}
}

func LowerThanOrEqual(v float64) LowerThanNumber {
	return LowerThanNumber{v, true}
}

// LowerThanNumber validates that a number must be lower than its value
type LowerThanNumber struct {
	Constraint float64
	Inclusive  bool
}

func (validator LowerThanNumber) Validate(v interface{}) error {

	switch val := v.(type) {
	case float64:

		diff := validator.Constraint - val

		if (validator.Inclusive && diff == 0) || diff > 0 {
			return nil
		}

		return ValidationErr("number.lower", "%v is not lower than %v", strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))

	case int:

		diff := validator.Constraint - float64(val)

		if (validator.Inclusive && diff == 0) || diff > 0 {
			return nil
		}

		return ValidationErr("number.lower", "%v is not lower than %v", strconv.Itoa(val), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
	}

	return nil
}

func GreaterThan(v float64) GreaterThanNumber {
	return GreaterThanNumber{v, false}
}

func GreaterThanOrEqual(v float64) GreaterThanNumber {
	return GreaterThanNumber{v, true}
}

// GreaterThanNumber validates that a number must be greater than its value
type GreaterThanNumber struct {
	Constraint float64
	Inclusive  bool
}

// Validate check value against constraint
func (validator GreaterThanNumber) Validate(v interface{}) error {

	switch val := v.(type) {
	case float64:

		diff := val - validator.Constraint

		if (validator.Inclusive && diff == 0) || diff > 0 {
			return nil
		}

		return ValidationErr("number.greater", "%v is not greater than %v", strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))

	case int:

		diff := float64(val) - validator.Constraint

		if (validator.Inclusive && diff == 0) || diff > 0 {
			return nil
		}

		return ValidationErr("number.greater", "%v is not greater than %v", strconv.Itoa(val), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
	}

	return nil
}

func Between(min, max float64) And {
	return And{
		GreaterThan(min),
		LowerThan(max),
	}
}

func BetweenInclusive(min, max float64) And {
	return And{
		GreaterThanOrEqual(min),
		LowerThanOrEqual(max),
	}
}
