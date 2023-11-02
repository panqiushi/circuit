package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID     uint   `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name" form:"name"`
	Desc   string `gorm:"column:desc" json:"desc" form:"desc"`
	UserId uint   `gorm:"foreignkey:UserId;column:user_id" json:"user_id"`
	User   User   `json:"-"`
}
