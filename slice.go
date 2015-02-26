package check

import (
	"reflect"
	"strconv"
)

type Slice struct {
	Inner Validator
}

func (c Slice) Validate(v interface{}) Error {
	err := CompoundValidationError{}

	t := reflect.ValueOf(v)

	switch t.Type().Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < t.Len(); i++ {
			if verr := c.Inner.Validate(t.Index(i).Interface()); verr != nil {
				err.AddError(strconv.Itoa(i), verr)
			}
		}
	default:
		return ValidationErr("slice.type", "not a slice", v)
	}

	return err.Value()
}
