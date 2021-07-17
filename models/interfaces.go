package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Query interface {
	Run(db *gorm.DB) (interface{}, *gorm.DB)
}

type Ressource interface {
	Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error)
	Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error)
	Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error)
	Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error)
	DeleteListQuery() Query
	ListQuery() Query
}
