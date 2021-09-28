package models

import "log"

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func InModel() {
	log.Println("in model")
}