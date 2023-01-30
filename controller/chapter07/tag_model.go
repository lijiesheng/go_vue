package chapter07

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:""`
	Desc    string `gorm:""`
}
