package postgres

import (
	"github.com/BoyerDamien/gapi/database"
	"gorm.io/driver/postgres"
)

func Open(dsn string) database.Dialector {
	return postgres.Open(dsn)
}
