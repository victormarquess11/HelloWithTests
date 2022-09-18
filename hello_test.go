package main

import "testing"

const englishLanguage = "English"
const spanishLanguage = "Spanish"

func TestHello(t *testing.T) {
	t.Run("english saying hello to people", func(t *testing.T) {
		got := Hello("Chris", englishLanguage)
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("english say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", englishLanguage)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish saying hello to people", func(t *testing.T) {
		got := Hello("Chris", spanishLanguage)
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", spanishLanguage)
		want := "Hola, Mundo"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
