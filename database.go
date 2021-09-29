package backendrepository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(username string, password string, address string, port string, dbName string) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbName)
	log.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	return
}