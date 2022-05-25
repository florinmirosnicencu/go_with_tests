package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Flo", "")
		want := "Hello, Flo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Flo", "Spanish")
		want := "Hola, Flo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Flo", "French")
		want := "Bonjour, Flo"

		assertCorrectMessage(t, got, want)
	})

}
