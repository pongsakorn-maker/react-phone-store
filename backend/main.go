package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pongsakorn-maker/react-phone-store/database"
	phone "github.com/pongsakorn-maker/react-phone-store/schema"
)

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConnection.Close()
	setupRoutes(app)
	app.Listen(3000)
}

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "phone.db")
	if err != nil {
		panic("Falied to connect to database")
	}
	fmt.Println("Database connection successfully open")
	database.DBConnection.AutoMigrate(&phone.Phone{})
	fmt.Println("Database Migrated")

}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/phone", phone.GetPhones)
	app.Get("api/v1/phone/:id", phone.GetPhone)
	app.Post("api/v1/phone", phone.CreatePhone)
	app.Delete("api/v1/phone/:id", phone.DeletePhone)
	app.Put("api/v1/phone/:id", phone.UpdatePhone)
}
