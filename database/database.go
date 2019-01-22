package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Initialize() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:@/pmo_demo?parseTime=true")
	Migrate(db)
	return db, err
}
