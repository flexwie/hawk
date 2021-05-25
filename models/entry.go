package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model `hawk:"ignore" json:"-"`
	Id         string `gorm:"primaryKey"`
	Name       string
	Start      string
	End        string
	CategoryId string
}
