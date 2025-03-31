package controllers

import (
	"gin/initializers"
	"gin/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func StaffPost(c *gin.Context) {

	//Create Staff
	var staff models.Staff

	if err := c.ShouldBindJSON(&staff); err != nil {
		// Return an error if binding fails
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	// Hash the password before saving it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}
	staff.Password = string(hashedPassword)
	staff.CreatedAt = time.Now()

	result := initializers.DB.Create(&staff)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Staff created successfully",
		"data":    staff,
	})
}
