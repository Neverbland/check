package check

// Struct allows validation of structs
type Map map[string]Validator

// Validate execute validation using the validators.
func (m Map) Validate(v interface{}) Error {

	val, ok := v.(map[string]interface{})

	if !ok {
		return ValidationErr("map.invalid", "not a map", v)
	}

	e := CompoundValidationError{}

	for fieldname, validator := range m {

		fieldvalue, exists := val[fieldname]

		if _, ok := validator.(Child); ok && !exists {
			e.AddError(fieldname, ValidationErr("child.empty", "no such key", fieldname))
			continue
		}

		if !exists {
			continue
		}

		if err := validator.Validate(fieldvalue); err != nil {
			e.AddError(fieldname, err)
		}
	}

	return e.Value()
}

type Child struct {
	Validator Validator
}

func (validator Child) Validate(v interface{}) Error {
	return validator.Validator.Validate(v)
}
