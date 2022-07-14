package main

import (
	"fmt"

	"github.com/maysapereira/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando projeto: API REST em GO")
	routes.HandleRequest()
}