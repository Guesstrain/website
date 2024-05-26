package database

import (
	"fmt"

	"github.com/Guesstrain/website/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseService interface {
	All() dto.PersonalInfo
}

type database struct {
	connection *gorm.DB
}

func NewDatabaseService() DatabaseService {
	dsn := "root:password@tcp(127.0.0.1:3306)/website?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&dto.PersonalInfo{})
	return &database{
		connection: db,
	}
}

func (db *database) All() dto.PersonalInfo {
	var pi dto.PersonalInfo
	db.connection.First(&pi, 1)
	return pi
}
