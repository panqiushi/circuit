package models

type User struct {
	ID    uint `gorm:"primary_key;auto_increment"`
	Name  string
	Email string `gorm:"unique"`
	Phone string `gorm:"unique"`
}
