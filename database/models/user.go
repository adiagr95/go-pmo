package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	FirstName string
	LastName string
	Password string
}

func (u *User) Serialize() map[string]interface{}  {
	return map[string]interface{}{
		"id": u.ID,
		"username": u.Username,
		"first_name": u.FirstName,
		"last_name": u.LastName,
	}
}