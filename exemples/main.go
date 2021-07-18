package main

import (
	"log"

	"github.com/BoyerDamien/gapi"
	"github.com/BoyerDamien/gapi/database/driver/sqlite"
)

func main() {

	app := gapi.New(sqlite.Open("test.db"), gapi.Config{})
	app.StaticFolders("./staticfiles")

	coll := app.Collection("/api/v1")
	coll.AddRessources(&Tag{})
	log.Fatal(app.Listen(":3000"))

	/*if err := config.InitStaticFiles(); err != nil {
		panic(err)
	}

	// Init database && migrate models
	log.Println("Connecting to database...")
	if config.ENV == "PROD" {
		db, err = gorm.Open(mysql.Open(config.DB_PROD_DSN), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(config.DB_TEST_NAME), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	}
	if err := models.MigrateAll(db); err != nil {
		panic(err)
	}*/

	/*
		// Init app
		test, err := database.Open(sqlite.Open("test.db"), &database.Config{})
		app := fiber.New()
		api := app.Group("/api")
		v1 := api.Group("/v1")

		// Middlewares
		middlewares.InitAll(app)

		// Endpoints
		endpoints.Setup(v1, db,
			&models.User{},
			&models.Media{},
			&models.Offer{},
			&models.Portfolio{},
			&models.Tag{},
		)
		log.Fatal(app.Listen(fmt.Sprintf(":%s", config.LISTEN_PORT)))*/
}
