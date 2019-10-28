package model

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

var DB *gorm.DB

func init() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=monster sslmode=disable password=8412098")
	if err != nil {
		panic("db error")
	}
	DB = db
}
