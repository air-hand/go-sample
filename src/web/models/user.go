package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string
	Groups []Group `gorm:"many2many:user_groups;"`
}

type Group struct {
	gorm.Model
	Name string
}
