package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Portfolio
//
// swagger:model
type Portfolio struct {
	// Base model
	Model `gorm:"embedded"`

	// Gallery
	// required: true
	Gallery []Media `json:"gallery" gorm:"foreignKey:Name" validate:"required"`

	// Logo
	// required: true
	Logo Media `json:"logo" validate:"required" gorm:"foreignKey:Name"`

	// Website
	// required: true
	Website string `json:"website" validate:"url,required"`

	// Name
	// required: true
	Name string `json:"name" validate:"required" gorm:"primaryKey"`

	// Description
	// required: true
	Description string `json:"description" validate:"required"`
}

func (s *Portfolio) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Model(s).Preload("Gallery").First(s), nil
}

func (s *Portfolio) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	if err := db.Association("Gallery").Replace(s.Gallery); err != nil {
		return nil, err
	}
	return db.Model(s).Updates(s), nil
}

func (s *Portfolio) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

func (s *Portfolio) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Name = ?", s.Name).Delete(s), nil
}

func (s *Portfolio) DeleteListQuery() Query {
	return &PortfolioDeleteQuery{}
}

func (s *Portfolio) ListQuery() Query {
	return &PortfolioListQuery{}
}

type PortfolioListQuery struct {
	ToFind  string `query:"tofind"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *PortfolioListQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {

	portfolios := new([]Portfolio)
	tmp := db

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("Name LIKE ?", "%"+s.ToFind+"%").Or("Website LIKE ?", "%"+s.ToFind+"%").Or("Description LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Preload("Gallery").Find(portfolios)
	return portfolios, result
}

type PortfolioDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *PortfolioDeleteQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {
	var portfolios []Portfolio

	if result := db.Where("Name IN ?", s.Names).Find(&portfolios); result.Error != nil {
		return result, nil
	}
	return nil, db.Delete(&portfolios, s.Names)
}
