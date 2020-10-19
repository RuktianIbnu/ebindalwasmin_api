package config

import (
	"fmt"
	"log"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"database/sql"
)

// const (
// 	DB_USER = "root"
// 	DB_PASS = ""
// 	DB_HOST = "127.0.0.1"
// 	DB_PORT = 3306
// 	DB_NAME = "ebindalwasmin"
// )

const (
	DB_USER = "user"
	DB_PASS = "GtXgl3ug!Do!"
	DB_HOST = "10.0.30.90"
	DB_PORT = 3306
	DB_NAME = "ebindalwasmin"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}