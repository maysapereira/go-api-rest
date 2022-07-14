package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maysapereira/go-api-rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personagens", controllers.TodosPersonagens)
	log.Fatal(http.ListenAndServe(":8000", r))
}