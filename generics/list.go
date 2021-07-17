package generics

import (
	"github.com/BoyerDamien/gapi/ressource"
	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func List(c *fiber.Ctx, db *gorm.DB, r ressource.ListRessource) error {

	query := r.ListQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	if err := ressource.Validate(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	result, store := query.Run(c, db)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": result.Error.Error()})
	}
	return c.JSON(store)
}
