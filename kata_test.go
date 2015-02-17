package main

import "testing"

func Test_Greet(t *testing.T) {
	person := Person{greeting: "Hello!"}
	assertEqual(t, "Hello!", person.Greet())
}

func assert(t *testing.T, predicate bool) {
	assertEqual(t, true, predicate)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("assertion failed: expected %#v, but got %#v", expected, actual)
	}
}
