package generics

import (
	"github.com/BoyerDamien/gapi/ressource"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx, db *gorm.DB, r ressource.CreateRessource) error {

	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
	}

	if err := ressource.Validate(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	result, err := r.Create(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err})
	}

	return c.JSON(r)
}
