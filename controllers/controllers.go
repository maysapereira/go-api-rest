package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maysapereira/go-api-rest/database"
	"github.com/maysapereira/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodosPersonagens(w http.ResponseWriter, r *http.Request) {
	var p []models.Personagen
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmPersonagem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personagen

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriaPersonagem(w http.ResponseWriter, r *http.Request) {
	var novoPersonagem models.Personagen
	json.NewDecoder(r.Body).Decode(&novoPersonagem)
	database.DB.Create(&novoPersonagem)
	json.NewEncoder(w).Encode(novoPersonagem)
}

func DeletaPersonagem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personagen
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}