package main

import "testing"

func Test_Greet(t *testing.T) {
	person := Person{greeting: "Hello!"}
	result := person.Greet()

	if result != "Hello!" {
		t.Fatalf("person.Greet(): %#v", result)
	}
}
