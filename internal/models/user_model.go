package models

type User struct {
	ID           uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name         string `gorm:"unique" json:"name" form:"name"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Phone        string `gorm:"unique" json:"phone" form:"phone"`
	HashPassword string
	Password     string `gorm:"-" form:"password"`
}
