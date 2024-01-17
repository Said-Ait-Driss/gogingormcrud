package controllers

import (
	"go-crud/db"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	var product models.Product

	if err := db.DB.First(&product, ctx.Param("id")).Error; err != nil {
		ctx.JSON(404, gin.H{"Error": "Product not found !"})
		return
	}

	ctx.JSON(200, product)
}

func CreateProduct(ctx *gin.Context) {
	var productInput models.Product

	if err := ctx.ShouldBindJSON(&productInput); err != nil {
		ctx.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	newProduct := models.Product{
		Title:       productInput.Title,
		Description: productInput.Description,
		Price:       productInput.Price,
	}

	if err := db.DB.Create(&newProduct).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to create product !",
		})
		return
	}
	ctx.JSON(201, newProduct)
}

func UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	var existingProduct models.Product
	var productInput models.Product

	if err := db.DB.First(&existingProduct, productId).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&productInput); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existingProduct.Title = productInput.Title
	existingProduct.Description = productInput.Description
	existingProduct.Price = productInput.Price

	if err := db.DB.Save(&existingProduct).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(200, existingProduct)

}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("id")

	var existingProduct models.Product
	if err := db.DB.First(&existingProduct, productId).Error; err != nil {
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	if err := db.DB.Delete(&existingProduct).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete Product"})
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
