package middlewares

import (
	"dbsite/config"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
)

func InitAll(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
	//app.Use(limiter.New())
	app.Use(func(c *fiber.Ctx) error {

		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.SECRET_KEY),
	}))
}
