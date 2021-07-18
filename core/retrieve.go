package core

import (
	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
)

func Retrieve(c *Ctx, db *database.DB, r RetrieveRessource) error {

	result, err := r.Retrieve(c, db)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(r)
}
