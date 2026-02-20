package ponteiroserros

import "fmt"

type Carteira struct {
	saldo int
}

type ErroSaldoInsuficiente struct {
	Solicitado int
	Disponivel int
}

func (e ErroSaldoInsuficiente) Error() string {
	return fmt.Sprintf("não é possível sacar %d, pois o saldo disponível é %d",
		e.Solicitado, e.Disponivel)
}

func (c *Carteira) Depositar(quantia int) {
	c.saldo += quantia
}

func (c *Carteira) Saldo() int {
	return c.saldo
}

func (c *Carteira) Sacar(quantia int) error {
	if quantia > c.saldo {
		return ErroSaldoInsuficiente{
			Solicitado: quantia,
			Disponivel: c.saldo,
		}
	}
	c.saldo -= quantia
	return nil
}
