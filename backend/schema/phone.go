package phone

import (
	"github.com/gofiber/fiber"
)

func GetPhones(c *fiber.Ctx) {
	c.Status(200).Send("Get all phones")
}

func GetPhone(c *fiber.Ctx) {
	c.Status(200).Send("Get a phone")
}

func CreatePhone(c *fiber.Ctx) {
	c.Status(201).Send("Create a new phone")
}

func UpdatePhone(c *fiber.Ctx) {
	c.Status(200).Send("Update a phone")
}

func DeletePhone(c *fiber.Ctx) {
	c.Status(204).Send("Delete a phone")
}
