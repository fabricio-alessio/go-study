package main

import (
	"fmt"

	"alessio.com/study/contas"
)

func main() {

	contaDoGuilherme := contas.ContaCorrente{
		Titular:       "Guilherme",
		NumeroAgencia: 2344,
		NumeroConta:   83929283,
		Saldo:         345,
	}
	contaDoGuilherme.NumeroAgencia = 33

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDoGuilherme.Sacar(234))
	fmt.Println(contaDoGuilherme)

	contaDaCris := new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"

	fmt.Println(contaDaCris)

	fmt.Println(contaDaCris.Sacar(200))

	fmt.Println(contaDaCris.Sacar(-200))
	status, valor := contaDaCris.Depositar(300)
	fmt.Println("Status deposito:", status, ", novo valor:", valor)
	fmt.Println(contaDaCris)

	fmt.Println(contaDoGuilherme, contaDaCris)
	contaDoGuilherme.Transferir(40, contaDaCris)
	fmt.Println(contaDoGuilherme, contaDaCris)
}
