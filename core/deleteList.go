package core

import (
	"github.com/BoyerDamien/gapi/database"
	fiber "github.com/gofiber/fiber/v2"
)

func deleteList(c *Ctx, db *database.DB, r DeleteListRessource) error {

	query := r.DeleteListQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	if err := Validate(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	result, _ := query.Run(c, db)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusAccepted)
	}
	return c.SendStatus(fiber.StatusOK)
}
