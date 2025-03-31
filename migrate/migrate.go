package main

import (
	"gin/initializers"
	"gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Patient{})
	initializers.DB.AutoMigrate(&models.Hospital{})
	initializers.DB.AutoMigrate(&models.Staff{})
}
