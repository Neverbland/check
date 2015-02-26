package check

import (
	"regexp"
	"strings"
)

type String struct {
	Inner Validator
}

func (s String) Validate(v interface{}) Error {
	if _, ok := v.(string); !ok {
		return ValidationErr("string.type", "not a string", v)
	}
	if s.Inner != nil {
		return s.Inner.Validate(v)
	}

	return nil
}

type NonEmptyString struct{}

func (s NonEmptyString) Validate(v interface{}) Error {
	return String{NonEmpty{}}.Validate(v)
}

// MinChar validates that a string must have a length minimum of its constraint
type MinChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MinChar) Validate(v interface{}) Error {

	return String{Callback(func(v interface{}) Error {
		if len(v.(string)) < validator.Constraint {
			return ValidationErr("string.min", "too short, minimum %v characters", validator.Constraint)
		}
		return nil
	})}.Validate(v)
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MaxChar) Validate(v interface{}) Error {

	return String{Callback(func(v interface{}) Error {
		if len(v.(string)) > validator.Constraint {
			return ValidationErr("string.max", "too long, minimum %v characters", validator.Constraint)
		}
		return nil
	})}.Validate(v)
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct{}

// Validate email addresses
func (validator Email) Validate(v interface{}) Error {

	return String{Callback(func(v interface{}) Error {
		if !strings.Contains(v.(string), "@") || string(v.(string)[0]) == "@" || string(v.(string)[len(v.(string))-1]) == "@" {
			return ValidationErr("string.email", "'%v' is an invalid email address", v)
		}
		return nil

	})}.Validate(v)
}

// Regex allow validation usig regular expressions
type Regex struct {
	Constraint string
}

// Validate using regex
func (validator Regex) Validate(v interface{}) Error {

	return String{Callback(func(v interface{}) Error {
		regex, err := regexp.Compile(validator.Constraint)
		if err != nil {
			panic(err)
		}

		if !regex.MatchString(v.(string)) {
			return ValidationErr("string.regex", "'%v' does not match '%v'", v, validator.Constraint)
		}

		return nil
	})}.Validate(v)
}

// UUID verify a string in the UUID format xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
type UUID struct{}

// Validate checks a string as correct UUID format
func (validator UUID) Validate(v interface{}) Error {

	return String{Callback(func(v interface{}) Error {
		regex := regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")

		if !regex.MatchString(v.(string)) {
			return ValidationErr("string.uuid", "'%v' is an invalid uuid", v)
		}
		return nil

	})}.Validate(v)

}
