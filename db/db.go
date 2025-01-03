package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Completed   bool
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
