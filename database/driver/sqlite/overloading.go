package sqlite

import (
	"github.com/BoyerDamien/gapi/database"
	"gorm.io/driver/sqlite"
)

/*
*		Overloading
 */

// Open
func Open(dsn string) database.Dialector {
	return sqlite.Open(dsn)
}
