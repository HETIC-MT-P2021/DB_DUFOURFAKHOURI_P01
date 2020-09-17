package database

import (
	"database/sql"
	"log"
)

var (DbConn *sql.DB)

func DbConnector(){
	db, err := sql.Open("mysql",
		"gobdd:gobdd@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//results, err := db.Query("SELECT customerName FROM customers")

	//fmt.Println(results)
}
