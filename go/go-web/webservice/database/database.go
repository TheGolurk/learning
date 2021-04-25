package database

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

func SetUpDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:golurk@tcp(127.0.0.1:3306)/nventorydb")
	if err != nil {
		log.Fatal(err)
	}
}
