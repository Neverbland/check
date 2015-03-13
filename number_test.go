package check

import "testing"

func TestNumberValidators(t *testing.T) {
	assertions := []assertion{
		{Number{LowerThan{2}}, 1, true},
		{Number{LowerThan{1}}, 2, false},
		{Number{LowerThan{2}}, 2, false},
		{Number{GreaterThan{1}}, 2, true},
		{Number{GreaterThan{2}}, 1, false},
		{Number{GreaterThan{2}}, 2, false},
	}

	runAssertions(assertions, t)
}
