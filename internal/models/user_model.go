package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint   `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Name         string `gorm:"unique;column:name" json:"name" form:"name"`
	Email        string `gorm:"unique;column:email" json:"email" form:"email"`
	Phone        string `gorm:"unique;column:phone" json:"phone" form:"phone"`
	HashPassword string `gorm:"column:hash_password"`
	Password     string `gorm:"-" form:"password"`
}
