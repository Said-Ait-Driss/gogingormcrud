package models

type Product struct {
	Id uint "gorm:primaryKey"
	Title string
	Price float64
	Description string
}