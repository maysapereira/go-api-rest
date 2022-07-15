package main

import (
	"fmt"

	"github.com/maysapereira/go-api-rest/database"
	"github.com/maysapereira/go-api-rest/models"
	"github.com/maysapereira/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando projeto: API REST em GO")

	models.Personagens = []models.Personagen{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaBancoDeDados()
	routes.HandleRequest()
}
