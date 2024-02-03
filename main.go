package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xvrlad/budgeting-app/controllers"
	"github.com/xvrlad/budgeting-app/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	itemsRoute := r.Group("/items") 
	{
		itemsRoute.POST("/", controllers.CreateItem)
		itemsRoute.GET("/", controllers.GetItems)
		itemsRoute.GET("/:id", controllers.GetItemById)
	}
	
	r.Run()
}