package contas

import "alessio.com/study/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Tipo          int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	if valorDoSaque < 0 {
		return "O valor do saque precisa ser positivo"
	}
	podeSacar := c.saldo > valorDoSaque
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do deposito precisa ser positivo", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
