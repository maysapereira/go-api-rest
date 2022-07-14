package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil)) //log Fatal faz com que qualquer problema que aconteça ao subir o servidor seja avisado

}

func main() {
	fmt.Println("Iniciando projeto: API REST em GO")
	HandleRequest()
}
