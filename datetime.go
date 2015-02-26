package check

import "time"

type Time struct {
	Inner Validator
}

func (t Time) Validate(v interface{}) Error {
	if _, ok := v.(time.Time); !ok {
		return ValidationErr("time.type", "not a type", v)
	}
	if t.Inner != nil {
		return t.Inner.Validate(v)
	}

	return nil
}

// Before check if a time in Value is before the time in Constraint
type Before struct {
	Constraint time.Time
}

// Validate check if a time in Value is before the time in Constraint
func (validator Before) Validate(v interface{}) Error {

	return Time{Callback(func(v interface{}) Error {
		if !v.(time.Time).Before(validator.Constraint) {
			return ValidationErr("datetime.before", "%v is not before %v", v.(time.Time).String(), validator.Constraint.String())
		}
		return nil
	})}.Validate(v)

}

// After check if a time in Value is before the time in Constraint
type After struct {
	Constraint time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (validator After) Validate(v interface{}) Error {
	return Time{Callback(func(v interface{}) Error {
		if !v.(time.Time).After(validator.Constraint) {
			return ValidationErr("datetime.after", "%v is not after %v", v.(time.Time).String(), validator.Constraint.String())
		}
		return nil
	})}.Validate(v)

}
