package main

import (
	"github.com/xvrlad/budgeting-app/initializers"
	"github.com/xvrlad/budgeting-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Item{})
}