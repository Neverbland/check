package check

import "testing"

func TestSlice(t *testing.T) {
	assertions := []assertion{
		{Slice{And{Number{}, NonEmpty{}}}, []int{1, 2, 3}, true},
		{Slice{And{Number{}, NonEmpty{}}}, []int{1, 2, 3, 0}, false},
		{Slice{And{Number{}, NonEmpty{}}}, []int{}, true},
		{Slice{}, 0, false},
	}
	runAssertions(assertions, t)
}
