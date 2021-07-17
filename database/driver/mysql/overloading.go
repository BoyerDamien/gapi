package mysql

import (
	"github.com/BoyerDamien/gapi/database"
	"gorm.io/driver/mysql"
)

/*
*		Overloading
 */
func Open(dsn string) database.Dialector {
	return mysql.Open(dsn)
}
