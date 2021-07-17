package security

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUserData(c *fiber.Ctx, name string) jwt.MapClaims {
	user := c.Locals(name).(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}
