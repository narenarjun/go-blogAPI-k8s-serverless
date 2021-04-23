package database

import (
	"blogapi/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		log.Fatalf("error connecting to the database,%d", err)
		return nil, err
	}
	return db, nil
}
