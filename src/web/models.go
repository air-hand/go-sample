package web

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func NewDBClient(config *AppConfig) *gorm.DB {
	location := url.QueryEscape(config.TZ)

	dsn := fmt.Sprintf("app:app@tcp(db:3306)/app_db?charset=utf8mb4&parseTime=True&loc=%s", location)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

func AddRecordSample(config *AppConfig) {
	db := NewDBClient(config)
	MigrateModels(db)

	db.Create(&Product{Code: "Foo", Price: 1000})
}
