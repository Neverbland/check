package check

import (
	"strconv"
)

type FloatValidator interface {
	ValidateFloat(float64) Error
}
type IntValidator interface {
	ValidateInt(int) Error
}

type NumberValidator interface {
	FloatValidator
	IntValidator
}

type Number []NumberValidator

func (validators Number) Validate(v interface{}) Error {

	switch v.(type) {
	case int:

		intv := make([]IntValidator, len(validators))

		for i, numv := range validators {
			intv[i] = numv.(IntValidator)
		}

		return Integer(intv).Validate(v)
	case float64:
		floatv := make([]FloatValidator, len(validators))

		for i, numv := range validators {
			floatv[i] = numv.(FloatValidator)
		}

		return Float(floatv).Validate(v)
	default:
		return ValidationErr("number.type", "%T not a number", v)
	}

	return nil
}

type Integer []IntValidator

func (validators Integer) Validate(v interface{}) Error {

	val, ok := v.(int)

	if !ok {
		return ValidationErr("integer.type", " %T not an integer", v)
	}

	errs := ErrorCollection{}

	for _, validator := range validators {
		if err := validator.ValidateInt(val); err != nil {
			errs.Add(err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil

}

type Float []FloatValidator

func (validators Float) Validate(v interface{}) Error {

	val, ok := v.(float64)

	if !ok {
		return ValidationErr("float.type", "%T not a float", v)
	}

	errs := ErrorCollection{}

	for _, validator := range validators {
		if err := validator.ValidateFloat(val); err != nil {
			errs.Add(err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil

	return nil
}

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator LowerThan) ValidateInt(v int) Error {

	if validator.Constraint <= float64(v) {
		return ValidationErr("number.lower", "%v is not lower than %v", strconv.Itoa(v), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
	}

	return nil
}

// Validate check value against constraint
func (validator LowerThan) ValidateFloat(v float64) Error {

	if validator.Constraint <= v {
		return ValidationErr("number.lower", "%v is not lower than %v", strconv.FormatFloat(v, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
	}

	return nil
}

// GreaterThan validates that a number must be greater than its value
type GreaterThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator GreaterThan) ValidateInt(v int) Error {

	if validator.Constraint >= float64(v) {
		return ValidationErr("number.greater", "%v is not greater than %v", strconv.Itoa(v), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
	}

	return nil
}

// Validate check value against constraint
func (validator GreaterThan) ValidateFloat(v float64) Error {

	if validator.Constraint >= v {
		return ValidationErr("number.greater", "%v is not greater than %v", strconv.FormatFloat(v, 'f', -1, 64), strconv.FormatFloat(validator.Constraint, 'f', -1, 64))
	}

	return nil
}
