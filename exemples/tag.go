package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Tag
//
// swagger:model
type Tag struct {
	// Base model

	// Nom du Tag
	// required: true
	// example: #python
	Name string `json:"name" validate:"required,min=3,max=255" gorm:"primaryKey"`
}

func (s *Tag) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Name = ?", c.Params("id")).First(s), nil
}

func (s *Tag) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Model(s).Updates(s), nil
}

func (s *Tag) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

func (s *Tag) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Name = ?", c.Params("id")).Delete(s), nil
}
