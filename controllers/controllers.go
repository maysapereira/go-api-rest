package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maysapereira/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodosPersonagens(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personagens)
}

func RetornaUmPersonagem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personagem := range models.Personagens {
		if strconv.Itoa(personagem.Id) == id {
			json.NewEncoder(w).Encode(personagem)
		}
	}

	//
}