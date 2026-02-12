package main

import "testing"

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("cumprimentar pessoas", func(t *testing.T) {
		got := Hello("Web II")
		want := "Olá, Web II"
		assertMessage(t, got, want)
	})

	t.Run("usar 'Mundo' quando string for vazia", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"
		assertMessage(t, got, want)
	})
}
