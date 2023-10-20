package main

import (
	"fmt"

	"github.com/falessio/go-rest-api/database"
	"github.com/falessio/go-rest-api/models"
	"github.com/falessio/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor rest")
	routes.HandleRequest()
}
