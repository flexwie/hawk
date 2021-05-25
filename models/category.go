package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id      string  `gorm:"primaryKey"`
	Name    string  `gorm:"unique"`
	Entries []Entry `gorm:"foreignKey:CategoryId"`
}
