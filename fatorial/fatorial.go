package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 150; i++ {
		resultado := fatorial(i)
		fmt.Println("fatorial", i, resultado)
	}
}

func fatorial(valor int) int {

	if valor == 0 {
		return 1
	} else if valor == 1 {
		return 1
	} else {
		return valor * fatorial(valor-1)
	}
}
