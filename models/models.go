package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Setup() error {
	InitializeConfig()
	tempDb, err := gorm.Open("sqlite3", "./gorm.db")
	db = &tempDb
	if err != nil {
		return err
	}
	db.LogMode(true)
	db.SingularTable(true)
	db.CreateTable(&Workout{})
	return nil
}
