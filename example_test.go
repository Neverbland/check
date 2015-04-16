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

	er := ErrorReader{s.Validate(*p)}
	if er.Count() != 0 {
		if er.Get("Name").Empty() {
			panic("key 'Name' does not exists")
		}
	}
}
