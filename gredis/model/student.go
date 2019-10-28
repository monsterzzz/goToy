package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name        string
	Age         int
	Sex         int
	Description string `gorm:"size:255"`
}
