package database

import (
	"database/sql"
	"log"
	"time"
)

var (DbConn *sql.DB)

func DbConnector(){
	db, err := sql.Open("mysql",
		"gobdd:gobdd@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}

	var dbErr error
	for i := 1; i <= 8; i++ {
		dbErr = db.Ping()
		if dbErr != nil {
			if i < 8 {
				log.Printf("db connection failed, %d retry : %v", i, dbErr)
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}

	defer db.Close()
}
