package models

type Personagen struct {
	Id int `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personagens []Personagen

/*
tive que mudar o nome do modelo de "Personagem" para "Personagen" 
porque o postgres coloca o modelo no plural automaticamente 
adicionando um "s" e isso fez dar erro.

O erro que estava dando:
2022/07/15 09:52:29 /home/maysa/go/src/estudos/first-api-rest-go/controllers/controllers.go:20 ERROR: relation "personagems" does not exist (SQLSTATE 42P01)
[0.398ms] [rows:0] SELECT * FROM "personagems"
*/