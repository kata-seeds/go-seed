package main

type Greetor interface {
	Greet() string
}

type Person struct {
	greeting string
}

func NewGreeter(greeting string) Greetor {
	return &Person{greeting: greeting}
}

func (person *Person) Greet() string {
	return person.greeting
}
