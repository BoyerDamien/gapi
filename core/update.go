package core

import (
	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
)

func Update(c *Ctx, db *database.DB, r UpdateRessource) error {

	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := Validate(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	result, err := r.Update(c, db)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusAccepted)
	}
	return c.JSON(r)
}
