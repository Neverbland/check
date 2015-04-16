package check

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type CustomStringContainValidator struct {
	Constraint string
}

func (validator CustomStringContainValidator) Validate(v interface{}) Error {
	if !strings.Contains(v.(string), validator.Constraint) {
		return ValidationErr("customStringContainValidator", "customStringContainValidator", v, validator.Constraint)
	}

	return nil
}

type User struct {
	Username string
	Password string
	Name     string
	Age      int
	Email    string
	Birthday time.Time
}

func (u *User) Validate() Error {
	s := Struct{
		"Username": String{
			NonEmpty{},
			Regexp(`^[a-zA-Z0-9]+$`),
		},
		"Password": String{
			NonEmpty{},
			MinChar{8},
		},
		"Name": NonEmptyString{},
		"Age": Number{
			GreaterThan{3},
			LowerThan{120},
		},
		"Email": String{
			Email{},
			CustomStringContainValidator{"test.com"},
		},
		"Birthday": Time{
			Before{time.Date(1990, time.January, 1, 1, 0, 0, 0, time.UTC)},
			After{time.Date(1900, time.January, 1, 1, 0, 0, 0, time.UTC)},
		},
	}
	return s.Validate(u)
}

func TestIntegration(t *testing.T) {

	assert := assert.New(t)

	invalidUser := &User{
		"not-valid-username*",
		"123", // Invalid password length
		"", // Cannot be empty
		150, // Invalid age
		"@test", // Invalid email address
		time.Date(1991, time.January, 1, 1, 0, 0, 0, time.UTC), // Invalid date
	}

	validUser := &User{
		"testuser",
		"validPassword123",
		"Good Name",
		20,
		"test@test.com",
		time.Date(1980, time.January, 1, 1, 0, 0, 0, time.UTC),
	}

	err := ErrorReader{invalidUser.Validate()}

	assert.False(err.Empty(), "Expected 'invalidUser' to be invalid")

	err = err.Get("Username")
	assert.False(err.Empty(), "Expected errors for 'Username'")
	assert.Equal(1, err.Count(), "Expected 1 error for 'Username'")
	assert.Equal("string.regex", err.Name(), "Expected regex error")

	assert.True(ErrorReader{validUser.Validate()}.Empty(), "Expected 'validUser' to be valid")
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		invalidUser := &User{
			"not-valid-username*",
			"123", // Invalid password length
			"", // Cannot be empty
			150, // Invalid age
			"@test", // Invalid email address
			time.Date(1991, time.January, 1, 1, 0, 0, 0, time.UTC), // Invalid date
		}

		invalidUser.Validate()
	}
}
