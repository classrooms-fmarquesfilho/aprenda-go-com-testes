package main

import "testing"

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithConfig(t *testing.T) {
	t.Run("config completa", func(T *testing.T) {
		cfg := HelloConfig{
			Name:   "Maria",
			Hour:   9,
			Region: RegionRN,
		}
		got := HelloWithConfig(cfg)
		want := "Bom dia, boy"
		assertMessage(t, got, want)
	})

	t.Run("config parcial - só nome e hora", func(t *testing.T) {
		cfg := HelloConfig{
			Name: "Maria",
			Hour: 9,
			// região omitida
		}
		got := HelloWithConfig(cfg)
		want := "Bom dia, Maria"
		assertMessage(t, got, want)
	})
}

func TestHelloWithRegion(t *testing.T) {
	t.Run("Rio Grande do Norte morning", func(t *testing.T) {
		got := HelloWithRegion(9, RegionRN)
		want := "Bom dia, boy"
		assertMessage(t, got, want)
	})

	t.Run("São Paulo afternoon", func(t *testing.T) {
		got := HelloWithRegion(14, RegionSP)
		want := "Boa tarde, mano"
		assertMessage(t, got, want)
	})

	t.Run("Minas Gerais morning", func(t *testing.T) {
		got := HelloWithRegion(10, RegionMG)
		want := "Bom dia, sô"
		assertMessage(t, got, want)
	})

	t.Run("Rio Grande do Sul evening", func(t *testing.T) {
		got := HelloWithRegion(22, RegionRS)
		want := "Boa noite, tchê"
		assertMessage(t, got, want)
	})
}

func TestHello(t *testing.T) {
	t.Run("apenas nome", func(t *testing.T) {
		got := Hello(WithName("Fernando"))
		want := "Olá, Fernando"
		assertMessage(t, got, want)
	})

	t.Run("nome e horário", func(t *testing.T) {
		got := Hello(
			WithName("Maria"),
			WithHour(9),
		)
		want := "Bom dia, Maria"
		assertMessage(t, got, want)
	})

	t.Run("completo - nome, hora e região", func(t *testing.T) {
		got := Hello(
			WithName("João"),
			WithHour(20),
			WithRegion(RegionRN),
		)
		want := "Boa noite, João"
		assertMessage(t, got, want)
	})

	t.Run("apenas região - usa vocativo", func(t *testing.T) {
		got := Hello(WithRegion(RegionRN))
		want := "Olá, boy"
		assertMessage(t, got, want)
	})

	t.Run("hora e região sem nome - usa vocativo", func(t *testing.T) {
		got := Hello(
			WithHour(14),
			WithRegion(RegionMG),
		)
		want := "Boa tarde, sô"
		assertMessage(t, got, want)
	})

	t.Run("nome com espaços em branco", func(t *testing.T) {
		got := Hello(WithName("  Fernando  "))
		want := "Olá, Fernando"
		assertMessage(t, got, want)
	})

	t.Run("nome vazio (só espaços) com região - usa vocativo", func(t *testing.T) {
		got := Hello(
			WithName("   "),
			WithRegion(RegionSP),
		)
		want := "Olá, mano"
		assertMessage(t, got, want)
	})

	t.Run("nome vazio explícito com região - usa vocativo", func(t *testing.T) {
		got := Hello(
			WithName(""),
			WithRegion(RegionRS),
		)
		want := "Olá, tchê"
		assertMessage(t, got, want)
	})

	t.Run("sem options - usa padrões", func(t *testing.T) {
		got := Hello()
		want := "Olá, Mundo"
		assertMessage(t, got, want)
	})

	t.Run("nome vazio sem região - usa Mundo", func(t *testing.T) {
		got := Hello(WithName(""))
		want := "Olá, Mundo"
		assertMessage(t, got, want)
	})

	t.Run("região inválida sem nome - usa Mundo", func(t *testing.T) {
		got := Hello(WithRegion("xyz"))
		want := "Olá, Mundo"
		assertMessage(t, got, want)
	})
}
