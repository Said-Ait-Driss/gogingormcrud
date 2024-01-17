package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/", controllers.HelloWorld)
	r.GET("/products/:id", controllers.GetProduct)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	return r
}
