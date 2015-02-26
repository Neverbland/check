package check

import "fmt"

func ExampleValidator() {
	username := "invalid*"
	validator := String{CustomStringContainValidator{"admin"}}
	e := validator.Validate(username)
	fmt.Println(e.Error())
}
