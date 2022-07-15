package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maysapereira/go-api-rest/controllers"
	"github.com/maysapereira/go-api-rest/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personagens", controllers.TodosPersonagens).Methods("Get")
	r.HandleFunc("/api/personagens/{id}", controllers.RetornaUmPersonagem).Methods("Get")
	r.HandleFunc("/api/personagens", controllers.CriaPersonagem).Methods("Post")
	r.HandleFunc("/api/personagens/{id}", controllers.DeletaPersonagem).Methods("Delete")
	r.HandleFunc("/api/personagens/{id}", controllers.EditaPersonagem).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
