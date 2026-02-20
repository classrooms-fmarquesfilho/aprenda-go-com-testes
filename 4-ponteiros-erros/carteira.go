package ponteiroserros

import "fmt"

type error interface {
	Error() string
}

type Carteira struct {
	saldo int
}

func (c *Carteira) Depositar(quantia int) {
	c.saldo += quantia
}

func (c Carteira) Saldo() int {
	return c.saldo
}

func (c *Carteira) Sacar(quantia int) error {
	if quantia > c.saldo {
		return fmt.Errorf("não é possível sacar %d, pois o saldo é %d",
			quantia, c.saldo)
	}
	c.saldo -= quantia
	return nil
}
