package gapi

import (
	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
)

func create(c *Ctx, db *database.DB, r CreateRessource) error {

	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err.Error()})
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
