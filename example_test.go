package check

type Person struct {
	Name string
}

func Example() {
	p := &Person{
		Name: "invalid*",
	}

	s := Struct{
		"Name": String{
			NonEmpty{},
			Regexp(`^[a-zA-Z0-9]+$`),
			MinChar{10},
		},
	}

	if er := Validate(s, p); er.IsError() {
		if er.Get("Name").IsError() {
			panic("key 'Name' does not exists")
		}
	}
}
