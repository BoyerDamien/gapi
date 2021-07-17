package database

import "gorm.io/gorm"

/*
*		Overloading
 */

// Dialector
type Dialector = gorm.Dialector

// Config
type Config = gorm.Config

// Options
type Option = gorm.Option

// DB
type DB = gorm.DB

// Open
func Open(dialector Dialector, opts ...Option) (*DB, error) {
	return gorm.Open(dialector, opts...)
}
