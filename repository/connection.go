package repository

import (
	"fmt"
	"sensores/global"
	"sensores/models"

	"github.com/jinzhu/gorm"

	// MySQL dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var connectionString string = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", global.Env("USER"), global.Env("PASSWORD"), global.Env("IP"), global.Env("DATABASE"))

var db, err = gorm.Open("mysql", connectionString)

func init() {
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the database!")

	db.AutoMigrate(&models.Sensor{})
}
