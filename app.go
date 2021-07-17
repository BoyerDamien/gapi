package gapi

import (
	"log"
	"os"

	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Fiber overloading
// Config
type Config = fiber.Config

// Context overloading
type Ctx = fiber.Ctx

// App wrapper
type App struct {
	fibApp *fiber.App // Fiber app content
	db     *gorm.DB   // Gorm db content
}

func (s *App) StaticFolders(paths ...string) {
	for _, path := range paths {
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(path, os.ModePerm)
				log.Printf("Path %s created\n", path)
			}
		}
	}
}

func (s *App) Listen(addr string) error {
	return s.fibApp.Listen(addr)
}

// Fiber New app overloading
func New(dialector database.Dialector, config ...Config) *App {
	db, err := database.Open(dialector)
	if err != nil {
		panic(err)
	}

	return &App{
		fibApp: fiber.New(config...),
		db:     db,
	}
}
