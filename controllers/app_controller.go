package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Sucess": "Hello world",
	})
}
