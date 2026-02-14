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
		got := Hello("Maria")
		want := "Olá, Maria"
		assertMessage(t, got, want)
	})

	t.Run("usar 'Mundo' quando string for vazia", func(t *testing.T) {
		got := Hello("")
		want := "Olá, Mundo"
		assertMessage(t, got, want)
	})

	t.Run("saudação de manhã", func(t *testing.T) {
		hour := 9
		got := HelloWithTime("Rodolfo", hour)
		want := "Bom dia, Rodolfo"
		assertMessage(t, got, want)
	})

	t.Run("saudação da tarde", func(t *testing.T) {
		hour := 14
		got := HelloWithTime("Maria", hour)
		want := "Boa tarde, Maria"
		assertMessage(t, got, want)
	})

	t.Run("saudação da noite", func(t *testing.T) {
		hour := 20
		got := HelloWithTime("João", hour)
		want := "Boa noite, João"
		assertMessage(t, got, want)
	})
}
