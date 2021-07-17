package generics

import (
	"github.com/BoyerDamien/gapi/ressource"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Retrieve(c *fiber.Ctx, db *gorm.DB, r ressource.RetrieveRessource) error {

	result, err := r.Retrieve(c, db)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(r)
}
