package database

import (
	"fmt"
	"info-bd-api/model"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	errf := godotenv.Load()
	if errf != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to load env: "+errf.Error())
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to connect DB: "+err.Error())
	}
	return nil
}

func MigrateSchema() {
	DB.AutoMigrate(&model.Division{}, &model.District{})
}
