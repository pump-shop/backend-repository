package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gormv1 "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

var DBv1 *gormv1.DB

func ConnectToDB(username string, password string, address string, port string, dbName string) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbName)
	log.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	return
}

func ConnectToDBv1(username string, password string, address string, port string, dbName string) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbName)
	log.Println(dsn)

	DBv1, err = gormv1.Open("mysql", dsn)
	
	return
}