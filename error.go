package check

import (
	"encoding/json"
	"fmt"
	"github.com/neverbland/access"
)

// Error is the default validation error. The Params() method returns the params
// to be used in error messages
type Error interface {
	Error() string
	Count() int
	Empty() bool
	Name() string
}

func ValidationErr(name, message string, params ...interface{}) ValidationError {
	return ValidationError{name, message, params}
}

// ValidationError implements Error
type ValidationError struct {
	name    string
	Message string
	Params  []interface{}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf(e.Message, e.Params...)
}

func (e ValidationError) Count() int   { return 1 }
func (e ValidationError) Empty() bool  { return false }
func (e ValidationError) Name() string { return e.name }

func (e ValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Error())
}

type ErrorCollection []Error

func (e ErrorCollection) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}
	str := ""
	for i, err := range e {
		str += fmt.Sprintf("%d: %s \n", i, err)
	}

	return str
}

func (e ErrorCollection) Count() int {
	count := 0
	for _, err := range e {
		count += err.Count()
	}

	return count
}

func (e ErrorCollection) Empty() bool {
	for _, err := range e {
		if !err.Empty() {
			return false
		}
	}

	return true
}

func (e ErrorCollection) Name() string { return "collection" }

func (e *ErrorCollection) Add(err Error) {
	*e = append(*e, err)
}

type ErrorMap map[string]Error

func (e ErrorMap) Error() string {
	str := ""
	for name, err := range e {
		str += fmt.Sprintf("'%s': %s\n", name, err)
	}

	return str
}

func (e ErrorMap) Count() int {
	count := 0
	for _, err := range e {
		count += err.Count()
	}

	return count
}

func (e ErrorMap) Empty() bool {
	for _, err := range e {
		if !err.Empty() {
			return false
		}
	}

	return true
}

func (e ErrorMap) Name() string { return "map" }

type ErrorReader struct {
	Err Error
}

func (e ErrorReader) Get(path ...string) ErrorReader {

	if len(path) == 0 {
		return e
	}

	var err Error = nil

	if t, ok := access.MustRead(path[0], e.Err).(Error); ok && t != nil {
		err = t
	}

	for {

		if ec, ok := err.(ErrorCollection); ok && ec.Count() == 1 {
			err = ec[0]
		} else {
			break
		}
	}

	return ErrorReader{err}
}

func (e ErrorReader) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func (e ErrorReader) Count() int {
	if e.Err != nil {
		return e.Err.Count()
	}
	return 0
}
func (e ErrorReader) Empty() bool {
	if e.Err != nil {
		return e.Err.Empty()
	}
	return true
}
func (e ErrorReader) Name() string {
	if e.Err != nil {
		return e.Err.Name()
	}
	return ""
}
