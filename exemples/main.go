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
}
