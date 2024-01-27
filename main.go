package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xvrlad/budgeting-app/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}