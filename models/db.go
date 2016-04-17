package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
)

var DB_CONNECTION *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql",
		"user:password@tcp(IpAddr:Port)/DBName")
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
    	fmt.Println(err)
	}

	fmt.Println("connected to db")

	DB_CONNECTION = db
}

func checkErr(err error, res http.ResponseWriter) {
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
}