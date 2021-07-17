package endpoints

import (
	"dbsite/config"
	"dbsite/models"
	"dbsite/security"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, db *gorm.DB) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")

	var user models.User
	result := db.Where("Email = ?", email).First(&user)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !security.ComparePwd(pass, user.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["admin"] = user.Role == "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
