package main

import "testing"

func TestHelloWorld(t *testing.T) {
	assertIsCorrectMessage := func(got string, expected string, t *testing.T) {
		t.Helper()
		if got != expected {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	}
	t.Run("when say hello with a name then response hello + name", func(t *testing.T) {
		expected := "Hello, Alex"
		got := HelloWorld("Alex")

		assertIsCorrectMessage(got, expected, t)
	})

	t.Run("when say hello with empty name then response is say hello world", func(t *testing.T) {
		expected := "Hello, world"
		got := HelloWorld("")

		assertIsCorrectMessage(got, expected, t)
	})
}
