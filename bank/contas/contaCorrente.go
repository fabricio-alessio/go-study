package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque < 0 {
		return "O valor do saque precisa ser positivo"
	}
	podeSacar := c.Saldo > valorDoSaque
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "O valor do deposito precisa ser positivo", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {

	if valor > 0 && c.Saldo > valor {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}
