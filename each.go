package check

import (
	"reflect"
)

type Each struct {
	Inner Validator
}

func (c Each) Validate(v interface{}) Error {

	rv := reflect.ValueOf(v)
	rvt := rv.Type()

	switch rvt.Kind() {
	case reflect.Array, reflect.Slice:

		err := make(ErrorCollection, rv.Len())

		for i := 0; i < rv.Len(); i++ {
			if verr := c.Inner.Validate(rv.Index(i).Interface()); verr != nil {
				err[i] = verr
			}
		}

		if len(err) > 0 {
			return err
		}

		return nil
	case reflect.Map:

		if rvt.Key().Kind() != reflect.String {
			return ValidationErr("each.invalid", "%T keys are not strings", v)
		}

		err := ErrorMap{}

		for _, k := range rv.MapKeys() {

			if verr := c.Inner.Validate(rv.MapIndex(k).Interface()); verr != nil {
				err[k.String()] = verr
			}
		}

		if len(err) > 0 {
			return err
		}

		return nil
	}

	return ValidationErr("each.invalid", "%T not a slice/array/map", v)
}
