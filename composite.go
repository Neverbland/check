package check

type Or []Validator

func (validators Or) Validate(v interface{}) Error {

	e := ErrorCollection{}

	for _, validator := range validators {
		if err := validator.Validate(v); err == nil {
			return nil
		} else {
			e.Add(err)
		}
	}

	return e
}

type And []Validator

// Validate implements Validator
func (validators And) Validate(v interface{}) Error {

	e := ErrorCollection{}

	for _, validator := range validators {
		if err := validator.Validate(v); err != nil {
			e.Add(err)
		}
	}

	if len(e) > 0 {
		return e
	}

	return nil
}
