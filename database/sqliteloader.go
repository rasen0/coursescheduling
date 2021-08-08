package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sqliteLoader(coursesdb string) (db *gorm.DB,err error) {
	db, err = gorm.Open(sqlite.Open(coursesdb), &gorm.Config{})
    return db, err
}



