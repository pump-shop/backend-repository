package main

import (
	"fmt"

	"github.com/pump-shop/backend-repository/database"
	"github.com/pump-shop/backend-repository/models"
)

func main() {
	err := database.ConnectToDB("root", "19121378", "localhost", "3306", "test")
	fmt.Println(err)
	p := models.Product{
		Name: "hasan",
		Price: 20,
		}
	fmt.Println(p)
}