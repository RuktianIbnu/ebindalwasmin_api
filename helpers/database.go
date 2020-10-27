package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

var (
	dbConn = &DB{}
)

// Init ...
func Init() *DB {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PSWD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Duration(300 * time.Second))

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	dbConn.SQL = db

	return dbConn
}
