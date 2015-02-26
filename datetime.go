package check

import "time"

type Time []TimeValidator

type TimeValidator interface {
	ValidateTime(time.Time) Error
}

func (validators Time) Validate(v interface{}) Error {
	t, ok := v.(time.Time)
	if !ok {
		return ValidationErr("time.type", "not a time", v)
	}

	errs := ErrorCollection{}

	for _, validator := range validators {
		if err := validator.ValidateTime(t); err != nil {
			errs.Add(err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil

}

// Before check if a time in Value is before the time in Constraint
type Before struct {
	Constraint time.Time
}

// Validate check if a time in Value is before the time in Constraint
func (validator Before) ValidateTime(v time.Time) Error {

	if !v.Before(validator.Constraint) {
		return ValidationErr("datetime.before", "%v is not before %v", v.String(), validator.Constraint.String())
	}
	return nil
}

// After check if a time in Value is before the time in Constraint
type After struct {
	Constraint time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (validator After) ValidateTime(v time.Time) Error {
	if !v.After(validator.Constraint) {
		return ValidationErr("datetime.after", "%v is not after %v", v.String(), validator.Constraint.String())
	}
	return nil
}
