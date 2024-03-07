package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	var dns = "root:123456@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	return db, nil
}
