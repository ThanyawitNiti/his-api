package controllers

import (
	"gin/initializers"
	"gin/models"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

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

func StaffLogin(c *gin.Context) {
	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
		// HospitalId int8  `json:"hospital_id"`
	}

	// Bind the incoming JSON request to the 'loginDetails' struct
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Retrieve the staff record by username from the database
	var staff models.Staff
	result := initializers.DB.Where("username = ?", loginDetails.Username).First(&staff)
	if result.Error != nil {
		// Staff not found
		c.JSON(404, gin.H{
			"error": "Staff not found",
		})
		return
	}
	// Compare the provided password with the hashed password stored in the database
	err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(loginDetails.Password))
	if err != nil {
		// Password does not match
		c.JSON(401, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Create a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims for the token (user info and expiration time)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = staff.Username
	claims["hospital_id"] = staff.HospitalId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tk, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal("Error signing token", err)
		c.JSON(500, gin.H{
			"error": "Error generating JWT token",
		})
		return
	}

	// If login is successful, return a success message (or a JWT token if you use one)
	c.JSON(200, gin.H{
		"message": "Login successful",
		"data":    staff, // You can return user data or a token
		"token":   tk,    // Send JWT token to client
	})
}
