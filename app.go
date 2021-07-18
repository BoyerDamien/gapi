package gapi

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/BoyerDamien/gapi/database"
	"github.com/gofiber/fiber/v2"
)

// Fiber overloading
// Config
type Config = fiber.Config

// Context overloading
type Ctx = fiber.Ctx

// App denotes Gapi application
type App struct {
	// Fiber app content
	*fiber.App
	// Gorm db content
	db *database.DB
}

// Collection allows to group handlers behind one endpoints.
// It's simply a wrapper to the fiber's group function:
//	router := app.collection("/api/v1", someHandler, someMiddleware)
func (s *App) Collection(endpoint string, handlers ...func(*Ctx) error) Router {
	return Router{
		s.Group(endpoint, handlers...),
		s.db,
		endpoint,
	}
}

// StaticFolders allows you to specify the locations of you static files:
//	app := gapi.New(sqlite.Open("test.db"), gapi.Config{})
//	app.StaticFolders("./staticfiles")
func (s *App) StaticFolders(paths ...string) {
	for _, path := range paths {
		s.App.Static("/static", path, fiber.Static{
			Compress:      true,
			CacheDuration: 10 * time.Minute,
			MaxAge:        3600,
		})
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(path, os.ModePerm)
				log.Printf("Path %s created\n", path)
			}
		}
	}
}

// Router denotes Gapi router
type Router struct {
	fiber.Router
	db   *database.DB
	base string
}

// Proceed to the ressource migration to the database
func (s *Router) migrateRessource(ressource interface{}) {
	log.Printf("%s migration in progress...\n", fmt.Sprintf("%T", ressource))
	if err := s.db.AutoMigrate(ressource); err != nil {
		panic(err)
	}
}

// Extract the name of ressource
func (s *Router) getRessourceName(ressource interface{}) string {
	name := strings.ToLower(fmt.Sprintf("%T", ressource))
	splitted := strings.Split(name, ".")
	return splitted[len(splitted)-1]
}

// Check whether a ressource has some method or not
func (s *Router) hasMethod(ressource interface{}, method string) bool {
	_, ok := reflect.TypeOf(ressource).MethodByName(method)
	return ok
}

// AddRessources add ressource to a Gapi Router intance.
// You can pass as many ressource as you want:
//	router := app.Collection("/api/v1")
//	router.AddRessources(&User{}, &Media{})
func (s *Router) AddRessources(ressources ...interface{}) {
	for _, val := range ressources {
		s.migrateRessource(val)

		ressourceName := s.getRessourceName(val)
		uri := fmt.Sprintf("%s/%s", s.base, ressourceName)
		valCopy := val

		if s.hasMethod(val, "Retrieve") {
			log.Printf("Add endpoint GET %s", uri)
			s.Get(fmt.Sprintf("%s/:id", ressourceName), func(c *Ctx) error {
				return retrieve(c, s.db, valCopy.(RetrieveRessource))
			})
		}
		if s.hasMethod(val, "Create") {
			log.Printf("Add endpoint POST %s", uri)
			s.Post(ressourceName, func(c *Ctx) error {
				return create(c, s.db, valCopy.(CreateRessource))
			})
		}

		if s.hasMethod(val, "Delete") {
			log.Printf("Add endpoint DELETE %s", uri)
			s.Delete(fmt.Sprintf("%s/:id", ressourceName), func(c *Ctx) error {
				return delete(c, s.db, valCopy.(DeleteRessource))
			})
		}

		if s.hasMethod(val, "Update") {
			log.Printf("Add endpoint PUT %s", uri)
			s.Put(ressourceName, func(c *Ctx) error {
				return update(c, s.db, valCopy.(UpdateRessource))
			})
		}

		if s.hasMethod(val, "ListQuery") {
			log.Printf("Add endpoint GET %ss", uri)
			s.Get(fmt.Sprintf("%ss", ressourceName), func(c *Ctx) error {
				return list(c, s.db, valCopy.(ListRessource))
			})
		}

		if s.hasMethod(val, "DeleteListQuery") {
			log.Printf("Add endpoint DELETE %ss", uri)
			s.Delete(fmt.Sprintf("%ss", ressourceName), func(c *Ctx) error {
				return deleteList(c, s.db, valCopy.(DeleteListRessource))
			})
		}
	}
}

// New creates a new Gapi named instance.
//  app := gapi.New(sqlite.Open("test.db"), gapi.Config{})
// You can pass optional configuration options by passing a Config struct:
//  app := gapi.New(sqlite.Open("test.db"), gapi.Config{
//      Prefork: true,
//      ServerHeader: "Gapi",
//  })
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
