package models

type User struct {
	ID           uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name         string `gorm:"unique" json:"name"`
	Email        string `gorm:"unique" json:"email"`
	Phone        string `gorm:"unique" json:"phone"`
	HashPassword string
	Password     string `gorm:"-"`
}
