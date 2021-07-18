package core

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
)

// Fiber overloading
// Config
type Config = fiber.Config

// Context overloading
type Ctx = fiber.Ctx

// App wrapper
type App struct {
	*fiber.App              // Fiber app content
	db         *database.DB // Gorm db content
}

func (s *App) Collection(endpoint string, handlers ...func(*Ctx) error) Router {
	return Router{
		s.Group(endpoint, handlers...),
		s.db,
		endpoint,
	}
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

type Router struct {
	fiber.Router
	db   *database.DB
	base string
}

func (s *Router) migrateRessource(ressource interface{}) {
	log.Printf("%s migration in progress...\n", fmt.Sprintf("%T", ressource))
	s.db.AutoMigrate(ressource)
}

func (s *Router) getRessourceName(ressource interface{}) string {
	name := strings.ToLower(fmt.Sprintf("%T", ressource))
	splitted := strings.Split(name, ".")
	return splitted[len(splitted)-1]
}

func (s *Router) hasMethod(ressource interface{}, method string) bool {
	_, ok := reflect.TypeOf(ressource).MethodByName(method)
	return ok
}

func (s *Router) AddRessources(ressources ...interface{}) {
	for _, val := range ressources {
		s.migrateRessource(val)

		ressourceName := s.getRessourceName(val)
		uri := fmt.Sprintf("%s/%s", s.base, ressourceName)
		valCopy := val

		if s.hasMethod(val, "Retrieve") {
			log.Printf("Add endpoint GET %s", uri)
			s.Get(fmt.Sprintf("%s/:id", ressourceName), func(c *Ctx) error {
				return Retrieve(c, s.db, valCopy.(RetrieveRessource))
			})
		}
		if s.hasMethod(val, "Create") {
			log.Printf("Add endpoint POST %s", uri)
			s.Post(ressourceName, func(c *Ctx) error {
				return Create(c, s.db, valCopy.(CreateRessource))
			})
		}

		if s.hasMethod(val, "Delete") {
			log.Printf("Add endpoint DELETE %s", uri)
			s.Delete(fmt.Sprintf("%s/:id", ressourceName), func(c *Ctx) error {
				return Delete(c, s.db, valCopy.(DeleteRessource))
			})
		}

		if s.hasMethod(val, "Update") {
			log.Printf("Add endpoint PUT %s", uri)
			s.Put(ressourceName, func(c *Ctx) error {
				return Update(c, s.db, valCopy.(UpdateRessource))
			})
		}

		if s.hasMethod(val, "List") {
			log.Printf("Add endpoint GET %ss", uri)
			s.Get(fmt.Sprintf("%ss", ressourceName), func(c *Ctx) error {
				return List(c, s.db, valCopy.(ListRessource))
			})
		}

		if s.hasMethod(val, "DeleteList") {
			log.Printf("Add endpoint DELETE %ss", uri)
			s.Delete(fmt.Sprintf("%ss", ressourceName), func(c *Ctx) error {
				return DeleteList(c, s.db, valCopy.(DeleteListRessource))
			})
		}
	}
}

// Fiber New app overloading
func New(dialector database.Dialector, config ...Config) *App {
	db, err := database.Open(dialector)
	if err != nil {
		panic(err)
	}

	return &App{
		fiber.New(config...),
		db,
	}
}
