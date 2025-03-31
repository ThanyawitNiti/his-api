package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env")
	}
}

var (
	db *gorm.DB
)

// Connect เชื่อมต่อฐานข้อมูล PostgreSQL
func Connect() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(filepath.Join(pwd, "../../pkg/config/.env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Printf("godotenv : %s = %s \n", "dbUser", dbUser)
	fmt.Printf("godotenv : %s = %s \n", "DB Host", dbHost)

	// เปลี่ยนจาก MySQL เป็น PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB คืนค่าการเชื่อมต่อฐานข้อมูล
func GetDB() *gorm.DB {
	return db
}
