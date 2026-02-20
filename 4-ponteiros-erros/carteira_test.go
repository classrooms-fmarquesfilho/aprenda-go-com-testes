package ponteiroserros

import "testing"

func TestCarteira(t *testing.T) {
	carteira := Carteira{}
	carteira.Depositar(10)

	got := carteira.Saldo()
	want := 10

	if got != want {
		t.Errorf("obtive %d esperava %d", got, want)
	}
}
