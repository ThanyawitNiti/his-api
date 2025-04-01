package controllers

import (
	"fmt"
	"gin/initializers"
	"gin/models"
	"log"
	"net/http"
	"os"
	"strings"
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
		"data": map[string]interface{}{
			"username":    staff.Username,
			"hospital_id": staff.HospitalId,
		},
	})
}

func StaffLogin(c *gin.Context) {
	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

	// If login is successful, return a success message
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   tk, // Send JWT token to client for /patient/search and /patient/:id
	})
}

func checkToken(c *gin.Context) (jwt.MapClaims, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, fmt.Errorf("authorization token is required")
	}

	// ลบ "Bearer " ที่นำหน้า token ออก
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// ตรวจสอบ token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// ตรวจสอบ signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		// ใช้ secret key จาก environment variable
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	// ตรวจสอบว่า token เป็น valid JWT
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func PatientSearch(c *gin.Context) {
	// ตรวจสอบ token
	claims, err := checkToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	hospitalID := claims["hospital_id"].(float64)

	// ค่าจาก query parameters
	firstName := c.DefaultQuery("first_name", "")
	middleName := c.DefaultQuery("middle_name", "")
	lastName := c.DefaultQuery("last_name", "")
	dateOfBirth := c.DefaultQuery("date_of_birth", "")
	nationalID := c.DefaultQuery("national_id", "")
	passportID := c.DefaultQuery("passport_id", "")
	phoneNumber := c.DefaultQuery("phone_number", "")
	email := c.DefaultQuery("email", "")

	// ค้นหาผู้ป่วยจากข้อมูลที่ได้รับ
	var patients []models.Patient
	query := initializers.DB.Model(&models.Patient{})

	// ถ้ามีการกรอกค่าต่างๆ มา, เราจะทำการ filter ค้นหาจากฟิลด์นั้นๆ
	if firstName != "" {
		query = query.Where("first_name_th ILIKE ?", "%"+firstName+"%")
	}
	if middleName != "" {
		query = query.Where("middle_name_th ILIKE ?", "%"+middleName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name_th ILIKE ?", "%"+lastName+"%")
	}
	if dateOfBirth != "" {
		query = query.Where("date_of_birth ILIKE ?", dateOfBirth)
	}
	if nationalID != "" {
		query = query.Where("national_id ILIKE ?", nationalID)
	}
	if passportID != "" {
		query = query.Where("passport_id ILIKE ?", passportID)
	}
	if phoneNumber != "" {
		query = query.Where("phone_number ILIKE ?", phoneNumber)
	}
	if email != "" {
		query = query.Where("email ILIKE ?", email)
	}

	// ตรวจสอบว่า hospital_id ใน JWT token ตรงกับ hospital_id ของผู้ป่วยหรือไม่
	query = query.Where("hospital_id = ?", hospitalID)

	// ดึงข้อมูลผู้ป่วยที่ตรงกับเงื่อนไข
	if err := query.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search for patients"})
		return
	}

	// หากไม่พบผู้ป่วยใดๆ
	if len(patients) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "ไม่พบข้อมูลคนไข้"})
		return
	}

	// ส่งผลลัพธ์กลับไป
	c.JSON(http.StatusOK, gin.H{"patients": patients})
}

// ฟังก์ชันสำหรับค้นหาผู้ป่วยโดยใช้ id
func PatientGetByid(c *gin.Context) {
	// รับ id จาก URL parameter
	id := c.Param("id")

	// ค้นหาผู้ป่วยโดยใช้ national_id หรือ passport_id
	var patient models.Patient

	// ค้นหาผู้ป่วยที่มีค่า ตรงกับ id ที่ส่งมา
	err := initializers.DB.Where("id = ? OR national_id = ? OR passport_id = ?", id, id, id).First(&patient).Error
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "ไม่พบผู้ป่วยที่ตรงกับข้อมูลที่ให้"})
		return
	}

	// ส่งข้อมูลผู้ป่วยกลับในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{
		"first_name_th":  patient.FirstNameTH,
		"middle_name_th": patient.MiddleNameTH,
		"last_name_th":   patient.LastNameTH,
		"first_name_en":  patient.FirstNameEN,
		"middle_name_en": patient.MiddleNameEN,
		"last_name_en":   patient.LastNameEN,
		"date_of_birth":  patient.DateOfBirth,
		"patient_hn":     patient.PatientHN,
		"national_id":    patient.NationalID,
		"passport_id":    patient.PassportID,
		"phone_number":   patient.PhoneNumber,
		"email":          patient.Email,
		"gender":         patient.Gender,
	})
}
