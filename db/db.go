package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "go-crud/config"
)

var DB *gorm.DB

func InitDB() {
    dsn := config.GetDSN()
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }
}