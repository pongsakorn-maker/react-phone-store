package phone

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/pongsakorn-maker/react-phone-store/database"
)

type Phone struct {
	gorm.Model
	PhoneModel string `json:"phonemodel"`
}

func GetPhones(c *fiber.Ctx) {
	db := database.DBConnection
	var phones []Phone
	db.Find(&phones)
	c.Status(200).Send("Get all phones")
	c.JSON(phones)
}

func GetPhone(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection
	var phone Phone
	db.Find(&phone, id)
	c.Status(200).Send("Get a phone")
	c.JSON(phone)
}

func CreatePhone(c *fiber.Ctx) {
	db := database.DBConnection
	var phone Phone
	phone.PhoneModel = "Iphone 13"
	db.Create(&phone)
	c.JSON(phone)
	c.Status(201).Send("Creates a new phone")
}

func UpdatePhone(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection
	var phone Phone
	db.First(&phone, id)
	db.Update(&phone)
	c.Status(200).Send("Updates a phone")
}

func DeletePhone(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection
	var phone Phone
	db.First(&phone, id)
	if phone.PhoneModel == "" {
		c.Status(500).Send("No Phone found with given ID")
		return
	}
	db.Delete(&phone)
	c.Status(204).Send("Deletes a phone")
}
