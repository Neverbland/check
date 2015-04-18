package check

import "time"

type Time []Validator

func (validators Time) Validate(v interface{}) error {
	t, ok := v.(time.Time)
	if !ok {
		return ValidationErr("time.type", "not a time", v)
	}

	return And(validators).Validate(t)
}

// Before check if a time in Value is before the time in Constraint
type Before struct {
	time.Time
}

// Validate check if a time in Value is before the time in Constraint
func (validator Before) Validate(v interface{}) error {

	val := v.(time.Time)

	if !val.Before(validator.Time) {
		return ValidationErr("datetime.before", "%v is not before %v", val.String(), validator.String())
	}
	return nil
}

// After check if a time in Value is before the time in Constraint
type After struct {
	time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (validator After) Validate(v interface{}) error {
	val := v.(time.Time)
	if !val.After(validator.Time) {
		return ValidationErr("datetime.after", "%v is not after %v", val.String(), validator.String())
	}
	return nil
}
