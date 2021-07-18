package gapi

import (
	"github.com/BoyerDamien/gapi/database"
)

type Query interface {
	Run(c *Ctx, db *database.DB) (*database.DB, interface{})
}

// List Deletable Ressource
type DeleteListRessource interface {
	DeleteListQuery() Query
}

// Listable Ressource
type ListRessource interface {
	ListQuery() Query
}

// Creatable Ressource
type CreateRessource interface {
	Create(c *Ctx, db *database.DB) (*database.DB, error)
}

// Retrievable Ressource
type RetrieveRessource interface {
	Retrieve(c *Ctx, db *database.DB) (*database.DB, error)
}

// Updatable Ressource
type UpdateRessource interface {
	Update(c *Ctx, db *database.DB) (*database.DB, error)
}

// Deletable Ressource
type DeleteRessource interface {
	Delete(c *Ctx, db *database.DB) (*database.DB, error)
}
