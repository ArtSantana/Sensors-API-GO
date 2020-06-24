package repository

import (
	"database/sql"
	"fmt"
	"sensores/global"

	_ "github.com/go-sql-driver/mysql" // Just for remove the warning
)

var connectionString string = fmt.Sprintf("%s:%s@tcp(%s)/%s", global.Env("USER"), global.Env("PASSWORD"), global.Env("IP"), global.Env("DATABASE"))
var db, err = sql.Open("mysql", connectionString)

func init() {
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the database!")
}
