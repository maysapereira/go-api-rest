package models

type Personagem struct {
	Id int `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personagens []Personagem

