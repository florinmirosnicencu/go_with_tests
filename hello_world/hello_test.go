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
		got := Hello("Flo")
		want := "Hello, Flo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World1"

		assertCorrectMessage(t, got, want)
	})

}
