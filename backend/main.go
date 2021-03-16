package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pongsakorn-maker/react-phone-store/backend/database"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	defer database.DBConnection.Close()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/phone", Phone.GetPhones)
	app.Get("/api/v1/phone/:id", Phone.GetPhone)
	app.Post("/api/v1/phone", Phone.CreatePhone)
	app.Delete("/api/v1/phone/:id", Phone.DeletePhone)
}

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConnection.AutoMigrate(&Phone.Phone{})
	fmt.Println("Database Migrated")
}
