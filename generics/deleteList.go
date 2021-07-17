package generics

import (
	"github.com/BoyerDamien/gapi/ressource"
	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteList(c *fiber.Ctx, db *gorm.DB, r ressource.DeleteListRessource) error {

	query := r.DeleteListQuery()
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	if err := ressource.Validate(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": err})
	}

	result, _ := query.Run(c, db)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusAccepted)
	}
	return c.SendStatus(fiber.StatusOK)
}
