package database

import (
	"github.com/jinzhu/gorm"
	"./models"
)

func Migrate(db *gorm.DB)  {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Token{})
	db.Model(&models.Token{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}
