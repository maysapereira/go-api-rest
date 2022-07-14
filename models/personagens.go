package models

type Personagem struct {
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personagens []Personagem

