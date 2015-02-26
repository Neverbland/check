package check

import "testing"

func TestComposite(t *testing.T) {
	assertions := []assertion{
		{And{Number{}, Integer{}}, int(1), true},
		{And{Number{}, Integer{}}, float64(1), false},

		{Or{Number{}, Integer{}}, int(1), true},
		{Or{Number{}, Integer{}}, float64(1), true},
		{Or{Integer{}, String{}}, float64(1), false},
	}
	runAssertions(assertions, t)
}
