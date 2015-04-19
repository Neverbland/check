package check

import (
	"regexp"
	"strings"
)

type String []Validator

func (validators String) Validate(v interface{}) error {
	strv, ok := v.(string)
	if !ok {
		return ValidationErr("string.type", "not a string")
	}

	return And(validators).Validate(strv)
}

//short form to check if value is a string and not empty
type NonEmptyString struct{}

func (validator NonEmptyString) Validate(v interface{}) error {
	return String{NonEmpty{}}.Validate(v)
}

// MinChar validates that a string must have a length minimum of its constraint
type MinChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MinChar) Validate(v interface{}) error {

	if len(v.(string)) < validator.Constraint {
		return ValidationErr("string.min", "too short, minimum %v characters", validator.Constraint)
	}
	return nil
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MaxChar) Validate(v interface{}) error {
	if len(v.(string)) > validator.Constraint {
		return ValidationErr("string.max", "too long, minimum %v characters", validator.Constraint)
	}
	return nil
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct{}

// Validate email addresses
func (validator Email) Validate(v interface{}) error {

	str := v.(string)

	if !strings.Contains(str, "@") || string(str[0]) == "@" || string(str[len(str)-1]) == "@" {
		return ValidationErr("string.email", "'%v' is an invalid email address", str)
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
func (validator RegexpValidator) Validate(v interface{}) error {

	if !validator.MatchString(v.(string)) {
		return ValidationErr("string.regex", "'%v' does not match '%v'", v, validator.String())
	}

	return nil
}

var UUIDRe = regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")

// UUID verify a string in the UUID format xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
type UUID struct{}

// Validate checks a string as correct UUID format
func (validator UUID) Validate(v interface{}) error {

	if !UUIDRe.MatchString(v.(string)) {
		return ValidationErr("string.uuid", "'%v' is an invalid uuid", v)
	}
	return nil
}
