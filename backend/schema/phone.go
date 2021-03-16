package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/pongsakorn-maker/react-phone-store/backend/database"
)

type Phone struct {
	gorm.Model
	Phone string `json:"Phone"`
}

func GetPhones(c *fiber.Ctx) {
	db := database.DBConnection
	var phones []Phone
	db.Find(&phones)
	c.JSON(phones)
}

func GetPhone(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection
	var phone Phone
	db.Find(&phone, id)
	c.JSON(phone)
}

func CreatePhone(c *fiber.Ctx) {
	db := database.DBConnection
	phone := new(Phone)
	if err := c.BodyParser(phone); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&phone)
	c.JSON(phone)
}

func DeletePhone(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection
	var phone Phone
	db.First(&phone, id)
	if phone.Phone == "" {
		c.Status(500).Send("No phone found with given ID")
		return
	}
	db.Delete(&phone)
	c.Send("phone successfully deleted")
}
