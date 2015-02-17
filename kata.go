package main

type Person struct {
	greeting string
}

func (person *Person) Greet() string {
	return person.greeting
}
