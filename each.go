package check

import (
	"github.com/neverbland/fail"
	"reflect"
)

type Each struct {
	Inner Validator
}

func (c Each) Validate(v interface{}) error {

	rv := reflect.ValueOf(v)
	rvt := rv.Type()

	switch rvt.Kind() {
	case reflect.Array, reflect.Slice:

		errs := make(fail.Collection, rv.Len())

		for i := 0; i < rv.Len(); i++ {
			if err := c.Inner.Validate(rv.Index(i).Interface()); fail.IsError(err) {
				errs[i] = err
			}
		}

		return fail.OrNil(errs)

	case reflect.Map:

		if rvt.Key().Kind() != reflect.String {
			return ValidationErr("each.invalid", "%T keys are not strings", v)
		}

		errs := make(fail.Map, rv.Len())

	for _, k := range rv.MapKeys() {

		if err := c.Inner.Validate(rv.MapIndex(k).Interface()); fail.IsError(err) {
			errs[k.String()] = err
		}
	}

		return fail.OrNil(errs)
	}

	return ValidationErr("each.invalid", "%T not a slice/array/map", v)
}
