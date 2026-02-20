package ponteiroserros

type Carteira struct {
	saldo int
}

func (c *Carteira) Depositar(quantia int) {
	c.saldo += quantia
}

func (c Carteira) Saldo() int {
	return c.saldo
}
