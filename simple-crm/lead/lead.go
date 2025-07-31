package lead

import (
	"simple-crm/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	db := database.DBConn
	if err := db.Find(&lead, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Lead not found")
	}
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	if err := db.Create(&lead).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	if err := db.First(&lead, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Lead not found to delete")
	}
	if err := db.Delete(&lead).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)

}
