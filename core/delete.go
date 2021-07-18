package core

import (
	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *Ctx, db *database.DB, r DeleteRessource) error {

	result, err := r.Delete(c, db)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusAccepted)
	}
	return c.SendStatus(fiber.StatusOK)
}
