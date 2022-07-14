package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maysapereira/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodosPersonagens(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personagens)
}