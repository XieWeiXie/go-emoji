package database

import (
	"fmt"
	"go-emoji/configs"
	"go-emoji/models"

	"github.com/jinzhu/gorm"
)

var (
	SQLITE3DIALECT  *gorm.DB
	POSTGRESDIALECT *gorm.DB
)

func DataBaeInit(dialect string) *gorm.DB {

	var typeDialect *gorm.DB
	if dialect == "sqlite3" {
		conn, err := gorm.Open(configs.SQLITE3, configs.SQLITEFILE)
		if err != nil {
			panic("connect sqlite3 fail")
		}
		conn.LogMode(true)
		SQLITE3DIALECT = conn
		typeDialect = SQLITE3DIALECT
	}
	if dialect == configs.POSTGRES {
		args := fmt.Sprintf("host=%s port=%s  user=%s password=%s dbname=%s sslmode=%s",
			configs.HOST, configs.PORT, configs.USER, configs.PASSWORD, configs.DATABASE, configs.SSLMODE)
		conn, err := gorm.Open(configs.POSTGRES, args)
		if err != nil {
			panic("connect postgres fail")
		}
		POSTGRESDIALECT = conn
		typeDialect = POSTGRESDIALECT
	}
	return typeDialect
}

func ModelsCollection() []interface{} {
	return []interface{}{
		models.Emoji{},
		models.Version{},
	}
}
