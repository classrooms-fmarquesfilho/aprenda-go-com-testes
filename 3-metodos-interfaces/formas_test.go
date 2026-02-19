package metodosinterfaces

import "testing"

func checarArea(t testing.TB, forma Forma, want float64) {
	t.Helper()
	got := forma.Area()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("retângulo", func(t *testing.T) {
		r := Retangulo{Largura: 12, Altura: 6}
		checarArea(t, r, 72.0)
	})

	t.Run("círculo", func(t *testing.T) {
		c := Circulo{Raio: 10}
		checarArea(t, c, 314.1592653589793)
	})
}
