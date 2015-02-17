package main

type Greeter interface {
	Greet() string
}

type Person struct {
	greeting string
}

func NewGreeter(greeting string) Greeter {
	return &Person{greeting: greeting}
}

func (person *Person) Greet() string {
	return person.greeting
}
