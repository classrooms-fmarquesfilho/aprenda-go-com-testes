package main

import "testing"

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	// --- casos normais ---
	t.Run("sem opções", func(t *testing.T) {
		assertMessage(t, Hello(), "Olá, Mundo")
	})

	t.Run("só nome", func(t *testing.T) {
		assertMessage(t, Hello(WithName("Fernando")), "Olá, Fernando")
	})

	t.Run("nome e hora", func(t *testing.T) {
		assertMessage(t,
			Hello(WithName("Maria"), WithHour(9)),
			"Bom dia, Maria")
	})

	t.Run("nome, hora e região — nome tem prioridade", func(t *testing.T) {
		assertMessage(t,
			Hello(WithName("João"), WithHour(20), WithRegion(RegionRN)),
			"Boa noite, João")
	})

	// --- edge cases: limites do intervalo de hora ---
	t.Run("hora 0 - meia-noite", func(t *testing.T) {
		assertMessage(t, Hello(WithHour(0)), "Boa noite, Mundo")
	})

	t.Run("hora 23 - limite superior", func(t *testing.T) {
		assertMessage(t, Hello(WithHour(23)), "Boa noite, Mundo")
	})

	// --- valores inválidos ---
	t.Run("hora 24 - inválida", func(t *testing.T) {
		assertMessage(t, Hello(WithHour(24)), "Olá, Mundo")
	})

	t.Run("hora negativa - inválida", func(t *testing.T) {
		assertMessage(t, Hello(WithHour(-5)), "Olá, Mundo")
	})

	// --- região ---
	t.Run("só região", func(t *testing.T) {
		assertMessage(t, Hello(WithRegion(RegionRN)), "Olá, boy")
	})

	t.Run("hora e região sem nome", func(t *testing.T) {
		assertMessage(t,
			Hello(WithHour(14), WithRegion(RegionMG)),
			"Boa tarde, sô")
	})

	t.Run("região inválida", func(t *testing.T) {
		assertMessage(t, Hello(WithRegion("xyz")), "Olá, Mundo")
	})

	// --- nome com espaços ---
	t.Run("nome com espaços - TrimSpace", func(t *testing.T) {
		assertMessage(t,
			Hello(WithName("  Fernando  ")),
			"Olá, Fernando")
	})

	t.Run("nome só espaços com região", func(t *testing.T) {
		assertMessage(t,
			Hello(WithName("   "), WithRegion(RegionSP)),
			"Olá, mano")
	})
}
