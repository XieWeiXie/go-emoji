package database

import (
	"fmt"
	"go-emoji/configs"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	POSTGRESDIALECT *gorm.DB
)

func DataBaeInit() *gorm.DB {

	args := fmt.Sprintf("host=%s port=%s  user=%s password=%s dbname=%s sslmode=%s",
		configs.HOST, configs.PORT, configs.USER, configs.PASSWORD, configs.DATABASE, configs.SSLMODE)
	log.Println(args)
	conn, err := gorm.Open(configs.POSTGRES, args)
	if err != nil {
		log.Println(err)
		panic("connect postgres fail")
	}
	POSTGRESDIALECT = conn
	POSTGRESDIALECT.LogMode(true)
	return POSTGRESDIALECT
}
