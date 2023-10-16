package main

import (
	"log"
	"net/http"

	"alessio.com/study/store/routes"
)

func main() {
	routes.CarregaRotas()
	log.Println("Starting on port 8000")
	http.ListenAndServe(":8000", nil)
}
