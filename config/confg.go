package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

// return => error or conection with Database
func Init() error {
	return nil
}
