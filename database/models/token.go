package models

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model
	UserID uint
	Token string
	Expiry int64
}