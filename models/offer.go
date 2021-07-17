package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Offer
//
// swagger:model
type Offer struct {
	// Base model
	Model `gorm:"embedded"`

	// Nom de l'offre
	// required: true
	Name string `json:"name" validate:"required" gorm:"primaryKey"`

	// Description de l'offre
	// required: true
	Description string `json:"description" validate:"required"`

	// Tags liés l'offre
	// required: true
	Tags []Tag `json:"tags" gorm:"many2many:offers_tags;" validate:"required"`
}

func (s *Offer) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Model(s).Preload("Tags").First(s), nil
}

func (s *Offer) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	if err := db.Model(s).Association("Tag").Replace(s.Tags); err != nil {
		return nil, err
	}
	return db.Model(s).Updates(s), nil
}

func (s *Offer) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.FirstOrCreate(s, &Offer{Name: s.Name}), nil
}

func (s *Offer) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Name = ?", s.Name).Delete(s), nil
}

func (s *Offer) DeleteListQuery() Query {
	return &OfferDeleteQuery{}
}

func (s *Offer) ListQuery() Query {
	return &OfferListQuery{}
}

type OfferListQuery struct {
	ToFind  string `query:"tofind" validate:"omitempty"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=name"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *OfferListQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {

	offers := new([]Offer)
	tmp := db

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("Name LIKE ?", "%"+s.ToFind+"%").Or("Description LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Preload("Tags").Find(offers)
	return offers, result
}

type OfferDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *OfferDeleteQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {
	var offers []Offer

	if result := db.Where("Name IN ?", s.Names).Find(&offers); result.Error != nil {
		return result, nil
	}
	return nil, db.Delete(&offers, s.Names)
}
