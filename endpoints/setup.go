package endpoints

import (
	"dbsite/generics"
	"dbsite/models"
	"fmt"
	"log"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB, ressources ...models.Ressource) {
	for _, val := range ressources {
		name := fmt.Sprintf("%T", val)
		splitted := strings.Split(name, ".")
		name = strings.ToLower(splitted[len(splitted)-1])
		uri := fmt.Sprintf("/%s", name)

		valCopy := val
		log.Printf("Setup endpoint: GET %s", uri)

		router.Get(uri, func(c *fiber.Ctx) error {
			return generics.Retrieve(c, db, valCopy)
		})

		log.Printf("Setup endpoint: POST %s", uri)
		router.Post(uri, func(c *fiber.Ctx) error {
			return generics.Create(c, db, valCopy)
		})

		log.Printf("Setup endpoint: PUT %s", uri)
		router.Put(uri, func(c *fiber.Ctx) error {
			return generics.Update(c, db, valCopy)
		})

		log.Printf("Setup endpoint: DELETE %s", uri)
		router.Delete(uri, func(c *fiber.Ctx) error {
			return generics.Delete(c, db, valCopy)
		})

		log.Printf("Setup endpoint: DELETE %ss", uri)
		router.Delete(fmt.Sprintf("%ss", uri), func(c *fiber.Ctx) error {
			return generics.DeleteList(c, db, valCopy.DeleteListQuery())
		})

		log.Printf("Setup endpoint: GET %ss", uri)
		router.Get(fmt.Sprintf("%ss", uri), func(c *fiber.Ctx) error {
			return generics.List(c, db, valCopy.ListQuery())
		})
	}

	log.Printf("Setup endpoint: POST /auth")
	router.Post("/auth", func(c *fiber.Ctx) error {
		return Login(c, db)
	})

}
