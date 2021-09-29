package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}

//create a product
func CreateProduct(db *gorm.DB, product *Product) error {
	result := db.Create(product)
	return result.Error
 }
 
 //get products
 func GetProducts(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
	   return err
	}
	return nil
 }
 
 //get product by id
 func GetProduct(db *gorm.DB, product *Product, id string) (err error) {
	err = db.Where("id = ?", id).First(product).Error
	if err != nil {
	   return err
	}
	return nil
 }
 
 //update product
 func UpdateProduct(db *gorm.DB, product *Product) (err error) {
	db.Save(product)
	return nil
 }
 
 //delete product
 func DeleteProduct(db *gorm.DB, product *Product, id string) (err error) {
	db.Where("id = ?", id).Delete(product)
	return nil
 }