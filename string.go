package check

import (
	"regexp"
	"strings"
)

type StringValidator interface {
	ValidateString(string) Error
}

type String []StringValidator

func (validators String) Validate(v interface{}) Error {
	s, ok := v.(string)
	if !ok {
		return ValidationErr("string.type", "not a string", v)
	}

	errs := ErrorCollection{}

	for _, validator := range validators {
		if err := validator.ValidateString(s); err != nil {
			errs.Add(err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

//short form to check if value is a string and not empty
type NonEmptyString struct{}

func (validator NonEmptyString) Validate(v interface{}) Error {
	return String{NonEmpty{}}.Validate(v)
}

// MinChar validates that a string must have a length minimum of its constraint
type MinChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MinChar) ValidateString(v string) Error {

	if len(v) < validator.Constraint {
		return ValidationErr("string.min", "too short, minimum %v characters", validator.Constraint)
	}
	return nil
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MaxChar) ValidateString(v string) Error {
	if len(v) > validator.Constraint {
		return ValidationErr("string.max", "too long, minimum %v characters", validator.Constraint)
	}
	return nil
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct{}

// Validate email addresses
func (validator Email) ValidateString(v string) Error {

	if !strings.Contains(v, "@") || string(v[0]) == "@" || string(v[len(v) - 1]) == "@" {
		return ValidationErr("string.email", "'%v' is an invalid email address", v)
	}

	return nil
}

func Regexp(re string) RegexpValidator {
	return RegexpValidator{regexp.MustCompile(re)}
}

// Regex allow validation using regular expressions
type RegexpValidator struct {
	*regexp.Regexp
}

// Validate using regex
func (validator RegexpValidator) ValidateString(v string) Error {

	if !validator.MatchString(v) {
		return ValidationErr("string.regex", "'%v' does not match '%v'", v, validator.String())
	}

	return nil
}

// UUID verify a string in the UUID format xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
type UUID struct{}

// Validate checks a string as correct UUID format
func (validator UUID) ValidateString(v string) Error {

	regex := regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")

	if !regex.MatchString(v) {
		return ValidationErr("string.uuid", "'%v' is an invalid uuid", v)
	}
	return nil
}
