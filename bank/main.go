package main

import (
	"fmt"

	"alessio.com/study/clientes"
	"alessio.com/study/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoGuilherme := contas.ContaCorrente{
		Titular:       clientes.Titular{Nome: "Guilherme"},
		NumeroAgencia: 2344,
		NumeroConta:   83929283,
	}
	contaDoGuilherme.NumeroAgencia = 33
	contaDoGuilherme.Depositar(400)

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDoGuilherme.Sacar(234))
	fmt.Println(contaDoGuilherme)

	contaDaCris := new(contas.ContaCorrente)
	contaDaCris.Titular = clientes.Titular{Nome: "Cris"}

	fmt.Println(contaDaCris)

	fmt.Println(contaDaCris.Sacar(200))

	fmt.Println(contaDaCris.Sacar(-200))
	status, valor := contaDaCris.Depositar(300)
	fmt.Println("Status deposito:", status, ", novo valor:", valor)
	fmt.Println(contaDaCris)

	fmt.Println(contaDoGuilherme, contaDaCris)
	contaDoGuilherme.Transferir(40, contaDaCris)
	fmt.Println(contaDoGuilherme, contaDaCris)

	fmt.Println(contaDoGuilherme.ObterSaldo())

	PagarBoleto(&contaDoGuilherme, 10)
	PagarBoleto(contaDaCris, 10)
	fmt.Println(contaDoGuilherme, contaDaCris)
}
