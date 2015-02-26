package check

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type assertion struct {
	v     Validator
	val   interface{}
	valid bool
}

func runAssertions(assertions []assertion, t *testing.T) {
	for _, a := range assertions {
		err := Validate(a.v, a.val)
		if a.valid {
			assert.True(t, err.Empty(), "%#v : %#v expected to be valid. %s", a.v, a.val, err)
		} else {
			assert.False(t, err.Empty(), "%#v: %#v expected to be invalid.", a.v, a.val)
		}

	}
}

type foo struct {
	b bool
}

func TestNonEmpty(t *testing.T) {
	assertions := []assertion{
		{NonEmpty{}, int(1), true},
		{NonEmpty{}, float64(1.0), true},
		{NonEmpty{}, "foo", true},
		{NonEmpty{}, true, true},
		{NonEmpty{}, foo{true}, true},
		{NonEmpty{}, []foo{foo{true}}, true},
		{NonEmpty{}, int(0), false},
		{NonEmpty{}, float64(0.0), false},
		{NonEmpty{}, "", false},
		{NonEmpty{}, false, false},
		{NonEmpty{}, foo{}, false},
		{NonEmpty{}, []foo{}, false},
	}
	runAssertions(assertions, t)
}
