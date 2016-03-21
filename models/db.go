package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB_CONNECTION *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql",
		"root:BobButtons2015!@tcp(127.0.0.1:3306)/NanaKalaDB")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
    	log.Fatal(err)
	}
	DB_CONNECTION = db
}

// func insert(statement string, data Data) {
// 	db, err := sql.Open("mysql",
// 		"root:BobButtons2015!@tcp(127.0.0.1:3306)/NanaKalaDB")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
	
// 	if err != nil {
//     	log.Fatal(err)
// 	}

// 	db.Prepare()
// }

// func query(statement string, data Data) {
// 	db, err := sql.Open("mysql",
// 		"root:BobButtons2015!@tcp(127.0.0.1:3306)/NanaKalaDB")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
	
// 	if err != nil {
//     	log.Fatal(err)
// 	}

// 	prep, err := db.Prepare(statement)

// 	if err != nil {

// 	}

// 	datalen := data.NumField()
// 	result, err := prep.Exec(data.Field())

// }

