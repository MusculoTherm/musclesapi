package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v3"
)

var db *gorm.DB
var cache *redis.Client

func Setup() error {
	InitializeConfig()
	db, err := gorm.Open("mysql", GlobalConfig.DBPath)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	db.SingularTable(true)
	return nil
}
