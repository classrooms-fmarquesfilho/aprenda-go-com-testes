package ponteiroserros

import "testing"

func TestCarteira(t *testing.T) {
	t.Run("depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(10)
		verificarSaldo(t, carteira, 10)
	})

	t.Run("sacar com saldo", func(t *testing.T) {
		carteira := Carteira{saldo: 20}
		err := carteira.Sacar(10)

		if err != nil {
			t.Errorf("não esperava erro mas obteve %v", err)
		}
		verificarSaldo(t, carteira, 10)
	})

	t.Run("sacar com saldo insuficiente", func(t *testing.T) {
		carteira := Carteira{saldo: 5}
		err := carteira.Sacar(10)

		if err == nil {
			t.Error("esperava erro mas não obteve nenhum")
		}
	})
}

func verificarSaldo(t testing.TB, carteira Carteira, esperado int) {
	t.Helper()
	obtido := carteira.Saldo()
	if obtido != esperado {
		t.Errorf("obtido %d esperado %d", obtido, esperado)
	}
}
