package core

import (
	"github.com/BoyerDamien/gapi/database"
	fiber "github.com/gofiber/fiber/v2"
)

func List(c *Ctx, db *database.DB, r ListRessource) error {

	query := r.ListQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	if err := Validate(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	result, store := query.Run(c, db)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": result.Error.Error()})
	}
	return c.JSON(store)
}
