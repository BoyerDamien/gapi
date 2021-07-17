package ressource

import (
	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database"
)

type Query interface {
	Run(c *gapi.Ctx, db *database.DB) (*database.DB, interface{})
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
	Create(c *gapi.Ctx, db *database.DB) (*database.DB, error)
}

// Retrievable Ressource
type RetrieveRessource interface {
	Retrieve(c *gapi.Ctx, db *database.DB) (*database.DB, error)
}

// Updatable Ressource
type UpdateRessource interface {
	Update(c *gapi.Ctx, db *database.DB) (*database.DB, error)
}

// Deletable Ressource
type DeleteRessource interface {
	Delete(c *gapi.Ctx, db *database.DB) (*database.DB, error)
}
