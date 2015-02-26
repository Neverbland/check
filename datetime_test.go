package check

import (
	"testing"
	"time"
)

func TestDatetimeValidators(t *testing.T) {
	assertions := []assertion{
		{
			Time{Before{time.Date(2014, time.January, 1, 1, 0, 0, 0, time.UTC)}},
			time.Date(2013, time.January, 1, 1, 0, 0, 0, time.UTC),
			true,
		},
		{
			Time{After{time.Date(2014, time.January, 1, 1, 0, 0, 0, time.UTC)}},
			time.Date(2015, time.January, 1, 1, 0, 0, 0, time.UTC),
			true,
		},
	}

	runAssertions(assertions, t)
}
