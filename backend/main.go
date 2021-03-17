package main

import (
	"github.com/gofiber/fiber"
	phone "github.com/pongsakorn-maker/react-phone-store/schema"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/phone", phone.GetPhones)
	app.Get("api/v1/phone/:id", phone.GetPhone)
	app.Post("api/v1/phone", phone.CreatePhone)
	app.Delete("api/v1/phone", phone.DeletePhone)
	app.Put("api/v1/phone", phone.UpdatePhone)
}
