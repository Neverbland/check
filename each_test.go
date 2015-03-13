package check

import "testing"

func TestEach(t *testing.T) {
	assertions := []assertion{
		{Each{And{Number{}, NonEmpty{}}}, []int{1, 2, 3}, true},
		{Each{And{Number{}, NonEmpty{}}}, []int{1, 2, 3, 0}, false},
		{Each{And{Number{}, NonEmpty{}}}, []int{}, true},

		{Each{And{Number{}, NonEmpty{}}}, map[string]int{"one": 1, "two": 2, "three": 3}, true},
		{Each{And{Number{}, NonEmpty{}}}, map[string]int{"1": 1, "2": 2, "3": 3, "0": 0}, false},
		{Each{And{Number{}, NonEmpty{}}}, map[string]int{}, true},
		{Each{And{Number{}, NonEmpty{}}}, map[int]int{}, false},

		{Each{}, 0, false},
	}
	runAssertions(assertions, t)
}
