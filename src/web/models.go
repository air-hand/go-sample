package web

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"local.packages/web/models"
)

func NewDBClient(config *DBConnectConfig) *gorm.DB {
	dsn := config.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.User{})
}
