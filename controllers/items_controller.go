package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xvrlad/budgeting-app/initializers"
	"github.com/xvrlad/budgeting-app/models"
)

// todo: add request validation
func CreateItem(c *gin.Context) {
	// Get data off request body
	var request struct {
		Name string
		Price float32
		WantDate time.Time
	}

	c.Bind(&request)

	// Create item
	item := models.Item{Name: request.Name, Price: request.Price, WantDate: request.WantDate}

	result := initializers.DB.Create(&item) 

	if (result.Error != nil) {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	// Return item
	c.JSON(200, gin.H{
		"item": item,
	})
}

func GetItems(c *gin.Context) {
	// Get items
	var items []models.Item
	result := initializers.DB.Find(&items)

	if (result.Error != nil) {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	// Return items
	c.JSON(200, gin.H{
		"items": items,
	})
}

func GetItemById(c *gin.Context) {
	id := c.Param("id")

	var item models.Item
	result := initializers.DB.First(&item, id)

	if (result.Error != nil) {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}