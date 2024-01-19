package main

import (
	"github.com/falessio/gin-api-rest/database"
	"github.com/falessio/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
