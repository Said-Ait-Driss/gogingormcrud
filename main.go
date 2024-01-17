package main


import (
	// "fmt"
	"log"
	"os"

	// "github.com/gin-gonic/gin"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"

	"go-crud/routes"
	"go-crud/db"
)

func init(){
	db.InitDB()
}

func main(){
	router := routes.SetupRouter()

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "9090"		
	}

	log.Println("Server is starting...")

	// Run the server
	if err := router.Run(":"+port); err != nil {
		log.Fatal("Failed to start the server:", err)
	}

	log.Println("Server started successfully")
}