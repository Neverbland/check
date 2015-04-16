package check

import "testing"

func TestNumberValidators(t *testing.T) {
	assertions := []assertion{
		{LowerThan(2), 1, true},
		{LowerThan(1), 2, false},
		{LowerThan(2), 2, false},
		{LowerThanOrEqual(2), 2, true},
		{GreaterThan(1), 2, true},
		{GreaterThan(2), 1, false},
		{GreaterThan(2), 2, false},
		{GreaterThanOrEqual(2), 2, true},
		{Between(0, 2), 1, true},
		{Between(0, 2), 0, false},
		{Between(0, 2), 2, false},
		{BetweenInclusive(0, 2), 0, true},
		{BetweenInclusive(0, 2), 2, true},
	}

	runAssertions(assertions, t)
}
