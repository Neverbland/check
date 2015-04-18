package check

import (
	"fmt"
	"github.com/neverbland/access"
	"github.com/neverbland/fail"
)

func ValidationErr(name, message string, params ...interface{}) ValidationError {
	return ValidationError{name, message, params}
}

// ValidationError implements Error
type ValidationError struct {
	Name    string
	Message string
	Params  []interface{}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf(e.Message, e.Params...)
}

func (e ValidationError) ErrorName() string { return e.Name }

type ErrorInterface error

//Error reader
type Reader struct {
	ErrorInterface
}

func (e Reader) Error() string {
	if e.IsError() {
		return e.ErrorInterface.Error()
	}
	return ""
}

func (e Reader) Get(path ...string) Reader {

	if len(path) == 0 {
		return e
	}

	if fail.IsError(e.ErrorInterface) {
		if err, ok := access.MustRead(path[0], e.ErrorInterface).(error); ok && fail.IsError(err) {
			return Reader{err}
		}
	}
	return Reader{}
}

func (e Reader) IsError() bool {
	return fail.IsError(e.ErrorInterface)
}

func (e Reader) OrNil() error {
	return fail.OrNil(e.ErrorInterface)
}
