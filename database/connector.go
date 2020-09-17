package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (DbConn *sql.DB)

func DbConnector(){
	var err error
	DbConn, err = sql.Open("mysql",
		"gobdd:gobdd@tcp(db:3306)/")
	fmt.Println(DbConn)
	if err != nil {
		log.Fatal(err)
	}

	var dbErr error
	for i := 1; i <= 8; i++ {
		dbErr = DbConn.Ping()
		if dbErr != nil {
			if i < 8 {
				log.Printf("db connection failed, %d retry : %v", i, dbErr)
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}
	defer DbConn.Close()
}
