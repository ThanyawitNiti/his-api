package main

import (
	"fmt"
	"gin/controllers"
	"gin/initializers"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Start project")
	r := gin.Default()

	r.POST("/staff/create", controllers.StaffPost)
	r.POST("/staff/login", controllers.StaffLogin)
	r.GET("/", controllers.PatientGet)

	// run at port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
