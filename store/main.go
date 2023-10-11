package main

import (
	"log"
	"net/http"

	"alessio.com/study/store/routes"
)

func main() {
	routes.CarregaRotas()
	log.Println("Starting")
	http.ListenAndServe(":8000", nil)
}
