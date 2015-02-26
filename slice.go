package check

import (
	"reflect"
)

type Slice struct {
	Inner Validator
}

func (c Slice) Validate(v interface{}) Error {

	t := reflect.ValueOf(v)

	switch t.Type().Kind() {
	case reflect.Array, reflect.Slice:

		err := make(ErrorCollection, t.Len())

		for i := 0; i < t.Len(); i++ {
			if verr := c.Inner.Validate(t.Index(i).Interface()); verr != nil {
				err[i] = verr
			}
		}

		if len(err) > 0 {
			return err
		}

		return nil

	}

	return ValidationErr("slice.invalid", "%T not a slice", v)
}
