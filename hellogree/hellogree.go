package main

import (
	"fmt"
	"log"

	"github.com/falessio/greetings"
)

func main() {
	log.SetPrefix("greetings ")
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
