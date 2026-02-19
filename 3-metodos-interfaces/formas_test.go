package metodosinterfaces

import "testing"

func TestArea(t *testing.T) {
	r := Retangulo{Largura: 12, Altura: 6}
	c := Circulo{Raio: 10}

	t.Run("retângulo", func(t *testing.T) {
		got := r.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("círculo", func(t *testing.T) {
		got := c.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
