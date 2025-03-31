package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start project")

	r := gin.Default()

	// ตั้งค่า routes ผ่าน controller
	// controller.SetupRoutes(r)

	// run at port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
