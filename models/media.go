package models

import (
	"fmt"
	"os"
	"path"

	"dbsite/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Media
//
// swagger:model
type Media struct {
	// Base model
	Model `gorm:"embedded"`

	Path string `json:"-"`

	// Nom du média
	// required: true
	// example: image.png
	Name string `json:"name" gorm:"primaryKey"`

	// Taille du fichier en ko
	// required: true
	// min: 1
	// example: 140
	Size uint `json:"file_size"`

	// Mime type du fichier
	// required: true
	// example: png
	Type string `json:"type"`

	// Url du fichier
	// required: false
	// example: /chemin/vers/le/fichier.png
	Url string `json:"url"`

	// Status du l'image
	// required: true
	// pattern: " protected | open"
	// example: open
	Status string `json:"status"`
}

func (s *Media) AfterDelete(tx *gorm.DB) (err error) {
	return os.Remove(s.Path)
}

func (s *Media) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Model(s).First(s), nil
}

func (s *Media) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	if s.Status == "open" || s.Status == "protected" {
		return db.Model(s).Select("Status").Updates(s), nil
	}
	return nil, fmt.Errorf("wrong status")
}

func (s *Media) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	media, err := c.FormFile("media")
	if err != nil {
		return nil, err
	}

	s.Type = media.Header.Get("Content-Type")
	s.Name = media.Filename
	s.Size = uint(media.Size)
	s.Path = path.Join(config.MEDIA_PATH, media.Filename)
	s.Url = path.Join(config.STATIC_FILES_MEDIA_URL, media.Filename)
	s.Status = "protected"

	result := db.FirstOrCreate(s, s)
	if result.Error != nil {
		return result, nil
	}

	if err := c.SaveFile(media, s.Path); err != nil {
		return db.Where("Name = ?", s.Name).Delete(s), err
	}
	return result, nil
}

func (s *Media) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.Where("Name = ?", s.Name).Delete(s), nil
}

func (s *Media) DeleteListQuery() Query {
	return &MediaDeleteQuery{}
}

func (s *Media) ListQuery() Query {
	return &MediaListQuery{}
}

type MediaListQuery struct {
	ToFind  string `query:"tofind" validate:"omitempty"`
	Type    string `query:"type" validate:"omitempty"`
	Status  string `query:"status" validate:"omitempty,eq=protected|eq=open"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=firstName|eq=name|eq=size"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *MediaListQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {

	medias := new([]Media)
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
	if len(s.Status) > 0 {
		tmp = tmp.Where("Status = ?", s.Status)
	}
	if len(s.Type) > 0 {
		tmp = tmp.Where("Type = ?", s.Type)
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Find(medias)
	return medias, result
}

type MediaDeleteQuery struct {
	Names []string `query:"names"`
}

func (s *MediaDeleteQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {
	var medias []Media

	if result := db.Where("Name IN ?", s.Names).Find(&medias); result.Error != nil {
		return result, nil
	}
	return nil, db.Delete(&medias, s.Names)
}
