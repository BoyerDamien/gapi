package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Tag
//
// swagger:model
type Tag struct {
	// Base model
	Model `gorm:"embedded"`

	// Nom du Tag
	// required: true
	// example: #python
	Name string `json:"name" validate:"required,min=3,max=255" gorm:"primaryKey"`

	// Offres liées au tag
	// required: false
	Offers []Offer `json:"offers" gorm:"many2many:offers_tags;"`
}

func (s *Tag) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Model(s).Preload("Offers").First(s), nil
}

func (s *Tag) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	if err := db.Model(s).Association("Offers").Replace(s.Offers); err != nil {
		return nil, err
	}
	return db.Model(s).Updates(s), nil
}

func (s *Tag) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

func (s *Tag) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Name = ?", s.Name).Delete(s), nil
}

func (s *Tag) DeleteListQuery() Query {
	return &TagDeleteQuery{}
}

func (s *Tag) ListQuery() Query {
	return &TagListQuery{}
}

type TagListQuery struct {
	ToFind  string `query:"tofind" validate:"omitempty"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *TagListQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {

	tags := new([]Tag)
	tmp := db

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("Name LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Preload("Offers").Find(tags)
	return tags, result
}

type TagDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *TagDeleteQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {
	var tags []Tag

	if result := db.Where("Name IN ?", s.Names).Find(&tags); result.Error != nil {
		return result, nil
	}
	return nil, db.Delete(&tags, s.Names)
}
